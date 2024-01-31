package resources

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/db"
	log "github.com/sirupsen/logrus"
)

// DeleteCluster
// @Summary Soft delete a cluster
// @Schemes
// @Description Soft delete a cluster will not completly erase from database, but will not show up anymore in the
// system. All queues,bindings,shovels and related artifacts will be soft delete to
// @Tags Cluster
// @Accept json
// @Produce json
// @Success 204 {object} bool
// @Param clusterId path int true "Id of a cluster to be soft deleted"
// @Router /cluster/{clusterId} [delete]
func (ctrl *clusterControllerDefaultImp) DeleteCluster(c *gin.Context) {
	clusterId, err := ctrl.parseGetClusterParams(c)

	if err != nil { return }

	log.WithField("clusterId", clusterId).Info("Cluster will be deleted")

	_, err = ctrl.DependencyLocator.PrismaClient.Cluster.FindUnique(db.Cluster.ID.Equals(clusterId)).Delete().Exec(c)
	if errors.Is(err, db.ErrNotFound) { 
		c.JSON(http.StatusNotFound, false)
		return 
	}else if err != nil {
		errorMsg := "Fail to delete clusterId"
		log.WithError(err).WithField("clusterId", clusterId).Error(errorMsg)
		c.JSON(http.StatusInternalServerError, errorMsg)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": fmt.Sprintf("cluster with id %d successufly deleted", clusterId)})
}


