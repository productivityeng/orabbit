package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/cluster/models"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/db"
	log "github.com/sirupsen/logrus"
)

// DeleteExchange godoc
// @Summary List all exchanges from cluster
// @Schemes
// @Description List all exchanges from cluster
// @Tags Exchange
// @Accept json
// @Produce json
// @Success 200 {number} Success
// @Param clusterId path int true "Cluster id from where retrieve exchanges"
// @Param exchangeId path int true "Exchange id from database that will be deleted from cluster"
// @Router /{clusterId}/exchange/{exchangeId}  [delete]
func (ctrl *ExchangeController) DeleteExchange(c *gin.Context)  {
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

	err = ctrl.DependencyLocator.ExchangeManagement.DeleteExchange(contracts.DeleteExchangeRequest{ 
		RabbitAccess: cluster_rabbitmq_access,
		Name: exchange.Name,
	},c)

	if err != nil {
		log.WithContext(c).WithError(err).WithField("exchangeId", exchangeId).Error("Fail to delete exchange from cluster")
		c.JSON(http.StatusInternalServerError, "Fail to delete exchange from cluster")
	}

	if err != nil { 
		log.WithContext(c).WithError(err).WithField("clusterId", clusterId).Error("Fail to retrieve exchanges")
		c.JSON(http.StatusInternalServerError, "Fail to retrieve exchanges")
		return
	}

	c.JSON(http.StatusOK,exchange)

}



func (ctrl *ExchangeController) parseExchangeIdParam(c *gin.Context) (exchangeId int,err error) {
	exchangeIdParam := c.Param("exchangeId")
	exchangeIdConv, err := strconv.ParseInt(exchangeIdParam, 10, 32)
	if err != nil {
		log.WithContext(c).WithError(err).WithField("exchangeId", exchangeIdParam).Error("Fail to parse exchangeId Param")
		c.JSON(http.StatusBadRequest, "Error parsing exchangeId from url route")
		return 0,err
	}
	return int(exchangeIdConv),nil
}