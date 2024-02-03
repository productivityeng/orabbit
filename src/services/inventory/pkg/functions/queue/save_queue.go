package functions

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core"
	"github.com/productivityeng/orabbit/db"
	log "github.com/sirupsen/logrus"
)

// SaveQueueInDatabase saves the queue in the database.
func SaveQueueInDatabase(DependencyLocator *core.DependencyLocator, 
	clusterId int,
	virtualHostId int,
	name string,
	description string,
	durable bool,
	autoDelete bool,
	queueType string,
	arguments map[string]interface{},c *gin.Context) (*db.QueueModel,error){

	argumentsJson,err:=  json.Marshal(arguments)
	if err != nil { 
		log.WithContext(c).WithError(err).Error("Fail to marshal queue arguments")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil,err
	
	}

	log.WithContext(c).Info("Saving queue in database")
	constraint_upset := db.Queue.UniqueNameClusterid(
		db.Queue.Name.Equals(name),
		db.Queue.ClusterID.Equals(clusterId),
	)

	queueSaved,err := DependencyLocator.PrismaClient.Queue.UpsertOne(constraint_upset).Create(
			db.Queue.Name.Set(name),
			db.Queue.Description.Set(description),
			db.Queue.Durable.Set(durable),
			db.Queue.Arguments.Set(argumentsJson),
			db.Queue.Type.Set(db.QueueType(queueType)),
			db.Queue.Cluster.Link(db.Cluster.ID.Equals(clusterId)),
			db.Queue.VirtualHost.Link(db.VirtualHost.ID.Equals(virtualHostId)),
		).Update(
			db.Queue.Name.Set(name),
			db.Queue.Description.Set(description),
			db.Queue.Durable.Set(durable),
			db.Queue.Arguments.Set(argumentsJson),
			db.Queue.Type.Set(db.QueueType(queueType)),
		).Exec(c)

	if err != nil {
		log.WithContext(c).WithError(err).Error("Fail to save queue in database")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil,err
	}
	return queueSaved,nil
}