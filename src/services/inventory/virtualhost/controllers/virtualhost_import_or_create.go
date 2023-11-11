package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// ImportOrCreateVirtualHost
// @Summary Import or Create a new VirtualHost
// @Schemes
// @Description Import or Create a new VirtualHost
// @Tags VirtualHost
// @Accept json
// @Produce json
// @Param clusterId path int true "Cluster id from where retrieve virtualhost"
// @Success 200
// @Success 201
// @Failure 404
// @Failure 500
// @Router /{clusterId}/virtualhost [post]
func (controller VirtualHostControllerImpl) ImportOrCreateVirtualHost(c *gin.Context) {
	clusterIdParam := c.Param("clusterId")
	clusterId, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse brokerId Param")
		c.JSON(http.StatusBadRequest, "Error parsing brokerId from url route")
		return
	}

	fields := log.Fields{"clusterId": clusterId}

	log.WithFields(fields).Info("Looking for rabbitmq cluster")
	_, err = controller.ClusterRepository.GetCluster(uint(clusterId), c)

	return
}
