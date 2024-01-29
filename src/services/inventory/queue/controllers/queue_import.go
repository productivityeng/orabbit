package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/queue/dto"
	"github.com/productivityeng/orabbit/rabbitmq/common"
	"github.com/productivityeng/orabbit/rabbitmq/queue"
	"github.com/productivityeng/orabbit/utils"
	log "github.com/sirupsen/logrus"
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

	clusterId, queueImportRequest, err := q.parseImportQueueFromCluster(c)
	if err != nil { return}

	fields := log.Fields{"request": fmt.Sprintf("%+v", queueImportRequest)}

	log.WithFields(fields).Info("looking for broker")

	cluster,err := q.getClusterByid(c,clusterId)
	if err != nil { return }


	queueFromCluster, err := q.QueueManagement.GetQueueFromCluster(queue.GetQueueRequest{
		RabbitAccess: common.RabbitAccess{
			Host:     cluster.Host,
			Port:     cluster.Port,
			Username: cluster.User,
			Password: cluster.Password,
		},
		Queue: queueImportRequest.QueueName,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	virtualHost,err :=q.DependencyLocator.PrismaClient.VirtualHost.FindUnique(db.VirtualHost.Name.Equals(queueFromCluster.Vhost)).Exec(c)
	if errors.Is(err, db.ErrNotFound) { 
		c.JSON(http.StatusConflict, gin.H{"error": fmt.Sprintf("VirtualHost %s is not registered in database",queueFromCluster.Vhost)})
		return
	}else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	

	if queueFromCluster == nil {
		log.WithFields(fields).Warn("Queue not found in cluster")
		c.JSON(http.StatusNotFound, gin.H{"error": "[QUEUE_NOTFOUND_INCLUSTER]"})
		return
	}


	err = utils.VerifyIfVirtualHostIsLockedById(q.DependencyLocator.PrismaClient, virtualHost.ID,c)
	if err != nil { return }

	argumentsJson,err:=  json.Marshal(queueFromCluster.Arguments)
	if err != nil { 
		log.WithFields(fields).WithError(err).Error("Fail to marshal queue arguments")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	
	}

	constraintUpsert := db.Queue.UniqueNameClusterid(db.Queue.Name.Equals(queueFromCluster.Name),db.Queue.ClusterID.Equals(clusterId))
	
	queueSaved,err := q.DependencyLocator.PrismaClient.Queue.UpsertOne(constraintUpsert).Create(
		db.Queue.Name.Set(queueFromCluster.Name),
		db.Queue.Description.Set(queueFromCluster.Name),
		db.Queue.Durable.Set(queueFromCluster.Durable),
		db.Queue.Arguments.Set(argumentsJson),
		db.Queue.Type.Set(db.ParseQueueType(queueFromCluster.Type)),
		db.Queue.Cluster.Link(db.Cluster.ID.Equals(clusterId)),
		db.Queue.VirtualHost.Link(db.VirtualHost.ID.Equals(virtualHost.ID)),
	).Update(
		db.Queue.Name.Set(queueFromCluster.Name),
		db.Queue.Description.Set(queueFromCluster.Name),
		db.Queue.Durable.Set(queueFromCluster.Durable),
		db.Queue.Arguments.Set(argumentsJson),
		db.Queue.Type.Set(db.ParseQueueType(queueFromCluster.Type)),
	).Exec(c)
	

	if err != nil { 
		log.WithFields(fields).WithError(err).Error("Fail to save queue")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	response := dto.GetQueueResponse{
		ID:          queueSaved.ID,
		ClusterID:   queueSaved.ClusterID,
		Name:        queueSaved.Name,
		VHost:       queueSaved.Name,
		Type:        queueSaved.Type.String(),
		IsInCluster: true,
	}
	json.Unmarshal(queueSaved.Arguments, &response.Arguments)
	c.JSON(http.StatusOK, response)
}

func(controler QueueControllerImpl) parseImportQueueFromCluster(c *gin.Context) (int,dto.QueueImportRequest,error) {
	clusterIdParam := c.Param("clusterId")
	clusterId, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse brokerId Param")
		c.JSON(http.StatusBadRequest, "Error parsing brokerId from url route")
		return 0,dto.QueueImportRequest{},err
	}

	var queueImportRequest dto.QueueImportRequest

	err = c.BindJSON(&queueImportRequest)
	if err != nil {
		log.WithContext(c).WithError(err).Error("Fail to parse user request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0,dto.QueueImportRequest{},err
	}

	return int(clusterId),queueImportRequest,nil
}