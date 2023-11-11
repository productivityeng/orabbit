package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/queue/dto"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/queue"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// SyncronizeQueue
// @Summary Syncronize a queue between cluster and ostern
// @Schemes
// @Description Create a queue in a cluster that not exist in cluster but is registered in ostern
// @Tags Queue
// @Accept json
// @Produce json
// @Param clusterId path int true "Cluster id"
// @Success 200
// @Failure 404
// @Failure 500
// @Param QueueImportRequest body dto.QueueSycronizeRequest true "Request"
// @Router /{clusterId}/queue/syncronize [post]
func (q QueueControllerImpl) SyncronizeQueue(c *gin.Context) {
	clusterIdParam := c.Param("clusterId")
	clusterId, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse clusterId Param")
		c.JSON(http.StatusBadRequest, "Error parsing clusterId from url route")
		return
	}

	var queueSyncronizeRequest dto.QueueSycronizeRequest

	err = c.BindJSON(&queueSyncronizeRequest)
	if err != nil {
		log.WithContext(c).WithError(err).Error("Fail to parse syncronize request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fields := log.Fields{"request": fmt.Sprintf("%+v", queueSyncronizeRequest), "clusterId": clusterId}

	queueFromOstern, err := q.QueueRepository.Get(uint(clusterId), queueSyncronizeRequest.QueueId, c)

	if err != nil {
		log.WithError(err).WithFields(fields).Error("Error trying to get a queue from database")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.WithFields(fields).Info("looking for cluster")
	cluster, err := q.ClusterRepository.GetCluster(uint(clusterId), c)

	createQueueResut := queue.CreateQueueRequest{
		RabbitAccess: cluster.GetRabbitMqAccess(),
		Queue:        queueFromOstern.Name,
		Vhost:        "/",
		Type:         queueFromOstern.Type,
		Durable:      queueFromOstern.Durable,
		Arguments:    queueFromOstern.Arguments,
	}
	_, err = q.QueueManagement.CreateQueue(createQueueResut)

	if err != nil {
		log.WithError(err).WithField("request", createQueueResut).Error("Erro ao tentar criar a fila no cluster")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"info": "Fila criada com sucesso no cluster"})

}
