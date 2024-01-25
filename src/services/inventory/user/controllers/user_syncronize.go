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
	log "github.com/sirupsen/logrus"
)

// SyncronizeUser godoc
// @Summary Sincronize um ususario no rabbitmq
// @Schemes
// @Description Cria um ususario que esteja na base do ostern e nao exista no cluster
// @Tags User
// @Accept json
// @Produce json
// @Success 201 {number} Success
// @Param clusterId path int true "Cluster id from where retrieve users"
// @Param userId path int true "User id registered"
// @Failure 400
// @Failure 500
// @Router /{clusterId}/user/{userId}/syncronize [post]
func (entity *UserControllerImpl) SyncronizeUser(c *gin.Context)  {
	userId, clusterId, err := entity.parseRequestParams(c)
	if err != nil {
		return
	}

	log.WithFields(log.Fields{
		"userId":    userId,
		"clusterId": clusterId,
	}).Info("looking for cluster and user")

	cluster,err := entity.getCluster(c, int(clusterId))
	if err != nil { return}
	
	user,err := entity.getUser(c, userId)
	if err != nil {return}
	
	err = entity.verifyIfUserIsLocked(c,user)
	if err != nil { return }

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

func (userController *UserControllerImpl) parseRequestParams(context *gin.Context) (userId int,clusterId int, error error) {
	clusterIdParam := context.Param("clusterId")
	clusterIdConv, err := strconv.ParseInt(clusterIdParam, 10, 32)

	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse clusterId Param")
		context.JSON(http.StatusBadRequest, "Error parsing clusterId from url route")
	}


	userIdParam := context.Param("userId")
	userIdConv, err := strconv.ParseInt(userIdParam, 10, 32)

	if err != nil {
		log.WithError(err).WithField("userId", userIdParam).Error("Fail to parse userId Param")
		context.JSON(http.StatusBadRequest, "Error parsing userId from url route")
	}


	return int(userIdConv), int(clusterIdConv), nil

}

func (UserController *UserControllerImpl) verifyIfUserIsLocked(context *gin.Context,user *db.UserModel) (error){
	locker,err := UserController.DependencyLocator.PrismaClient.LockerUser.FindFirst(db.LockerUser.And(db.LockerUser.UserID.Equals(user.ID),
	db.LockerUser.Enabled.Equals(true))).Exec(context)
	
	if errors.Is(err, db.ErrNotFound) { 
		log.WithContext(context).Info("No locker founded")
		return nil
	} else if err != nil { 
		log.WithContext(context).WithError(err).Error("Fail to find locker")
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	log.WithField("locker",locker).Info("Locker founded")
	context.JSON(http.StatusBadRequest, gin.H{"error": "User is locked"})
	return errors.New("User is locked")
}