package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/queue/dto"
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

	q.QueueManagement.GetAllQueuesFromCluster()
}
