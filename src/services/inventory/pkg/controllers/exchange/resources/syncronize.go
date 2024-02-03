package resources

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/models"
	log "github.com/sirupsen/logrus"
)

// SyncronizeExchange godoc
// @Summary Syncronize a exchange between cluster and database
// @Schemes
// @Description Syncronize a exchange between cluster and database
// @Tags Exchange
// @Accept json
// @Produce json
// @Success 200 {number} Success
// @Param clusterId path int true "Cluster id from where retrieve exchanges"
// @Param exchangeId path int true "Exchange id from database that will be deleted from cluster"
// @Router /{clusterId}/exchange/{exchangeId}/syncronize  [post]
func (ctrl *ExchangeController) SyncronizeExchange(c *gin.Context)  {
	clusterId,err := ctrl.parseClusterIdParam(c)
	if err != nil { return}
	exchangeId,err := ctrl.parseExchangeIdParam(c)
	if err != nil { return}

	err = ctrl.verifyIfExchangeIsLocked(exchangeId,c)
	if err != nil { return}

	cluster,err := ctrl.DependencyLocator.PrismaClient.Cluster.FindUnique(db.Cluster.ID.Equals(clusterId)).Exec(c)
	if err != nil { 
		log.WithContext(c).WithError(err).WithField("clusterId", clusterId).Error("Fail to retrieve cluster")
		c.JSON(http.StatusInternalServerError, "Fail to retrieve cluster")
		return
	}

	exchange,err := ctrl.DependencyLocator.PrismaClient.Exchange.FindUnique(db.Exchange.ID.Equals(exchangeId)).Exec(c)
	if err != nil { 
		log.WithContext(c).WithError(err).WithField("exchangeId", exchangeId).Error("Fail to retrieve exchange")
		c.JSON(http.StatusInternalServerError, "Fail to retrieve exchange")
		return
	}


	cluster_rabbitmq_access := models.GetRabbitMqAccess(cluster)

	arguments := make(map[string]interface{})
	err = json.Unmarshal(exchange.Arguments,&arguments)
	if err != nil {
		log.WithContext(c).WithError(err).Error("Fail to unmarshal exchange arguments")
	 }
	 
	err = ctrl.DependencyLocator.ExchangeManagement.CreateExchange(contracts.CreateExchangeRequest{
		RabbitAccess: cluster_rabbitmq_access,
		Name: exchange.Name,
		Type: exchange.Type,
		Durable: exchange.Durable,
		Arguments: arguments,
	})

	if err != nil {
		log.WithContext(c).WithError(err).WithField("exchangeId", exchangeId).Error("Fail to create exchange in cluster")
		c.JSON(http.StatusInternalServerError, "Fail to create exchange in cluster")
	}


	c.JSON(http.StatusOK,exchange)

}
