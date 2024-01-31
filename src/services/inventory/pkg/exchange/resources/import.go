package resources

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/models"
	"github.com/productivityeng/orabbit/pkg/exchange/dto"
	"github.com/productivityeng/orabbit/pkg/exchange/functions"
	"github.com/productivityeng/orabbit/pkg/utils"
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

	err = functions.VerifyIfExchangesIsAlreadyInDatabase(ctrl.DependencyLocator.PrismaClient,requestBody.Name,c)
	if err != nil { return }

	virtualHost,err := ctrl.DependencyLocator.PrismaClient.VirtualHost.FindUnique(db.VirtualHost.Name.Equals(requestBody.VirtualHostName)).Exec(c)
	if errors.Is(err, db.ErrNotFound) { 
		c.JSON(http.StatusBadRequest, gin.H{"error": "VirtualHost not found"})
		return 
	}else if err != nil { 
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to retrieve virtual host"})
		return 
	}

	err = utils.VerifyIfVirtualHostIsLockedById(ctrl.DependencyLocator.PrismaClient, virtualHost.ID,c)
	if err != nil { return }


	rabbit_access := models.GetRabbitMqAccess(cluster)
	
	exchange,err :=ctrl.DependencyLocator.ExchangeManagement.GetExchangeByName(contracts.GetExchangeRequest{
		RabbitAccess: rabbit_access,
		Name: requestBody.Name,
	},c)

	if err != nil { 
		log.WithContext(c).WithError(err).Error("Fail to retrieve exchange from cluster")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to retrieve exchange from cluster"})
		return
	}

	exchangeSaved,err := functions.SaveExchangeInDatabase(ctrl.DependencyLocator,dto.CreateExchangeDto{
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

