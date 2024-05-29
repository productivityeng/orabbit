package resources

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/models"
	"github.com/productivityeng/orabbit/pkg/controllers/virtualhost/dto"
	"github.com/productivityeng/orabbit/pkg/rabbitmq/virtualhost"
	log "github.com/sirupsen/logrus"
)

// Import
// @Summary Import a new VirtualHost
// @Schemes
// @Description Import  a new VirtualHost
// @Tags VirtualHost
// @Accept json
// @Produce json
// @Param clusterId path int true "Cluster id from where retrieve virtualhost"
// @Param ImportVirtualHostRequest body dto.ImportVirtualHostRequest true "Request"
// @Success 200 {string} string	"ok"
// @Success 201
// @Failure 404
// @Failure 500
// @Router /{clusterId}/virtualhost/import [post]
func (controller VirtualHostControllerImpl) Import(c *gin.Context) {
	clusterId,err := controller.parseClusterIdParams(c)
	if err != nil { return }

	request,err := controller.parseImportVirtualHostBody(c)
	if err != nil { return }

	fields := log.Fields{"clusterId": clusterId}
	log.WithFields(fields).Info("Looking for rabbitmq cluster")

	cluster,err := controller.DependencyLocator.PrismaClient.Cluster.FindUnique(
		db.Cluster.ID.Equals(clusterId),
	).Exec(c)
		
	if err != nil && errors.Is(err, db.ErrNotFound) { 
		log.WithFields(fields).WithError(err).Error("Cluster not found")
		c.JSON(http.StatusNotFound, gin.H{"message": "Cluster not found"})
		return
	} else if err != nil {
		log.WithFields(fields).WithError(err).Error("Error retrieving cluster")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving cluster"})
		return
	}
	log.WithContext(c).Info("Cluster found")

	rabbitmq_access := models.GetRabbitMqAccess(cluster)
	
	vHostFromCluster, err := controller.DependencyLocator.VirtualHostManagement.GetVirtualHost(virtualhost.GetVirtualHostRequest{
		RabbitAccess: rabbitmq_access,
		Name: request.Name,
	})

	if err != nil { 
		log.WithFields(fields).WithError(err).Error("Error retrieving virtual host")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving virtual host"})
		return
	}

	response,err := controller.saveVirtualHost(dto.SaveVirtualHostDto{
		Name: vHostFromCluster.Name,
		Description: vHostFromCluster.Description,
		DefaultQueueType: vHostFromCluster.DefaultQueueType,
		Tags: vHostFromCluster.Tags,
		ClusterId: clusterId,
	},c)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error trying to save imported virtualHost"})
		return
	}
	c.JSON(http.StatusCreated, response)
}

func (ctrl *VirtualHostControllerImpl) saveVirtualHost(save dto.SaveVirtualHostDto,c *gin.Context) (*dto.GetVirtualHostDto,error) {

		tagsBytes,_ := json.Marshal(save.Tags)
		unique_constraint := db.VirtualHost.UniqueNameClusterid(db.VirtualHost.Name.Equals(save.Name), db.VirtualHost.ClusterID.Equals(save.ClusterId))

		result,err := ctrl.DependencyLocator.PrismaClient.VirtualHost.UpsertOne(unique_constraint).Create(
			db.VirtualHost.Name.Set(save.Name),
			db.VirtualHost.Description.Set(save.Description),
			db.VirtualHost.DefaultQueueType.Set(db.QueueType(save.DefaultQueueType)),
			db.VirtualHost.Tags.Set(tagsBytes),
			db.VirtualHost.Cluster.Link(db.Cluster.ID.Equals(save.ClusterId)),
		).Update(
			db.VirtualHost.Name.Set(save.Name),
			db.VirtualHost.Description.Set(save.Description),
			db.VirtualHost.DefaultQueueType.Set(db.QueueType(save.DefaultQueueType)),
			db.VirtualHost.Tags.Set(tagsBytes),
			db.VirtualHost.Cluster.Link(db.Cluster.ID.Equals(save.ClusterId)),
		).Exec(c)
    if err != nil { 
		log.WithError(err).Error("Error creating virtual host")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating virtual host"})
		return nil,err
	}

	var tagsArray []string

	json.Unmarshal(result.Tags,&tagsArray)

	return &dto.GetVirtualHostDto{
		Name: result.Name,
		Description: result.Description,
		DefaultQueueType: result.DefaultQueueType.String(),
		Tags: tagsArray,
		Id: result.ID,
		ClusterId: result.ClusterID,
		IsInDatabase: true,
		IsInCluster: true,
	},nil
}