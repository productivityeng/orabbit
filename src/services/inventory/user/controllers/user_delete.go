package controllers

import (
	"github.com/gin-gonic/gin"
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
