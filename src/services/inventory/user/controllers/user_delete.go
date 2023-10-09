package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/user"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// DeleteUser PingExample godoc
// @Summary Delete a mirror user
// @Schemes
// @Description Delete a mirrored user from the registry, the user will not be deleted from the cluster
// @Tags User
// @Accept json
// @Produce json
// @Param userId path int true "User id registered"
// @Success 204
// @Failure 404
// @Failure 500
// @Router /{clusterId}/user/{userId} [delete]
func (entity *UserControllerImpl) DeleteUser(c *gin.Context) {
	userIdParam := c.Param("userId")

	userId, err := strconv.ParseInt(userIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("brokerId", userIdParam).Error("Fail to parse brokerId Param")
		c.JSON(http.StatusBadRequest, "Error parsing cluster from url route")
		return
	}

	clusterIdParam := c.Param("clusterId")
	clusterId, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("brokerId", clusterIdParam).Error("Fail to parse clusterId Param")
		c.JSON(http.StatusBadRequest, "Error parsing clusterId from url route")
		return
	}

	userFromDb, err := entity.UserRepository.GetUser(uint(clusterId), uint(userId), c)
	if err != nil {
		log.WithError(err).WithField("userId", userId).Error("Erro ao buscar usuario na base")
		c.JSON(http.StatusBadRequest, "Erro ao buscar ususario na base")
		return
	}

	if userFromDb.IsLocked() {
		log.WithField("user", userFromDb.ID).Warn("Usuario com interacao bloqueada")
		c.JSON(http.StatusBadRequest, gin.H{"error": "usuario com interacao bloqueada"})
		return
	}

	cluster, err := entity.ClusterRepository.GetCluster(uint(clusterId), c)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterId).Error("Erro ao buscar cluster na base")
		c.JSON(http.StatusBadRequest, "Erro ao buscar cluster na base")
		return
	}

	deleteUserRequest := user.DeleteUserRequest{
		RabbitAccess: cluster.GetRabbitMqAccess(),
		Username:     userFromDb.Username,
	}
	err = entity.UserManagement.DeleteUser(deleteUserRequest, c)

	if err != nil {
		log.WithError(err).WithField("request", deleteUserRequest).Error("Erro ao deletar usuario no rabbit")
		c.JSON(http.StatusInternalServerError, "Erro ao deletar usuario na base")
	}
	c.JSON(http.StatusNoContent, "Deleted")
	return
}
