package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/common"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/user"
	"github.com/productivityeng/orabbit/user/dto"
	"github.com/productivityeng/orabbit/user/entities"
	log "github.com/sirupsen/logrus"
	"net/http"
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
func (entity *UserControllerImpl) CreateUser(c *gin.Context) {
	var importUserReuqest dto.ImportOrCreateUserRequest

	err := c.BindJSON(&importUserReuqest)
	if err != nil {
		log.WithContext(c).WithError(err).Error("Fail to parse user request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fields := log.Fields{"request": fmt.Sprintf("%+v", importUserReuqest)}

	log.WithFields(fields).Info("looking for broker")
	broker, err := entity.ClusterRepository.GetCluster(importUserReuqest.ClusterId, c)

	if err != nil {
		log.WithContext(c).WithError(err).Error("Fail to retrieve broker from user")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.WithFields(fields).WithContext(c).Info("verifying if user already exists for this broker")

	exists, err := entity.UserRepository.CheckIfUserExistsForCluster(importUserReuqest.ClusterId, importUserReuqest.Username, c)
	if err != nil {
		log.WithError(err).WithContext(c).Error("Fail to verify if username already exists for this cluster")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if exists {
		log.WithContext(c).Warn("User already exists in this cluster")
		c.JSON(http.StatusBadRequest, gin.H{"error": "[USER_ALREADY_EXISTS_IN_THIS_CLUSTERS]"})
		return
	}

	var passwordHash string
	if importUserReuqest.Create {
		log.WithFields(fields).Info("User want to create a new user")
		result, err := entity.UserManagement.CreateNewUser(user.CreateNewUserRequest{
			RabbitAccess: common.RabbitAccess{
				Host:     broker.Host,
				Port:     broker.Port,
				Username: broker.User,
				Password: broker.Password,
			},
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
		passwordHash, err = entity.UserManagement.GetUserHash(user.GetUserHashRequest{
			RabbitAccess: common.RabbitAccess{
				Host:     broker.Host,
				Port:     broker.Port,
				Username: broker.User,
				Password: broker.Password,
			},
			UserToRetrieveHash: importUserReuqest.Username,
		}, c)

		if err != nil {
			log.WithContext(c).WithError(err).Error("Fail to retrieve password hash for user")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		log.WithFields(fields).WithField("exists", exists).WithContext(c).Info("user not exists for this broker, creating now")
	}

	userCreated, err := entity.UserRepository.CreateUser(&entities.UserEntity{
		Username:     importUserReuqest.Username,
		PasswordHash: passwordHash,
		ClusterId:    broker.ID,
	})

	if err != nil {
		log.WithContext(c).WithError(err).Error("Fail to save user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.GetUserResponse{
		Id:           userCreated.ID,
		Username:     userCreated.Username,
		PasswordHash: userCreated.PasswordHash,
		ClusterId:    userCreated.ClusterId,
		LockedReason: userCreated.Locker.Reason,
	})

}
