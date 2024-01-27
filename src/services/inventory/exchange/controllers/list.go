package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/cluster/models"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/exchange/dto"
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

	exchangesFromCluster,err := ctrl.DependencyLocator.ExchangeManagement.GetAllExchangesFromCluster(contracts.ListExchangeRequest{
		RabbitAccess: cluster_rabbitmq_access,
	},c)

	if err != nil { 
		log.WithContext(c).WithError(err).WithField("clusterId", clusterId).Error("Fail to retrieve exchanges from cluster")
		c.JSON(http.StatusInternalServerError, "Fail to retrieve exchanges from cluster")
		return
	}

	exchangesFromDatabase, err := ctrl.DependencyLocator.PrismaClient.Exchange.FindMany(db.Exchange.ClusterID.Equals(clusterId)).With(
		db.Exchange.Lockers.Fetch(),
	).Exec(c)
	if err != nil { 
		log.WithContext(c).WithError(err).WithField("clusterId", clusterId).Error("Fail to retrieve exchanges from database")
		c.JSON(http.StatusInternalServerError, "Fail to retrieve exchanges from database")
		return
	}

	resultExchanges := make([]dto.GetExchangeDto,0)

	loop_cluster: for _,exchangeFromCluster := range exchangesFromCluster { 
		for _,exchangeFromDatabase := range exchangesFromDatabase { 
			//if exchanges from cluster exists in database
			if exchangeFromCluster.Name == exchangeFromDatabase.Name { 
				resultExchanges = append(resultExchanges, dto.GetExchangeDto{
					Name: exchangeFromCluster.Name,
					Type: exchangeFromCluster.Type,
					Durable: exchangeFromCluster.Durable,
					Internal: exchangeFromCluster.Internal,
					Arguments: exchangeFromCluster.Arguments,
					ClusterId: exchangeFromDatabase.ClusterID,
					Id: exchangeFromDatabase.ID,
					IsInCluster: true,
					IsInDatabase: true,
					Lockers: exchangeFromDatabase.Lockers(),
				})
				continue loop_cluster
			}
		}

		resultExchanges = append(resultExchanges, dto.GetExchangeDto{ 
				Name: exchangeFromCluster.Name,
				Type: exchangeFromCluster.Type,
				Durable: exchangeFromCluster.Durable,
				Internal: exchangeFromCluster.Internal,
				Arguments: exchangeFromCluster.Arguments,
				ClusterId: 0,
				IsInCluster: true,
				IsInDatabase: false,
		})
	}

	loop_database: for _,exchangeFromDatabase := range exchangesFromDatabase { 
		for _,exchangeInResult := range resultExchanges { 
			if exchangeFromDatabase.Name == exchangeInResult.Name { 
				continue loop_database
			}
		}


		arguments := make(map[string]interface{})
		err = json.Unmarshal(exchangeFromDatabase.Arguments,&arguments)
		if err != nil {
			log.WithContext(c).WithError(err).Error("Fail to unmarshal exchange arguments")
		}
		resultExchanges = append(resultExchanges, dto.GetExchangeDto{ 
			Id: exchangeFromDatabase.ID,
			Name: exchangeFromDatabase.Name,
			Type: exchangeFromDatabase.Type,
			Durable: exchangeFromDatabase.Durable,
			Internal: exchangeFromDatabase.Internal,
			Arguments: arguments,
			ClusterId: exchangeFromDatabase.ClusterID,
			IsInCluster: false,
			IsInDatabase: true,
			Lockers: exchangeFromDatabase.Lockers(),
		})
	}

	c.JSON(http.StatusOK,resultExchanges)

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