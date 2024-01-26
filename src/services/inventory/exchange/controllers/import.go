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

// ImportExchange godoc
// @Summary Import an exchange from cluster
// @Schemes
// @Description Import an exchange from cluster
// @Tags Exchange
// @Accept json
// @Produce json
// @Success 200 {number} Success
// @Param clusterId path int true "Cluster id from where retrieve exchanges"
// @Param ImportExchangeRequest body dto.ImportExchangeRequest true "Request"
// @Router /{clusterId}/exchange/import  [post]
func (ctrl *ExchangeController) ImportExchange(c *gin.Context)  {
	
	log.WithContext(c).Info("Parsing clusterId param")
	clusterId,err := ctrl.parseClusterIdParam(c)
	if err != nil { return}

	log.WithContext(c).Info("Parsing body")
	requestBody ,err := ctrl.parseImportExchangeBody(c)
	if err != nil { return}

	cluster,err := ctrl.getCluster(clusterId,c)
	if err != nil { return }

	err = ctrl.verifyIfExchangesIsAlreadyInDatabase(requestBody.Name,c)
	if err != nil { return }

	rabbit_access := models.GetRabbitMqAccess(cluster)
	
	exchange,err :=ctrl.DependencyLocator.ExchangeManagement.GetExchangeByName(contracts.GetExchangeRequest{
		RabbitAccess: rabbit_access,
		Name: requestBody.Name,
	},c)

	if err != nil { 
		log.WithContext(c).WithError(err).Error("Fail to retrieve exchange from cluster")
		c.JSON(http.StatusInternalServerError, "Fail to retrieve exchange from cluster")
		return
	}

	exchangeSaved,err := ctrl.saveExchangeInDatabase(dto.CreateExchangeDto{
		Name: exchange.Name,
		Type: exchange.Type,
		Durable: exchange.Durable,
		Internal: exchange.Internal,
		Arguments: exchange.Arguments,
	},clusterId,c)
	if err != nil { return }

	arguments := make(map[string]interface{})
	err = json.Unmarshal(exchangeSaved.Arguments,&arguments)
	if err != nil {
		log.WithContext(c).WithError(err).Error("Fail to unmarshal exchange arguments")
	 }
	returnDto := dto.GetExchangeDto{
		Name: exchangeSaved.Name,
		Type: exchangeSaved.Type,
		Durable: exchangeSaved.Durable,
		Internal: exchangeSaved.Internal,
		Arguments: arguments,
		ClusterId: exchangeSaved.ClusterID,
		IsInCluster: true,
		IsInDatabase: true,
		Id: exchangeSaved.ID,
	 }

	c.JSON(http.StatusOK,returnDto)

}

//verifyIfExchangesIsAlreadyInDatabase 
// verify if exchange is already in database
func (ctrl *ExchangeController) verifyIfExchangesIsAlreadyInDatabase(exchangeName string, c *gin.Context) (error) {
	_,err :=ctrl.DependencyLocator.PrismaClient.Exchange.FindUnique(db.Exchange.Name.Equals(exchangeName)).Exec(c)
	if errors.Is(err, db.ErrNotFound) { 
		return nil
	}
	c.JSON(http.StatusBadRequest, "Exchange already in database")
	return errors.New("Exchange already in database")
}


