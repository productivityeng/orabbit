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

// ListAllExchanges godoc
// @Summary List all exchanges from cluster
// @Schemes
// @Description List all exchanges from cluster
// @Tags Exchange
// @Accept json
// @Produce json
// @Success 200 {number} Success
// @Param clusterId path int true "Cluster id from where retrieve exchanges"
// @Router /{clusterId}/exchange  [get]
func (ctrl *ExchangeController) ListAllExchanges(c *gin.Context)  {
	clusterId,err := ctrl.parseClusterIdParam(c)
	if err != nil { return}

	cluster,err := ctrl.DependencyLocator.PrismaClient.Cluster.FindUnique(db.Cluster.ID.Equals(clusterId)).Exec(c)
	if err != nil { 
		log.WithContext(c).WithError(err).WithField("clusterId", clusterId).Error("Fail to retrieve cluster")
		c.JSON(http.StatusInternalServerError, "Fail to retrieve cluster")
		return
	}

	cluster_rabbitmq_access := models.GetRabbitMqAccess(cluster)

	exchanges,err := ctrl.DependencyLocator.ExchangeManagement.GetAllExchangesFromCluster(contracts.ListExchangeRequest{
		RabbitAccess: cluster_rabbitmq_access,
	},c)

	if err != nil { 
		log.WithContext(c).WithError(err).WithField("clusterId", clusterId).Error("Fail to retrieve exchanges")
		c.JSON(http.StatusInternalServerError, "Fail to retrieve exchanges")
		return
	}

	c.JSON(http.StatusOK,exchanges)

}



func (ctrl *ExchangeController) parseClusterIdParam(c *gin.Context) (clusterId int,err error) {
	clusterIdParam := c.Param("clusterId")
	clusterIdConv, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithContext(c).WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse clusterId Param")
		c.JSON(http.StatusBadRequest, "Error parsing clusterId from url route")
		return 0,err
	}
	return int(clusterIdConv),nil
}