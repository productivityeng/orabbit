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
// import all bindings from a exchange, if the other side of binding is a queue or exchange, import it too
func ImportExchangeBindings(DependencyLocator *core.DependencyLocator, cluster *db.ClusterModel,exchangeName string,virtualHostId int,c *gin.Context) error {
	
	virtualHost,err := DependencyLocator.PrismaClient.VirtualHost.FindUnique(db.VirtualHost.ID.Equals(virtualHostId)).Exec(c)
	if err != nil {
		return err
	}

	access := models.GetRabbitMqAccess(cluster)
	
	bindingsFromCluster,err := DependencyLocator.ExchangeManagement.GetExchangeBindings(contracts.GetExchangeBindings{
		RabbitAccess: access,
		ExchangeName: exchangeName,
		VHost: virtualHost.Name,
	},c)

	if err != nil { return err }

	for _, binding := range bindingsFromCluster {

		//binding with default exchange
		if binding.Source == "" { 
			continue
		}

		switch binding.DestinationType { 
			case "queue": {
				break
			}
			case "exchange": { 

				exchange,err := getExchangeByNameAndImportIfIsNotInDatabase(DependencyLocator,cluster.ID,binding.Source,virtualHostId,c)
				if err != nil { 
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return err
				}
				js_arguments_in_bytes,_ := json.Marshal(binding.Arguments)

				binding_created,err := DependencyLocator.PrismaClient.BindingExchangeToExchange.CreateOne(
					db.BindingExchangeToExchange.VirtualHost.Link(db.VirtualHost.ID.Equals(virtualHostId)),
					db.BindingExchangeToExchange.SourceExchange.Link(db.Exchange.ID.Equals(exchange.ID)),
					db.BindingExchangeToExchange.DestinationExchange.Link(db.Exchange.ID.Equals(exchange.ID)),
					db.BindingExchangeToExchange.RoutingKey.Set(binding.RoutingKey),
					db.BindingExchangeToExchange.Arguments.Set(js_arguments_in_bytes),
				).Exec(c)
				if err == nil {
					log.WithField("binding",binding_created).Info("Binding created")
				}else {
					log.WithError(err).Error("Error creating binding")
				}
				break
			}
		}
	}
	return nil
}




func  getExchangeByNameAndImportIfIsNotInDatabase(DependencyLocator *core.DependencyLocator, clusterId int,exchangeName string,virtualHostId int,c *gin.Context) (*db.ExchangeModel,error) {
	exchange,err := DependencyLocator.PrismaClient.Exchange.FindFirst(db.Exchange.Name.Equals(exchangeName)).Exec(c)
	
	if errors.Is(err, db.ErrNotFound) { 
		exchange,errForImport := ImportExchangeByName(DependencyLocator,clusterId,exchangeName,virtualHostId,c)
		if errForImport != nil { 
			return nil,errForImport
		}
		return exchange,nil
	} else if err != nil { 
		return nil,err
	}
	return exchange,nil

}


