package resources

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/controllers/utils"
	"github.com/productivityeng/orabbit/controllers/virtualhost/dto"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/models"
	"github.com/productivityeng/orabbit/pkg/rabbitmq/virtualhost"
	log "github.com/sirupsen/logrus"
)

// DeleteVirtualHost
// @Summary Delete a virtualhost from cluster
// @Schemes
// @Description Delete a virtualhost from cluster
// @Tags VirtualHost
// @Accept json
// @Produce json
// @Param clusterId path int true "Cluster id from where retrieve virtualhost"
// @Param virtualHostId path int true "VirtualHost id from database to delete"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /{clusterId}/virtualhost/{virtualHostId} [delete]
func (controller VirtualHostControllerImpl) DeleteVirtualHost(c *gin.Context) {
	clusterId, err := controller.parseClusterIdParams(c)
	if err != nil { return }
	virtualHostId,err := controller.parseVirtualHostIdParams(c)
	if err != nil { return }

	err = utils.VerifyIfVirtualHostIsLockedById(controller.DependencyLocator.PrismaClient, virtualHostId,c)
	if err != nil { return }

	cluster,err := controller.DependencyLocator.PrismaClient.Cluster.FindUnique(db.Cluster.ID.Equals(clusterId)).Exec(c)
	if errors.Is(err, db.ErrNotFound) {
		log.WithError(err).WithField("clusterId", clusterId).Error("Cluster not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Cluster not found"})
		return 
	}else if err != nil { 
		log.WithError(err).WithField("clusterId", clusterId).Error("Fail to retrieve cluster")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	virtualHost,err := controller.DependencyLocator.PrismaClient.VirtualHost.FindUnique(db.VirtualHost.ID.Equals(virtualHostId)).Exec(c)
	if errors.Is(err, db.ErrNotFound) {
		log.WithError(err).WithField("virtualHostId", virtualHostId).Error("VirtualHost not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "VirtualHost not found"})
		return 
	} else if err != nil {
		log.WithError(err).WithField("virtualHostId", virtualHostId).Error("Fail to retrieve virtualHost")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	rAccess := models.GetRabbitMqAccess(cluster)
	err = controller.DependencyLocator.VirtualHostManagement.DeleteVirtualHost(virtualhost.DeleteVirtualHostRequest{
		RabbitAccess: rAccess,
		Name: virtualHost.Name,
	})
	if err != nil {
		log.WithError(err).WithField("virtualHostId", virtualHostId).Error("Fail to delete virtualHost")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	var tags_array []string

	json.Unmarshal(virtualHost.Tags,&tags_array)
	c.JSON(http.StatusOK, dto.GetVirtualHostDto{
		Id: virtualHost.ID,
		Name: virtualHost.Name,
		Description: virtualHost.Description,
		IsInDatabase: true,
		IsInCluster: false,
		DefaultQueueType: string(virtualHost.DefaultQueueType),
		Tags: tags_array,
		ClusterId: virtualHost.ClusterID,
	})
}