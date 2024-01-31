package functions

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/core"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/models"
	log "github.com/sirupsen/logrus"
)

// ImportQueueBindings
// import all bindings from a queue
func ImportQueueBindings(DependencyLocator *core.DependencyLocator, cluster *db.ClusterModel,queueId int,queueName string,virtualHostId int,c *gin.Context) error {
	
	virtualHost,err := DependencyLocator.PrismaClient.VirtualHost.FindUnique(db.VirtualHost.ID.Equals(virtualHostId)).Exec(c)
	if err != nil {
		return err
	}

	access := models.GetRabbitMqAccess(cluster)
	
	bindingsFromCluster,err := DependencyLocator.QueueManagement.GetQueueBindingsFromCluster(contracts.GetQueueBindingsRequest{
		RabbitAccess: access,
		Name: queueName,
		VirtualHostName: virtualHost.Name,

	})

	if err != nil { return err }

	for _, binding := range bindingsFromCluster {

		//binding with default exchange
		if binding.Source == "" { 
			continue
		}

		exchange,err := getExchangeByNameAndImportIfIsNotInDatabase(DependencyLocator,cluster,binding.Source,virtualHostId,c)
		if err != nil { 
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return err
		}

		js_arguments_in_bytes,_ := json.Marshal(binding.Arguments)

		binding_created,err := DependencyLocator.PrismaClient.BindingExchangeToQueue.CreateOne(
			db.BindingExchangeToQueue.VirtualHost.Link(db.VirtualHost.ID.Equals(virtualHostId)),
			db.BindingExchangeToQueue.SourceExchange.Link(db.Exchange.ID.Equals(exchange.ID)),
			db.BindingExchangeToQueue.DestinationQueue.Link(db.Queue.ID.Equals(queueId)),
			db.BindingExchangeToQueue.RoutingKey.Set(binding.RoutingKey),
			db.BindingExchangeToQueue.Arguments.Set(js_arguments_in_bytes),
		).Exec(c)
		if err == nil {
			log.WithField("binding",binding_created).Info("Binding created")
		}else {
			log.WithError(err).Error("Error creating binding")
		}
	}
	return nil
}



func  getExchangeByNameAndImportIfIsNotInDatabase(DependencyLocator *core.DependencyLocator, cluster *db.ClusterModel,exchangeName string,virtualHostId int,c *gin.Context) (*db.ExchangeModel,error) {
	exchange,err := DependencyLocator.PrismaClient.Exchange.FindFirst(db.Exchange.Name.Equals(exchangeName)).Exec(c)
	
	if errors.Is(err, db.ErrNotFound) { 
		exchange,errForImport := importExchangeForImportQueueBindings(DependencyLocator,cluster,exchangeName,virtualHostId,c)
		if errForImport != nil { 
			return nil,errForImport
		}
		return exchange,nil
	} else if err != nil { 
		return nil,err
	}
	return exchange,nil

}


func  importExchangeForImportQueueBindings(DependencyLocator *core.DependencyLocator,cluster *db.ClusterModel,exchangeName string,virtualHostId int,c *gin.Context) (*db.ExchangeModel,error) {
	
	virtualHost,err := DependencyLocator.PrismaClient.VirtualHost.FindUnique(db.VirtualHost.ID.Equals(virtualHostId)).Exec(c)
	if err != nil {
		return nil,err
	}

	access := models.GetRabbitMqAccess(cluster)
	
	exchangesFromCluster,err := DependencyLocator.ExchangeManagement.GetExchangeByName(contracts.GetExchangeRequest{
		RabbitAccess: access,
		Name: exchangeName,
		VirtualHostName: virtualHost.Name,
	},c)
	if err != nil { return nil,err}

	jsonArguments,_ := json.Marshal(exchangesFromCluster.Arguments)
	
	exchange,err := DependencyLocator.PrismaClient.Exchange.CreateOne(
		db.Exchange.Cluster.Link(db.Cluster.ID.Equals(cluster.ID)),
		db.Exchange.VirtualHost.Link(db.VirtualHost.ID.Equals(virtualHost.ID)),
		db.Exchange.Name.Set(exchangesFromCluster.Name),
		db.Exchange.Internal.Set(exchangesFromCluster.Internal),
		db.Exchange.Durable.Set(exchangesFromCluster.Durable),
		db.Exchange.Arguments.Set(jsonArguments),
		db.Exchange.Type.Set(exchangesFromCluster.Type),

	).Exec(c)

	if err != nil {
		return nil,err
	}
	return exchange,nil
	
}
