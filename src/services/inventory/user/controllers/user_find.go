package controllers

import (
	"github.com/gin-gonic/gin"
	dto "github.com/productivityeng/orabbit/user/dto"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// PingExample godoc
// @Summary Retrieve a mirror user from broker
// @Schemes
// @Description Recovery the details of a specific mirror user that is already imported from the cluster
// @Tags User
// @Accept json
// @Produce json
// @Param userId path int true "User id registered"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /{clusterId}/user/{userId} [get]
func (entity *UserControllerImpl) FindUser(c *gin.Context) {
	userIdParam := c.Param("userId")
	userId, err := strconv.ParseInt(userIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("brokerId", userIdParam).Error("Fail to parse brokerId Param")
		c.JSON(http.StatusBadRequest, "Error parsing brokerId from url route")
		return
	}

	user, err := entity.UserRepository.GetUser(int32(userId), c)
	if err != nil {
		log.WithError(err).WithContext(c).Error("Fail to retrieve user by id")
		c.JSON(http.StatusNotFound, gin.H{"error": "[USER_NOT_FOUND]"})
		return
	}

	c.JSON(http.StatusOK, dto.GetUserResponseFromUserEntity(user))
	return

}
