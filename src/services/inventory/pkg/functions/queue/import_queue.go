package functions

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/core"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/pkg/controllers/utils"
	functions "github.com/productivityeng/orabbit/pkg/functions/cluster"
	virtualhost_functions "github.com/productivityeng/orabbit/pkg/functions/virtualhost"
	"github.com/productivityeng/orabbit/pkg/rabbitmq/common"
	"github.com/sirupsen/logrus"
)

// ImportQueue imports a queue from a cluster to the database
func ImportQueue(DependencyLocator *core.DependencyLocator,clusterId int,queueName string,virtualHostId int,c *gin.Context) (*db.QueueModel,error){
	cluster,err := functions.GetCluster(DependencyLocator,1,c)
	if err != nil { 
		logrus.WithField("clusterId",clusterId).WithError(err).Error("Fail to get cluster")
		return nil,err
	}

	virtualHost,err := virtualhost_functions.GetVirtualHostById(DependencyLocator,c,virtualHostId)
	if err != nil { 
		logrus.WithField("virtualHostId",virtualHostId).WithError(err).Error("Fail to get virtual host")
		return nil,err
	}

	queueFromCluster ,err := DependencyLocator.QueueManagement.GetQueueFromCluster(contracts.GetQueueRequest{
		RabbitAccess: common.RabbitAccess{
			Host:     cluster.Host,
			Port:     cluster.Port,
			Username: cluster.User,
			Password: cluster.Password,
		},
		Queue: queueName,
	})
	
	err = utils.VerifyIfVirtualHostIsLockedById(DependencyLocator.PrismaClient, virtualHost.ID,c)


	argumentsJson,err:=  json.Marshal(queueFromCluster.Arguments)
	if err != nil { 
		logrus.WithError(err).Error("Fail to marshal queue arguments")
		return nil,err
	}

	constraintUpsert := db.Queue.UniqueNameClusterid(db.Queue.Name.Equals(queueFromCluster.Name),db.Queue.ClusterID.Equals(clusterId))
	
	queueSaved,err := DependencyLocator.PrismaClient.Queue.UpsertOne(constraintUpsert).Create(
		db.Queue.Name.Set(queueFromCluster.Name),
		db.Queue.Description.Set(queueFromCluster.Name),
		db.Queue.Durable.Set(queueFromCluster.Durable),
		db.Queue.Arguments.Set(argumentsJson),
		db.Queue.Type.Set(db.ParseQueueType(queueFromCluster.Type)),
		db.Queue.Cluster.Link(db.Cluster.ID.Equals(clusterId)),
		db.Queue.VirtualHost.Link(db.VirtualHost.ID.Equals(virtualHostId)),
	).Update(
		db.Queue.Name.Set(queueFromCluster.Name),
		db.Queue.Description.Set(queueFromCluster.Name),
		db.Queue.Durable.Set(queueFromCluster.Durable),
		db.Queue.Arguments.Set(argumentsJson),
		db.Queue.Type.Set(db.ParseQueueType(queueFromCluster.Type)),
	).Exec(c)
	
	return queueSaved,err
}