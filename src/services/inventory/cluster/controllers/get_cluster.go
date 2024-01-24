package cluster

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// GetCluster
// @Summary Retrieve a single rabbitmq cluster
// @Schemes
// @Description Retrieve a single rabbitmq cluster
// @Tags Cluster
// @Accept json
// @Produce json
// @Success 200 {object} db.ClusterModel
// @NotFound 404 {object} bool
// @Param clusterId path int true "Id of a cluster to be retrived"
// @Router /cluster/{clusterId} [get]
func (ctrl *clusterControllerDefaultImp) GetCluster(c *gin.Context) {
	clusterId, err := ctrl.parseGetClusterParams(c)
	if err != nil { return }

	log.WithField("clusterId", clusterId).Info("Cluster will be deleted")
	cluster, err := ctrl.getClusterByid(c,clusterId)
	if err != nil { return }

	c.JSON(http.StatusOK, cluster)
	return
}

