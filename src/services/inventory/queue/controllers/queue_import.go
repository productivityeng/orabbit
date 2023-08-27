package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/queue/dto"
	"github.com/productivityeng/orabbit/queue/entities"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/common"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/queue"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (q QueueControllerImpl) ImportQueueFromCluster(c *gin.Context) {
	var queueImportRequest dto.QueueImportRequest

	err := c.BindJSON(&queueImportRequest)
	if err != nil {
		log.WithContext(c).WithError(err).Error("Fail to parse user request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fields := log.Fields{"request": fmt.Sprintf("%+v", queueImportRequest)}

	log.WithFields(fields).Info("looking for broker")
	broker, err := q.ClusterRepository.GetCluster(queueImportRequest.ClusterId, c)

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

	entityToSave := &entities.QueueEntity{
		ClusterID: queueImportRequest.ClusterId,
		Name:      queueFromCluster.Name,
		Type:      queueFromCluster.Type,
		Durable:   queueFromCluster.Durable,
		Arguments: queueFromCluster.Arguments,
	}

	err = q.QueueRepository.Save(entityToSave)
	if err != nil {
		log.WithError(err).Error("Fail to save item")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entityToSave)
}
