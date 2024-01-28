package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	"github.com/productivityeng/orabbit/cluster/models"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/rabbitmq/virtualhost"
	"github.com/productivityeng/orabbit/virtualhost/dto"
	log "github.com/sirupsen/logrus"
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
	clusterId, err := controller.parseClusterIdParams(c)
	if err != nil { return }
	fields := log.Fields{"clusterId": clusterId}

	log.WithFields(fields).Info("Looking for rabbitmq cluster")
	cluster, err := controller.getClusterById(clusterId, c)
	if err != nil { return }

	vhosts, err := controller.DependencyLocator.VirtualHostManagement.ListVirtualHosts(virtualhost.ListVirtualHostRequest{
		RabbitAccess: models.GetRabbitMqAccess(cluster),
	})

	if err != nil {
		log.WithError(err).WithFields(fields).Error("Erro ao obter vhosts do cluster")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	vhostsFromDatabase, err := controller.DependencyLocator.PrismaClient.VirtualHost.FindMany(db.VirtualHost.ClusterID.Equals(clusterId)).With(
		db.VirtualHost.Lockers.Fetch(),
	).Exec(c)
	if err != nil {
		log.WithError(err).Error("Erro ao obter a lista de VirtualHosts do banco de dados")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	response := controller.buildListVirtualHostResponse(vhosts,vhostsFromDatabase,clusterId)

	c.JSON(http.StatusOK, response)
}

func (ctrl *VirtualHostControllerImpl) buildListVirtualHostResponse(vhostFromCluster []rabbithole.VhostInfo,vhostFromDatabase []db.VirtualHostModel,clusterId int) []dto.GetVirtualHostDto{
	response := make([]dto.GetVirtualHostDto, 0)


	loop_cluster:for _, vhostFromCluster := range vhostFromCluster {
		for _, vhostFromDatabase := range vhostFromDatabase { 
			if vhostFromCluster.Name == vhostFromDatabase.Name {
				response = append(response, dto.GetVirtualHostDto{
					Id:           vhostFromDatabase.ID,
					Description:  vhostFromDatabase.Description,
					Name:         vhostFromDatabase.Name,
					ClusterId: vhostFromDatabase.ClusterID,
					IsInDatabase: true,
					IsInCluster:  true,
					Tags: 	   vhostFromCluster.Tags,
					DefaultQueueType: vhostFromCluster.DefaultQueueType,
					Lockers: vhostFromDatabase.Lockers(),
				})
				continue loop_cluster
			}
		}
		response = append(response, dto.GetVirtualHostDto{
			Id:          0,
			Description: vhostFromCluster.Description,
			Name:        vhostFromCluster.Name,
			ClusterId: clusterId,
			IsInCluster: true,
			IsInDatabase: false,
			Tags: 	   vhostFromCluster.Tags,
			DefaultQueueType: vhostFromCluster.DefaultQueueType,
			Lockers: make([]db.LockerVirtualHostModel,0),
		})
	}

	loop_database:for _, vhostFromDatabase := range vhostFromDatabase { 
		for _,vhostInResponse := range response { 
			if vhostFromDatabase.Name == vhostInResponse.Name { continue loop_database }
		}

		var tags_array []string

		json.Unmarshal(vhostFromDatabase.Tags,&tags_array)
		response = append(response, dto.GetVirtualHostDto{ 
			Id: vhostFromDatabase.ID,
			Description: vhostFromDatabase.Description,
			Name: vhostFromDatabase.Name,
			ClusterId: vhostFromDatabase.ClusterID,
			IsInCluster: false,
			IsInDatabase: true,
			Tags: tags_array,
			DefaultQueueType: vhostFromDatabase.DefaultQueueType.String(),
			Lockers: vhostFromDatabase.Lockers(),
		})	
	}

	return response
}


