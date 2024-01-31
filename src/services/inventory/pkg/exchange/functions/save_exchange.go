package functions

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/pkg/exchange/dto"
	log "github.com/sirupsen/logrus"
)

// saveExchangeInDatabase saves the exchange in the database.
// It takes the requestBody of type dto.CreateExchangeDto, clusterId of type int, and c of type *gin.Context as input parameters.
// It returns the saved exchange of type *db.ExchangeModel and an error if any.
func SaveExchangeInDatabase(DependencyLocator *core.DependencyLocator, requestBody dto.CreateExchangeDto,clusterId int,c *gin.Context) (*db.ExchangeModel,error){
	argumentsJson,err:=  json.Marshal(requestBody.Arguments)
	if err != nil { 
		log.WithContext(c).WithError(err).Error("Fail to marshal queue arguments")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil,err
	
	}

	log.WithContext(c).Info("Saving exchange in database")
	constraint_upset := db.Exchange.UniqueClusteridName(db.Exchange.ClusterID.Equals(clusterId),db.Exchange.Name.Equals(requestBody.Name))

	exchangeSaved,err := DependencyLocator.PrismaClient.Exchange.UpsertOne(constraint_upset).Create(
		db.Exchange.Cluster.Link(db.Cluster.ID.Equals(clusterId)),
		db.Exchange.VirtualHost.Link(db.VirtualHost.ID.Equals(requestBody.VirtualHostId)),
		db.Exchange.Name.Set(requestBody.Name),
		db.Exchange.Internal.Set(requestBody.Internal),
		db.Exchange.Durable.Set(requestBody.Durable),
		db.Exchange.Arguments.Set(argumentsJson),
		db.Exchange.Type.Set(requestBody.Type),
		).Update(
			db.Exchange.Cluster.Link(db.Cluster.ID.Equals(clusterId)),
			db.Exchange.Name.Set(requestBody.Name),
			db.Exchange.Internal.Set(requestBody.Internal),
			db.Exchange.Durable.Set(requestBody.Durable),
			db.Exchange.Arguments.Set(argumentsJson),
			db.Exchange.Type.Set(requestBody.Type),
		).Exec(c)

	if err != nil {
		log.WithContext(c).WithError(err).Error("Fail to save exchange in database")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil,err
	}
	return exchangeSaved,nil
}