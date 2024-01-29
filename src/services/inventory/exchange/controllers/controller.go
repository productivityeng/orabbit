package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/cluster/models"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/exchange/dto"
	log "github.com/sirupsen/logrus"
)

// parseCreateExchangeBody parses the request body into a CreateExchangeDto struct.
// It takes a gin.Context as a parameter and returns a CreateExchangeDto and an error.
// If there is an error parsing the body, it logs the error and returns a bad request response.
func (ctrl *ExchangeController) parseCreateExchangeBody(c *gin.Context) (request dto.CreateExchangeDto,err error){
	err = c.ShouldBindJSON(&request)
	if err != nil {
		log.WithContext(c).WithError(err).Error("Fail to parse body")
		c.JSON(http.StatusBadRequest, "Error parsing body")
		return
	}
	return
}


// parseImportExchangeBody parses the request body into a ImportExchangeRequest struct.
// It takes a gin.Context as a parameter and returns a ImportExchangeRequest and an error.
// If there is an error parsing the body, it logs the error and returns a bad request response.
func (ctrl *ExchangeController) parseImportExchangeBody(c *gin.Context) (request dto.ImportExchangeRequest,err error){
	err = c.ShouldBindJSON(&request)
	if err != nil {
		log.WithContext(c).WithError(err).Error("Fail to parse body")
		c.JSON(http.StatusBadRequest, "Error parsing body")
		return
	}
	return
}




// createExchangeInCluster creates an exchange in the specified cluster.
// It takes a cluster model, a create exchange request body, and a Gin context as parameters.
// It returns an error if the exchange creation fails.
func (ctrl *ExchangeController) createExchangeInCluster(cluster *db.ClusterModel, requestBody dto.CreateExchangeDto,c *gin.Context) (err error) { 
	cluster_rabbitmq_access := models.GetRabbitMqAccess(cluster)

	log.WithContext(c).Info("Creating exchange in cluster")
	err = ctrl.DependencyLocator.ExchangeManagement.CreateExchange(contracts.CreateExchangeRequest{ 
		RabbitAccess: cluster_rabbitmq_access,
		Name: requestBody.Name,
		Type: requestBody.Type,
		ClusterId: cluster.ID,
		Internal: requestBody.Internal,
		Durable: requestBody.Durable,
		Arguments: requestBody.Arguments,
	})

	if err != nil { 
		log.WithContext(c).WithError(err).WithField("clusterId", cluster.ID).Error("Fail to create exnchange in cluster")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	log.WithContext(c).Info("Exchange created in cluster")
	return nil
}



// saveExchangeInDatabase saves the exchange in the database.
// It takes the requestBody of type dto.CreateExchangeDto, clusterId of type int, and c of type *gin.Context as input parameters.
// It returns the saved exchange of type *db.ExchangeModel and an error if any.
func (ctrl *ExchangeController) saveExchangeInDatabase(requestBody dto.CreateExchangeDto,clusterId int,c *gin.Context) (*db.ExchangeModel,error){
	argumentsJson,err:=  json.Marshal(requestBody.Arguments)
	if err != nil { 
		log.WithContext(c).WithError(err).Error("Fail to marshal queue arguments")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil,err
	
	}

	log.WithContext(c).Info("Saving exchange in database")
	constraint_upset := db.Exchange.UniqueClusteridName(db.Exchange.ClusterID.Equals(clusterId),db.Exchange.Name.Equals(requestBody.Name))

	exchangeSaved,err := ctrl.DependencyLocator.PrismaClient.Exchange.UpsertOne(constraint_upset).Create(
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

// getCluster retrieves a cluster from the database based on the given cluster ID.
// It returns the cluster model if found, otherwise it returns an error.
func (ctrl *ExchangeController) getCluster(clusterId int,c *gin.Context) (*db.ClusterModel,error){
	cluster,err := ctrl.DependencyLocator.PrismaClient.Cluster.FindUnique(db.Cluster.ID.Equals(clusterId)).Exec(c)
	if err != nil { 
		log.WithField("clusterId", clusterId).Error("Fail to retrieve cluster")
		return nil,err
	}
	return cluster,nil
}

func (ctrl *ExchangeController) getExchange(exchangeId int,c *gin.Context) (*db.ExchangeModel,error){
	exchange,err := ctrl.DependencyLocator.PrismaClient.Exchange.FindUnique(db.Exchange.ID.Equals(exchangeId)).With(
		db.Exchange.Lockers.Fetch(),
	).Exec(c)
	if err != nil { 
		log.WithField("exchangeId", exchangeId).Error("Fail to retrieve exchange")
		return nil,err
	}
	return exchange,nil
}

func(ctrl ExchangeController) verifyIfExchangeIsLocked(queueId int,c *gin.Context) error {
	result,err :=ctrl.DependencyLocator.PrismaClient.LockerExchange.FindFirst(
		db.LockerExchange.ExchangeID.Equals(queueId),
		db.LockerExchange.Enabled.Equals(true),
	).Exec(c)

	if errors.Is(err, db.ErrNotFound) { 
		return nil
	}else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Exchange is locked"})
		return errors.New("queue is locked")
	}

	return nil
}