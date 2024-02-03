package functions

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core"
	"github.com/productivityeng/orabbit/db"
	log "github.com/sirupsen/logrus"
)


func GetQueueById(DependencyLocator *core.DependencyLocator, c *gin.Context,queueId int) (*db.QueueModel, error) { 
	queueFromDb,err := DependencyLocator.PrismaClient.Queue.FindUnique(db.Queue.ID.Equals(queueId)).With(
		db.Queue.LockerQueues.Fetch(),
		db.Queue.VirtualHost.Fetch(),
		db.Queue.Cluster.Fetch(),
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

func GetQueueByName(DependencyLocator *core.DependencyLocator,clusterId int,virtualId int,queueName string,c *gin.Context) *db.QueueModel {
	queue,err := DependencyLocator.PrismaClient.Queue.FindFirst(
		db.Queue.Name.Equals(queueName),
		db.Queue.ClusterID.Equals(clusterId),
		db.Queue.VirtualHostID.Equals(virtualId),

	).Exec(c)
	if err != nil { 
		log.WithContext(c).WithError(err).Error("Fail to find queue")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}
	return queue
}