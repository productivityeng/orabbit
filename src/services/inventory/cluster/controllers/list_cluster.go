package cluster

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/rabbitmq/common"
	log "github.com/sirupsen/logrus"
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

	params,err  := ctrl.parsePageParams(c)
	log.WithField("params", params).Info("Received request")
	if err != nil { return }

	log.WithField("params", params).Info("Retrieving clusters")
	paginatedClusters, err := ctrl.DependencyLocator.PrismaClient.Cluster.FindMany().Take(params.PageSize).Skip(params.PageNumber-1).Exec(c)
	
	if err != nil {
		log.WithError(err).Error("Error retrieving clusters from repository")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	log.WithField("paginatedClusters", paginatedClusters).Info("Retrieved clusters")
	c.JSON(http.StatusOK, paginatedClusters)
}

func (ctrl *clusterControllerDefaultImp) parsePageParams(c *gin.Context) (params *common.PageParam, err error) {
	err = c.BindQuery(&params)
	if err != nil {
		log.WithError(err).Error("Error trying to parse query params")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	return params, nil
 }

