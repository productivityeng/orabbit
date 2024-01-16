package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
	clusterId,err := controller.parseClusterIdParams(c)
	if err != nil { return }

	fields := log.Fields{"clusterId": clusterId}

	log.WithFields(fields).Info("Looking for rabbitmq cluster")
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
	return
}
