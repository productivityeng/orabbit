package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/virtualhost"
	"github.com/productivityeng/orabbit/virtualhost/dto"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// ListVirtualHost
// @Summary Retrieve all virtual hosts from cluster and database
// @Schemes
// @Description Retrieve all virtual hosts from cluster and database
// @Tags VirtualHost
// @Accept json
// @Produce json
// @Param clusterId path int true "Cluster id from where retrieve virtualhost"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /{clusterId}/virtualhost [get]
func (controller VirtualHostControllerImpl) ListVirtualHost(c *gin.Context) {
	clusterIdParam := c.Param("clusterId")
	clusterId, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse brokerId Param")
		c.JSON(http.StatusBadRequest, "Error parsing brokerId from url route")
		return
	}

	fields := log.Fields{"clusterId": clusterId}

	log.WithFields(fields).Info("Looking for rabbitmq cluster")
	cluster, err := controller.ClusterRepository.GetCluster(uint(clusterId), c)

	if err != nil {
		log.WithError(err).WithFields(fields).Error("Fail to retrieve cluster")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	vhosts, err := controller.VirtualHostManagement.ListVirtualHosts(virtualhost.ListVirtualHostRequest{
		RabbitAccess: cluster.GetRabbitMqAccess(),
	})

	if err != nil {
		log.WithError(err).WithFields(fields).Error("Erro ao obter vhosts do cluster")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]dto.GetVirtualHostDto, 0)

	for _, vhostFromCluster := range vhosts {
		response = append(response, dto.GetVirtualHostDto{
			Id:          0,
			Description: vhostFromCluster.Description,
			Name:        vhostFromCluster.Name,
			IsInCluster: true,
		})
	}

	vhostsFromDatabase, err := controller.VirtualHostRepository.ListVirtualHosts(uint(clusterId), 100, 5, c)

	if err != nil {
		log.WithError(err).Error("Erro ao obter a lista de VirtualHosts do banco de dados")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, vhostFromDatabase := range vhostsFromDatabase {
		for _, vhostFromClusterInResponse := range response {
			if vhostFromClusterInResponse.Name == vhostFromDatabase.Name {
				vhostFromClusterInResponse.Id = vhostFromDatabase.ID
				vhostFromClusterInResponse.IsInDatabase = true
				continue
			}
		}

		response = append(response, dto.GetVirtualHostDto{
			Id:           vhostFromDatabase.ID,
			Description:  vhostFromDatabase.Description,
			Name:         vhostFromDatabase.Name,
			IsInDatabase: true,
			IsInCluster:  false,
		})
	}

	c.JSON(http.StatusOK, response)
	return
}
