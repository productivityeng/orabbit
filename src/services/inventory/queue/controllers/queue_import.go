package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/productivityeng/orabbit/queue/dto"
	"github.com/productivityeng/orabbit/queue/entities"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/common"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/queue"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// ImportQueueFromCluster
// @Summary Import or create queue
// @Schemes
// @Description Import existing queue from cluster or creater another one
// @Tags Queue
// @Accept json
// @Produce json
// @Param clusterId path int true "Cluster id from where retrieve users"
// @Success 200
// @Failure 404
// @Failure 500
// @Param QueueImportRequest body dto.QueueImportRequest true "Request"
// @Router /{clusterId}/queue/import [post]
func (q QueueControllerImpl) ImportQueueFromCluster(c *gin.Context) {
	clusterIdParam := c.Param("clusterId")
	clusterId, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse brokerId Param")
		c.JSON(http.StatusBadRequest, "Error parsing brokerId from url route")
		return
	}

	var queueImportRequest dto.QueueImportRequest

	err = c.BindJSON(&queueImportRequest)
	if err != nil {
		log.WithContext(c).WithError(err).Error("Fail to parse user request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fields := log.Fields{"request": fmt.Sprintf("%+v", queueImportRequest)}

	log.WithFields(fields).Info("looking for broker")
	broker, err := q.ClusterRepository.GetCluster(uint(clusterId), c)

	queueFromCluster, err := q.QueueManagement.GetQueueFromCluster(queue.GetQueueRequest{
		RabbitAccess: common.RabbitAccess{
			Host:     broker.Host,
			Port:     broker.Port,
			Username: broker.User,
			Password: broker.Password,
		},
		Queue: queueImportRequest.QueueName,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if queueFromCluster == nil {
		log.WithFields(fields).Warn("Queue not found in cluster")
		c.JSON(http.StatusNotFound, gin.H{"error": "[QUEUE_NOTFOUND_INCLUSTER]"})
		return
	}

	queueToSave := &entities.QueueEntity{
		ClusterId: uint(clusterId),
		Name:      queueFromCluster.Name,
		Type:      queueFromCluster.Type,
		Durable:   queueFromCluster.Durable,
		Arguments: queueFromCluster.Arguments,
	}
	err = q.QueueRepository.Save(queueToSave)

	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			log.WithError(err).Warn(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "[QUEUE_ALREADY_TRACKED]"})
			return
		}
		log.WithError(err).Error("Fail to save item")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.GetQueueResponse{
		ID:          queueToSave.ID,
		ClusterID:   queueToSave.ClusterId,
		Name:        queueToSave.Name,
		VHost:       queueToSave.Name,
		Type:        queueToSave.Type,
		Arguments:   queueToSave.Arguments,
		IsInCluster: true,
	})
}
