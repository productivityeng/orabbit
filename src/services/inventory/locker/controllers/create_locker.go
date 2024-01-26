package locker

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/locker/dto"
	"github.com/sirupsen/logrus"
)

// CreateLocker
// @Summary Create a locker for a specific artifact in a cluster
// @Schemes
// @Description Create a locker for a specific artifact in a cluster
// @Tags Locker
// @Accept json
// @Produce json
// @Param CreateLockerRequest body dto.CreateLockerRequest true "Request"
// @Param clusterId path int true "Cluster id from where retrieve users"
// @Param lockerType path string true "Artifact name from where retrieve users"
// @Param artifactId path int true "Artifact id from where retrieve users"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /{clusterId}/locker/{lockerType}/{artifactId} [post]
func (ctrl *LockerController) CreateLocker(c *gin.Context) {
	clusterId,lockerType,artifactId,err := ctrl.parseRouteParams(c)
	if err != nil { return}
	createLockerRequest,err := ctrl.parseCreateLockerBody(c)

	if err != nil { return}

	

	switch lockerType { 
		case "queue": {
			ctrl.handleCreateLockerForQueue(*createLockerRequest,clusterId,lockerType,artifactId,c)
			return
		}

		case "user": { 
			ctrl.handleCreateLockerForUser(*createLockerRequest,clusterId,lockerType,artifactId,c)
			return
		}

		case "exchange":{
			ctrl.handleCreateLockerForExchange(*createLockerRequest,clusterId,lockerType,artifactId,c)
			return
		}

		default:
			c.JSON(http.StatusBadRequest,gin.H{"message":"invalid locker type"})
				return

		}
	
}


func (ctrl *LockerController) handleCreateLockerForQueue(createLockerRequest dto.CreateLockerRequest,clusterId int,lockerType string,artifactId int,c *gin.Context) {
			enabledLocker,err := ctrl.getEnabledLockerQueue(clusterId,lockerType,artifactId,c)
			if !errors.Is(err,db.ErrNotFound){
				c.JSON(http.StatusInternalServerError,gin.H{"message":"error retrieving locker for queue"})
				return 
			}
			if enabledLocker != nil {
				c.JSON(http.StatusConflict,gin.H{"message":"enabled locker already exists"})
				return 
			
			}
			lockerQueue,err := ctrl.DependencyLocator.PrismaClient.LockerQueue.CreateOne(
				db.LockerQueue.Queue.Link(db.Queue.ID.Equals(artifactId)),
				db.LockerQueue.Enabled.Set(true),
				db.LockerQueue.Reason.Set(createLockerRequest.Reason),
				db.LockerQueue.UserResponsibleEmail.Set(createLockerRequest.Responsible),
				).Exec(c)

			if err != nil { 
				c.JSON(http.StatusInternalServerError,gin.H{"message":"error creating locker","error": err.Error()})
				return 
			}
			c.JSON(http.StatusCreated,lockerQueue)
}



func (ctrl *LockerController) handleCreateLockerForUser(createLockerRequest dto.CreateLockerRequest,clusterId int,lockerType string,artifactId int,c *gin.Context) {
			enabledLocker,err := ctrl.getEnabledLockerUser(clusterId,lockerType,artifactId,c)
			if err != nil && !errors.Is(err,db.ErrNotFound){
				c.JSON(http.StatusInternalServerError,gin.H{"message":"error retrieving locker for user","error": err.Error()})
				return 
			}
			if enabledLocker != nil {
				c.JSON(http.StatusConflict,gin.H{"message":"enabled locker already exists"})
				return 
			
			}
			lockerQueue,err := ctrl.DependencyLocator.PrismaClient.LockerUser.CreateOne(
				db.LockerUser.User.Link(db.User.ID.Equals(artifactId)),
				db.LockerUser.Enabled.Set(true),
				db.LockerUser.Reason.Set(createLockerRequest.Reason),
				db.LockerUser.UserResponsibleEmail.Set(createLockerRequest.Responsible),
			).Exec(c)

			if err != nil { 
				c.JSON(http.StatusInternalServerError,gin.H{"message":"error creating locker","error": err.Error()})
				return 
			}
			c.JSON(http.StatusCreated,lockerQueue)
}

func (ctrl *LockerController) handleCreateLockerForExchange(createLockerRequest dto.CreateLockerRequest,clusterId int,lockerType string,artifactId int,c *gin.Context) {
			enabledLocker,err := ctrl.getEnabledLockerExhange(clusterId,lockerType,artifactId,c)
			if err != nil && !errors.Is(err,db.ErrNotFound){
				logrus.WithError(err).Error("error retrieving locker for exchange")
				c.JSON(http.StatusInternalServerError,gin.H{"message":"error retrieving locker for exchange","error": err.Error()})
				return 
			}
			if enabledLocker != nil {
				c.JSON(http.StatusConflict,gin.H{"message":"enabled locker already exists"})
				return 
			
			}
			lockerQueue,err := ctrl.DependencyLocator.PrismaClient.LockerExchange.CreateOne(
				db.LockerExchange.Exchange.Link(db.Exchange.ID.Equals(artifactId)),
				db.LockerExchange.Enabled.Set(true),
				db.LockerExchange.Reason.Set(createLockerRequest.Reason),
				db.LockerExchange.UserResponsibleEmail.Set(createLockerRequest.Responsible),
				).Exec(c)

			if err != nil { 
				c.JSON(http.StatusInternalServerError,gin.H{"message":"error creating locker","error": err.Error()})
				return 
			}
			c.JSON(http.StatusCreated,lockerQueue)
}

