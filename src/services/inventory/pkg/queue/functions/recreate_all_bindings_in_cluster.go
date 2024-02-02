package functions

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/core"
	"github.com/productivityeng/orabbit/models"
	"github.com/sirupsen/logrus"
)

// RecreateAllBindingInCluster
// Get all bindinds from a queue in database and recreate in cluster. Bindings where the exchanges are not tracked will be ignored
func RecreateAllBindingInCluster(DependencyLocator *core.DependencyLocator,queueId int,c *gin.Context) error {
	
	queue,err := GetQueueById(DependencyLocator,c,queueId)
	if err != nil { return err }

	virtualHost := queue.VirtualHost()
	cluster := queue.Cluster()

	access := models.GetRabbitMqAccess(cluster)
	
	bindingsForQueue,err := GetBindingsWhereQueueIsDest(DependencyLocator,c,queueId)
	if err != nil { return err }

	logrus.WithField("queue", queue).Info("Recreating all bindings in cluster ...")

	for _,binding := range bindingsForQueue { 

		var arguments map[string]interface{}
		json.Unmarshal([]byte(binding.Arguments), &arguments)
		createBindingRequest := contracts.CreateQueueBindingRequest{
			RabbitAccess: access,
			QueueName: queue.Name,
			ExchangeName: binding.SourceExchange().Name,
			RoutingKey: binding.RoutingKey,
			Arguments: arguments,
			VHost: virtualHost.Name,
		}

		logrus.WithField("request", createBindingRequest).Info("Creating queue binding ...")
		err := DependencyLocator.QueueManagement.CreateQueueBinding(createBindingRequest)
		if err != nil { 
			logrus.WithError(err).WithField("request", createBindingRequest).Error("Fail to create queue binding")
		}else {
			logrus.WithField("request", createBindingRequest).Info("Queue binding created")
		}
	}

	logrus.WithField("queue", queue).Info("All bindings recreated")


	return nil
}
