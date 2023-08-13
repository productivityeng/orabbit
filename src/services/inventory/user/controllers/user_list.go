package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/src/packages/common"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// ListUsers
// @Summary Retrieve a mirror user from broker
// @Schemes
// @Description Recovery the details of a specific mirror user that is already imported from the cluster
// @Tags User
// @Accept json
// @Produce json
// @Param clusterId path int true "Cluster id from where retrieve users"
// @Param params query common.PageParam true "Number of items in one page"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /{clusterId}/user [get]
func (userCtrl *UserControllerImpl) ListUsers(c *gin.Context) {

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
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse brokerId Param")
		c.JSON(http.StatusBadRequest, "Error parsing brokerId from url route")
		return
	}
	log.WithField("parameter", clusterIdParam).Info("Looking for list of users")

	result, err := userCtrl.UserRepository.ListUsers(int32(clusterId), param.PageSize, param.PageNumber, c)

	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to retrieve users for the cluster")
		c.JSON(http.StatusBadRequest, "Fail to retrieve users for the cluster")
		return
	}

	c.JSON(http.StatusOK, result)
	return
}
