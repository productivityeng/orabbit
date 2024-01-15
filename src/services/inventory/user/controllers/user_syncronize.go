package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/cluster/models"
	"github.com/productivityeng/orabbit/db"
	user2 "github.com/productivityeng/orabbit/src/packages/rabbitmq/user"
	"github.com/productivityeng/orabbit/user/dto"
	log "github.com/sirupsen/logrus"
)

// SyncronizeUser godoc
// @Summary Sincronize um ususario no rabbitmq
// @Schemes
// @Description Cria um ususario que esteja na base do ostern e nao exista no cluster
// @Tags User
// @Accept json
// @Produce json
// @Param ImportOrCreateUserRequest body dto.UserSyncronizeRequest true "Request"
// @Success 201 {number} Syccess
// @Failure 400
// @Failure 500
// @Router /{clusterId}/user/syncronize [post]
func (entity *UserControllerImpl) SyncronizeUser(c *gin.Context)  {
	clusterId, syncronizeUserRequest, err := entity.parseRequestParams(c)
	if err != nil {
		return
	}

	fields := log.Fields{"request": fmt.Sprintf("%+v", syncronizeUserRequest)}

	log.WithFields(fields).Info("looking for cluster and user")

	cluster,err := entity.getCluster(c, int(clusterId))
	user,err := entity.getUser(c, syncronizeUserRequest.UserId)

	if err != nil {
		return
	}
	
	err = entity.verifyIfUserIsLocked(c,user)

	if err != nil { 
		return
	}

	createUserRequest := user2.CreateNewUserWithHashPasswordRequest{
		RabbitAccess:     models.GetRabbitMqAccess(cluster),
		UsernameToCreate: user.Username,
		PasswordHash:     user.PasswordHash,
	}
	_, err = entity.UserManagement.CreateNewUserWithHashPassword(createUserRequest, c)

	if err != nil {
		log.WithContext(c).WithField("request", createUserRequest).WithError(err).Error("Erro ao criar usuario no cluster")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar ususario no cluster"})
		return
	}

	c.JSON(http.StatusCreated, "Usuario criado no cluster")
	return

}

func (userController *UserControllerImpl) getUser(context *gin.Context, userId int) (*db.UserModel, error) {
	fields := log.Fields{"userId": fmt.Sprintf("%+v", userId)}

	user, err := userController.DependencyLocator.PrismaClient.User.FindUnique(db.User.ID.Equals(userId)).Exec(context)

	if errors.Is(err, db.ErrNotFound) { 
		log.WithContext(context).WithFields(fields).WithError(err).Error("User not found")
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil,err
	} else if err != nil { 
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil,err
	}

	return user,nil
}

func (userController *UserControllerImpl) getCluster(context *gin.Context, clusterId int) (*db.ClusterModel, error){

	cluster,err := userController.DependencyLocator.PrismaClient.Cluster.FindUnique(db.Cluster.ID.Equals(int(clusterId))).Exec(context)

	if errors.Is(err, db.ErrNotFound) { 
		log.WithContext(context).WithError(err).Error("Cluster not found")
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil,err
	} else if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil,err
	}
	return cluster,nil
}

func (userController *UserControllerImpl) parseRequestParams(context *gin.Context) (int, *dto.UserSyncronizeRequest,  error) {
	clusterIdParam := context.Param("clusterId")
	clusterId, err := strconv.ParseInt(clusterIdParam, 10, 32)

	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse clusterId Param")
		context.JSON(http.StatusBadRequest, "Error parsing clusterId from url route")
	}

	var syncronizeUserRequest dto.UserSyncronizeRequest

	err = context.BindJSON(&syncronizeUserRequest)

	if err != nil {
		log.WithContext(context).WithError(err).Error("Fail to parse user request")
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0,nil,err
	}

	return int(clusterId), &syncronizeUserRequest,nil

}

func (UserController *UserControllerImpl) verifyIfUserIsLocked(context *gin.Context,user *db.UserModel) (error){
	_,err := UserController.DependencyLocator.PrismaClient.Locker.FindFirst(db.Locker.And(db.Locker.ArtifactName.Equals(user.Username),db.Locker.Type.Equals(db.LockerTypeUser))).Exec(context)
	if errors.Is(err, db.ErrNotFound) { 
		log.WithContext(context).Info("No locker founded")
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	} else if err != nil { 
		log.WithContext(context).WithError(err).Error("Fail to find locker")
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	
	context.JSON(http.StatusBadRequest, gin.H{"error": "User is locked"})
	return errors.New("User is locked")
}