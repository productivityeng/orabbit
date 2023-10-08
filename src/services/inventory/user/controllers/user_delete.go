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
	clusterId, err := strconv.ParseInt(userIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("brokerId", clusterIdParam).Error("Fail to parse clusterId Param")
		c.JSON(http.StatusBadRequest, "Error parsing clusterId from url route")
		return
	}

	err = entity.UserRepository.DeleteUser(uint(clusterId), uint(userId), c)
	if err != nil {
		log.WithError(err).WithContext(c).Error("Fail to retrieve user by id")
		c.JSON(http.StatusNotFound, gin.H{"error": "[USER_NOT_FOUND]"})
		return
	}

	c.JSON(http.StatusNoContent, "Deleted")
	return
}

// DeleteRabbitUser
// @Summary Delete a user from rabbitmq cluster and ostern database
// @Schemes
// @Description Completely delete a rabbitmq user
// @Tags User
// @Accept json
// @Produce json
// @Param userId path int true "User id registered"
// @Success 204
// @Failure 404
// @Failure 500
// @Router /{clusterId}/user/rabbitmq/{userId} [delete]
func (entity *UserControllerImpl) DeleteRabbitUser(c *gin.Context) {
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

	userFromDatabase, err := entity.UserRepository.GetUser(uint(clusterId), uint(userId), c)
	if err != nil {
		log.WithError(err).Error("Erro ao buscar usuario")
		c.JSON(http.StatusBadRequest, "Erro ao buscar usuario")
		return
	}

	cluster, err := entity.ClusterRepository.GetCluster(uint(clusterId), c)

	if err != nil {
		log.WithError(err).Error("Erro ao buscar cluster")
		c.JSON(http.StatusBadRequest, "Erro ao buscar cluster")
		return
	}

	err = entity.UserManagement.DeleteUser(user.DeleteUserRequest{
		RabbitAccess: cluster.GetRabbitMqAccess(),
		Username:     userFromDatabase.Username,
	}, c)

	if err != nil {
		log.WithError(err).Error("Erro ao remover usuario do rabbitmq no cluster")
		c.JSON(http.StatusInternalServerError, "Erro ao remover usuario do rabbitmq no cluster")
		return
	}

	err = entity.UserRepository.DeleteUser(uint(clusterId), uint(userId), c)
	if err != nil {
		log.WithError(err).WithContext(c).Error("Fail to retrieve user by id")
		c.JSON(http.StatusNotFound, gin.H{"error": "[USER_NOT_FOUND]"})
		return
	}

	c.JSON(http.StatusNoContent, "Deleted")
	return
}
