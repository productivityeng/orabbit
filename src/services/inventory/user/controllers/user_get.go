package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/src/packages/common"
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
func (userCtrl *UserControllerImpl) ListUser(c *gin.Context) {
	var param common.PageParam

	err := c.BindQuery(&param)
	if err != nil {
		log.WithError(err).Error("Error trying to parse query params")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clusterIdParam := c.Param("clusterId")
	clusterId, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("brokerId", clusterIdParam).Error("Fail to parse brokerId Param")
		c.JSON(http.StatusBadRequest, "Error parsing brokerId from url route")
		return
	}

	userCtrl.UserRepository.ListUsers(int32(clusterId), param.PageNumber, param.PageSize, c)
}
