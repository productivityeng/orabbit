package functions

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/core"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/models"
)


func  ImportExchangeByName(DependencyLocator *core.DependencyLocator,clusterId int,exchangeName string,virtualHostId int,c *gin.Context) (*db.ExchangeModel,error) {
	
	cluster,err := DependencyLocator.PrismaClient.Cluster.FindUnique(db.Cluster.ID.Equals(clusterId)).Exec(c)
	if err != nil { 
		return nil,err
	}
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
