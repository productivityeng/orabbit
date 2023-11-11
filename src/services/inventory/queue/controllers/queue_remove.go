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

// RemoveQueueFromCluster
// @Summary Remove a fila do cluster
// @Schemes
// @Description Remove a fila do cluster mas nao altera o cadastro no ostern
// @Tags Queue
// @Accept json
// @Produce json
// @Param clusterId path int true "Cluster id"
// @Success 200
// @Failure 404
// @Failure 500
// @Param QueueImportRequest body dto.QueueRemoveRequest true "Request"
// @Router /{clusterId}/queue/remove [delete]
func (q QueueControllerImpl) RemoveQueueFromCluster(c *gin.Context) {
	clusterIdParam := c.Param("clusterId")
	clusterId, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse clusterId Param")
		c.JSON(http.StatusBadRequest, "Error parsing clusterId from url route")
		return
	}

	var queueRemoveRequest dto.QueueRemoveRequest

	err = c.BindJSON(&queueRemoveRequest)
	if err != nil {
		log.WithContext(c).WithError(err).Error("Falha ao interpretar a requisicao")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fields := log.Fields{"request": fmt.Sprintf("%+v", queueRemoveRequest), "clusterId": clusterId}

	queueFromOstern, err := q.QueueRepository.Get(uint(clusterId), queueRemoveRequest.QueueId, c)

	if err != nil {
		log.WithError(err).WithFields(fields).Error("Error trying to get a queue from database")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.WithFields(fields).Info("looking for cluster")
	cluster, err := q.ClusterRepository.GetCluster(uint(clusterId), c)

	deleteQueueRequest := queue.DeleteQueueRequest{
		RabbitAccess: cluster.GetRabbitMqAccess(),
		Queue:        queueFromOstern.Name,
	}
	err = q.QueueManagement.DeleteQueue(deleteQueueRequest)

	if err != nil {
		log.WithError(err).WithField("request", deleteQueueRequest).Error("Erro ao tentar deletar a fila no cluster")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"info": "Fila removida do  cluster"})

}
