package cluster

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/src/packages/common"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// ListClusters
// @Summary Retrieve a list of rabbitmq clusters registered
// @Schemes
// @Description Retrieve a paginated list of cluster that the user has access
// @Tags Cluster
// @Accept json
// @Produce json
// @Success 200
// @Param params query common.PageParam true "Number of items in one page"
// @Router /cluster [get]
func (ctrl *clusterControllerDefaultImp) ListClusters(c *gin.Context) {

	var param common.PageParam
	err := c.BindQuery(&param)
	if err != nil {
		log.WithError(err).Error("Error trying to parse query params")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paginatedClusters, err := ctrl.ClusterRepository.ListCluster(param.PageSize, param.PageNumber)
	if err != nil {
		log.WithError(err).Error("Error retrieving clusters from repository")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, paginatedClusters)

}
