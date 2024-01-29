package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/cluster/models"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/queue/dto"
	"github.com/productivityeng/orabbit/rabbitmq/queue"
	log "github.com/sirupsen/logrus"
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
	
	clusterId,queueSyncronizeRequest,err := q.parseSyncronizeQueueParams(c)
	if err != nil { 
		return
	}

	fields := log.Fields{"request": fmt.Sprintf("%+v", queueSyncronizeRequest), "clusterId": clusterId}

	err = q.verifyIfQueueIsLocked(queueSyncronizeRequest.QueueId,c)
	if err != nil { 
		return
	}

	queueFromDb,err := q.getQueueById(c,queueSyncronizeRequest.QueueId)
	if err != nil { return }

	cluster,err := q.getClusterByid(c,clusterId)

	if err != nil { return }
	createQueueResut := queue.CreateQueueRequest{
		RabbitAccess: models.GetRabbitMqAccess(cluster),
		Queue:        queueFromDb.Name,
		Vhost:        "/",
		Type:         queueFromDb.Type.String(),
		Durable:      queueFromDb.Durable,
	}

	err = json.Unmarshal(queueFromDb.Arguments, &createQueueResut.Arguments)
	
	if err != nil { 
		log.WithFields(fields).WithError(err).Error("Fail to unmarshal queue arguments")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}	
	_, err = q.QueueManagement.CreateQueue(createQueueResut)

	if err != nil {
		log.WithError(err).WithField("request", createQueueResut).Error("Erro ao tentar criar a fila no cluster")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"info": "Fila criada com sucesso no cluster"})

}

func (controller QueueControllerImpl) parseSyncronizeQueueParams(c *gin.Context)(clusterId int,request *dto.QueueSycronizeRequest, err error) { 
	clusterIdParam := c.Param("clusterId")
	clusterIdConv, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse clusterId Param")
		c.JSON(http.StatusBadRequest, "Error parsing clusterId from url route")
		return 0,nil,err
	}

	var queueSyncronizeRequest dto.QueueSycronizeRequest

	err = c.BindJSON(&queueSyncronizeRequest)
	if err != nil {
		log.WithContext(c).WithError(err).Error("Fail to parse syncronize request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0,nil,err
	}
	return int(clusterIdConv),&queueSyncronizeRequest,nil
}

func (controller QueueControllerImpl) getClusterByid(c *gin.Context,clusterId int) (*db.ClusterModel,error) {
	cluster,err := controller.DependencyLocator.PrismaClient.Cluster.FindUnique(db.Cluster.ID.Equals(clusterId)).Exec(c)

	if errors.Is(err, db.ErrNotFound) {
		log.WithContext(c).Error("Cluster not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Cluster not found"})
		return nil,err
	 } else if err != nil {
		log.WithContext(c).WithError(err).Error("Fail to find cluster")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil,err
	  }
	return cluster,nil
}

func(controller QueueControllerImpl) getQueueById(c *gin.Context,queueId int) (*db.QueueModel, error) { 
	queueFromDb,err := controller.DependencyLocator.PrismaClient.Queue.FindUnique(db.Queue.ID.Equals(queueId)).With(
		db.Queue.LockerQueues.Fetch(),
		db.Queue.VirtualHost.Fetch(),
	).Exec(c)

	if errors.Is(err, db.ErrNotFound) { 
		log.WithContext(c).Error("Queue not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Queue not found"})
		return nil,err
	} else if err != nil { 
		log.WithContext(c).WithError(err).Error("Fail to find queue")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil,err
	}
	return queueFromDb,nil
}
