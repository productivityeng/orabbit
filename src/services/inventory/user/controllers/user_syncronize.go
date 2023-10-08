package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	user2 "github.com/productivityeng/orabbit/src/packages/rabbitmq/user"
	"github.com/productivityeng/orabbit/user/dto"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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
func (entity *UserControllerImpl) SyncronizeUser(c *gin.Context) {
	clusterIdParam := c.Param("clusterId")
	clusterId, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse clusterId Param")
		c.JSON(http.StatusBadRequest, "Error parsing clusterId from url route")
		return
	}

	var syncronizeUserRequest dto.UserSyncronizeRequest

	err = c.BindJSON(&syncronizeUserRequest)
	if err != nil {
		log.WithContext(c).WithError(err).Error("Fail to parse user request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fields := log.Fields{"request": fmt.Sprintf("%+v", syncronizeUserRequest)}

	log.WithFields(fields).Info("looking for cluster")
	cluster, err := entity.ClusterRepository.GetCluster(uint(clusterId), c)

	if err != nil {
		log.WithContext(c).WithError(err).WithField("clusterId", clusterId).Error("Erro ao encontrar o cluster")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := entity.UserRepository.GetUser(uint(clusterId), syncronizeUserRequest.UserId, c)
	if err != nil {
		log.WithContext(c).WithError(err).WithField("userid", syncronizeUserRequest.UserId).Error("Erro ao encontrar usuario")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createUserRequest := user2.CreateNewUserWithHashPasswordRequest{
		RabbitAccess:     cluster.GetRabbitMqAccess(),
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
