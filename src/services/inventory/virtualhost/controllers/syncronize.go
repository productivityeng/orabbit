package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/cluster/models"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/rabbitmq/virtualhost"
	"github.com/productivityeng/orabbit/virtualhost/dto"
	"github.com/sirupsen/logrus"
)

// Syncronize
// @Summary Syncronize a virtualhost from database with cluster
// @Schemes
// @Description Syncronize a virtualhost from database with cluster
// @Tags VirtualHost
// @Accept json
// @Produce json
// @Param clusterId path int true "Cluster id from where retrieve virtualhost"
// @Param virtualHostId path int true "VirtualHost id from database to delete"
// @Success 200 {string} string	"ok"
// @Success 201
// @Failure 404
// @Failure 500
// @Router /{clusterId}/virtualhost/{virtualHostId}/syncronize [post]
func (controller VirtualHostControllerImpl) Syncronize(c *gin.Context) {
	logrus.WithContext(c).Info("Syncronize virtualhost")

	clusterId,err := controller.parseClusterIdParams(c)
	if err != nil { return }
	logrus.WithContext(c).WithField("clusterId",clusterId).Info("Parsed clusterId param")

	virtualHostId,err := controller.parseVirtualHostIdParams(c)
	if err != nil { return }
	logrus.WithContext(c).WithField("virtualHostId",virtualHostId).Info("Parsed virtualHostId param")

	err = controller.verifyIfVirtualHostIsLocked(virtualHostId,c)
	if err != nil { return }
	logrus.WithContext(c).WithField("virtualHostId",virtualHostId).Info("VirtualHost is not locked")
	
	cluster,err := controller.getClusterById(clusterId,c)
	if err != nil { return }
	logrus.WithContext(c).WithField("clusterId",clusterId).Info("Cluster found")


	virtualHost,err := controller.DependencyLocator.PrismaClient.VirtualHost.FindUnique(db.VirtualHost.ID.Equals(virtualHostId)).Exec(c)
	if errors.Is(err, db.ErrNotFound) {
		logrus.WithContext(c).WithField("virtualHostId",virtualHostId).Info("VirtualHost not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "VirtualHost not found"})
		return 
	} else if err != nil { 
		logrus.WithContext(c).WithError(err).Error("Error while retrieving virtualHost from database")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithContext(c).WithField("virtualHostId",virtualHostId).Info("VirtualHost found")

	rAccess := models.GetRabbitMqAccess(cluster)
	var tags []string
	json.Unmarshal([]byte(virtualHost.Tags), &tags)

	logrus.WithContext(c).WithField("virtualHostId",virtualHostId).Info("Creating virtualHost in cluster")
	err = controller.DependencyLocator.VirtualHostManagement.CreateVirtualHost(
		virtualhost.CreateVirtualHostRequest{
			RabbitAccess: rAccess,
			Name: virtualHost.Name,
			Description: virtualHost.Description,
			DefaultQueueType: virtualHost.DefaultQueueType.String(),
			Tags: tags,
		},
	)
	if err != nil { 
		logrus.WithContext(c).WithError(err).Error("Error while creating virtualHost in cluster")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logrus.WithContext(c).WithField("virtualHostId",virtualHostId).Info("VirtualHost created in cluster")

	c.JSON(http.StatusCreated, dto.GetVirtualHostDto{
		Id: virtualHost.ID,
		Name: virtualHost.Name,
		Description: virtualHost.Description,
		DefaultQueueType: virtualHost.DefaultQueueType.String(),
		ClusterId: virtualHost.ClusterID,
		Tags: tags,
		IsInDatabase: true,
		IsInCluster: true,
	} )
}
