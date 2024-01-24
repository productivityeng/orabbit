package locker

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/locker/dto"
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

	enabledLocker,err := ctrl.getEnabledLockerQueue(clusterId,lockerType,artifactId,c)
	if !errors.Is(err,db.ErrNotFound){
		c.JSON(http.StatusInternalServerError,gin.H{"message":"error retrieving locker"})
		return
	}
	if enabledLocker != nil {
		c.JSON(http.StatusConflict,gin.H{"message":"enabled locker already exists"})
		return
	
	}

	switch lockerType { 
		case "queue": {
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
	

			c.JSON(http.StatusOK,lockerQueue)
		}

	default:
		c.JSON(http.StatusBadRequest,gin.H{"message":"invalid locker type"})
			return

	}
	
}

func (ctrl *LockerController) parseCreateLockerBody(c *gin.Context) (*dto.CreateLockerRequest,error){
	var createLockerRequest dto.CreateLockerRequest
	err := c.BindJSON(&createLockerRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"error parsing request body"})
		return nil,err
	}
	return &createLockerRequest,nil
}

func (ctrl *LockerController) parseDisableLockerBody(c *gin.Context) (*dto.DisableLockerRequest,error){
	var disableLockerRequest dto.DisableLockerRequest
	err := c.BindJSON(&disableLockerRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"error parsing request body"})
		return nil,err
	}
	return &disableLockerRequest,nil
}