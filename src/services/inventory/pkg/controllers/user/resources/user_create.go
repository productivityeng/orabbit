package resources

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/models"
	"github.com/productivityeng/orabbit/pkg/controllers/user/dto"
	log "github.com/sirupsen/logrus"
)

// CreateUser godoc
// @Summary Syncronize a existing RabbitMQ user from the broker.
// @Schemes
// @Description Create a new <b>RabbitMQ User mirror</b> from the broker. The user must exist in the cluster, the login and hashpassword will be imported
// @Tags User
// @Accept json
// @Produce json
// @Param ImportOrCreateUserRequest body dto.ImportOrCreateUserRequest true "Request"
// @Success 201 {number} Syccess
// @Failure 400
// @Failure 500
// @Router /{clusterId}/user [post]
func (controller *UserControllerImpl) CreateUser(c *gin.Context) {
	
	importUserReuqest,err := controller.parseInputRequest(c)
	if err != nil { 
		return
	}
	fields := log.Fields{"request": fmt.Sprintf("%+v", importUserReuqest)}

	log.WithFields(fields).Info("looking for broker")
	cluster,err := controller.getCluster(c, importUserReuqest.ClusterId)
	if err != nil {
		
		return
	}

	log.WithFields(fields).WithContext(c).Info("verifying if user already exists for this broker")
	
	userFromDatabase,err := controller.DependencyLocator.PrismaClient.User.FindUnique(db.User.UniqueUsernameClusterid(db.User.Username.Equals(importUserReuqest.Username), db.User.ClusterID.Equals(importUserReuqest.ClusterId))).Exec(c)

	if userFromDatabase != nil{ 
		log.WithContext(c).Warn("User already exists in this cluster")
		c.JSON(http.StatusBadRequest, gin.H{"error": "[USER_ALREADY_EXISTS_IN_THIS_CLUSTERS]"})
		return
	} else if err != nil  { 
		if !errors.Is(err, db.ErrNotFound) {	
			log.WithError(err).WithContext(c).Error("Fail to verify if username already exists for this cluster")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return 
		}

	
	}
	

	access := models.GetRabbitMqAccess(cluster)

	var passwordHash string
	if importUserReuqest.Create {
		log.WithFields(fields).Info("User want to create a new user")
		result, err := controller.DependencyLocator.UserManagement.CreateNewUser(contracts.CreateNewUserRequest{
			RabbitAccess: access,
			UserToCreate:            importUserReuqest.Username,
			PasswordOfUsertToCreate: importUserReuqest.Password,
		}, c)

		if err != nil {
			log.WithError(err).Error("Fail to create a new user")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		passwordHash = result.PasswordHash
	} else {
		log.WithFields(fields).Info("broker founded")
		log.WithFields(fields).Info("looking for passwordhash")

		access := models.GetRabbitMqAccess(cluster)
		passwordHash, err = controller.DependencyLocator.UserManagement.GetUserHash(contracts.GetUserHashRequest{
			RabbitAccess:access,
			UserToRetrieveHash: importUserReuqest.Username,
		}, c)

		if err != nil {
			log.WithContext(c).WithError(err).Error("Fail to retrieve password hash for user")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	}

	userCreated,err := controller.DependencyLocator.PrismaClient.User.CreateOne(
		db.User.Username.Set(importUserReuqest.Username),
		db.User.PasswordHash.Set(passwordHash),
		db.User.Cluster.Link(
			db.Cluster.ID.Equals(importUserReuqest.ClusterId),
		),
	).Exec(c)

	if err != nil { 
		log.WithError(err).Error("Fail to create user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusCreated, dto.GetUserResponse{
		Id:           userCreated.ID,
		Username:     userCreated.Username,
		PasswordHash: userCreated.PasswordHash,
		ClusterId:    userCreated.ClusterID,
	})

}

func (ctrl *UserControllerImpl) parseInputRequest(c *gin.Context) (*dto.ImportOrCreateUserRequest, error) {
	var importUserReuqest dto.ImportOrCreateUserRequest

	fields := log.Fields{"request": fmt.Sprintf("%+v", importUserReuqest)}

	err := c.BindJSON(&importUserReuqest)
	if err != nil {
		log.WithContext(c).WithFields(fields).WithError(err).Error("Fail to parse user request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil,err
	}
	return &importUserReuqest,nil

}
