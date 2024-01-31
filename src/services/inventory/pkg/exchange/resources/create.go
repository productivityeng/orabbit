package resources

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/pkg/exchange/dto"
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
// @Param CreateExchangeDto body dto.CreateExchangeDto true "Request"
// @Router /{clusterId}/exchange  [post]
func (ctrl *ExchangeController) CreateExchange(c *gin.Context)  {
	
	log.WithContext(c).Info("Parsing clusterId param")
	clusterId,err := ctrl.parseClusterIdParam(c)
	if err != nil { return}

	log.WithContext(c).Info("Parsing body")
	requestBody ,err := ctrl.parseCreateExchangeBody(c)
	if err != nil { return}

	cluster,err := ctrl.getCluster(clusterId,c)
	if err != nil { return }

	err = ctrl.createExchangeInCluster(cluster,requestBody,c)
	if err != nil { return }



	exchangeSaved,err := ctrl.saveExchangeInDatabase(requestBody,clusterId,c)
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


