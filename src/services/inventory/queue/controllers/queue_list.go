package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/queue/dto"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/common"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/queue"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// ListQueuesFromCluster
// @Summary Retrieve all users from rabbitmq cluster
// @Schemes
// @Description Retrieve all users that exist on rabbit cluster. Event if it its registered in ostern
// @Tags Queue
// @Accept json
// @Produce json
// @Param clusterId path int true "Cluster id from where retrieve users"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /{clusterId}/queue/queuesfromcluster [get]
func (q QueueControllerImpl) ListQueuesFromCluster(c *gin.Context) {
	clusterIdParam := c.Param("clusterId")
	clusterId, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse brokerId Param")
		c.JSON(http.StatusBadRequest, "Error parsing brokerId from url route")
		return
	}

	fields := log.Fields{"clusterId": clusterId}

	log.WithFields(fields).Info("Looking for rabbitmq cluster")
	cluster, err := q.ClusterRepository.GetCluster(uint(clusterId), c)

	if err != nil {
		log.WithError(err).WithFields(fields).Error("Fail to retrieve cluster")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	queuesFromCluster, err := q.QueueManagement.GetAllQueuesFromCluster(queue.ListQueuesRequest{
		RabbitAccess: common.RabbitAccess{
			Host:     cluster.Host,
			Port:     cluster.Port,
			Username: cluster.User,
			Password: cluster.Password,
		},
	})
	if err != nil {
		log.WithError(err).WithFields(fields).Error("Fail to retrieve all queuesFromCluster from cluster")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	queuesFromDatabase, err := q.QueueRepository.List(uint(clusterId), c)

	var getQueueResponse dto.GetQueueResponseList

	for _, queueFromCluster := range queuesFromCluster {

		var queueFromClusterExistsInDatabase = false
		var queueIdFromDatabase = uint(0)
		if queueFromBd := queuesFromDatabase.GetQueueFromListByName(queueFromCluster.Name); queueFromBd != nil {
			queueFromClusterExistsInDatabase = true
			queueIdFromDatabase = queueFromBd.ID
		}

		getQueueResponse = append(getQueueResponse, dto.GetQueueResponse{
			ID:           queueIdFromDatabase,
			ClusterID:    uint(clusterId),
			Name:         queueFromCluster.Name,
			VHost:        queueFromCluster.Vhost,
			Type:         queueFromCluster.Type,
			IsInCluster:  true,
			IsInDatabase: queueFromClusterExistsInDatabase,
			Arguments:    queueFromCluster.Arguments,
			Durable:      queueFromCluster.Durable,
		})

	}

	for _, queueFromDb := range queuesFromDatabase {
		if queueFromResponse := getQueueResponse.GetByName(queueFromDb.Name); queueFromResponse == nil {
			getQueueResponse = append(getQueueResponse, dto.GetQueueResponse{
				ID:           queueFromDb.ID,
				ClusterID:    uint(clusterId),
				Name:         queueFromDb.Name,
				VHost:        "/",
				Type:         queueFromDb.Type,
				IsInCluster:  false,
				IsInDatabase: true,
				Arguments:    queueFromDb.Arguments,
				Durable:      queueFromDb.Durable,
			})
		}
	}

	c.JSON(http.StatusOK, getQueueResponse)
}
