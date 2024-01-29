package resources

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/db"
	log "github.com/sirupsen/logrus"
)


type FindClusterByHost struct {
	Host string `json:"Host" binding:"required"`
	Port int  `json:"Port" binding:"required"`
}

// FindCluster
// @Summary Verify if exists a rabbitmqcluster
// @Schemes
// @Description Check if exists an rabbitmq cluster with host es
// @Tags Cluster
// @Accept json
// @Produce json
// @Success 200
// @Param params query FindClusterByHost true "Number of items in one page"
// @Router /cluster/exists [get]
func (ctrl *clusterControllerDefaultImp) FindCluster(c *gin.Context) {
	
	params,err := ctrl.parseFindCluster(c)
	if err != nil { return }

	unique := db.Cluster.UniqueNameHostPort(
		db.Cluster.Host.Equals(params.Host),
		db.Cluster.Port.Equals(params.Port),
	)
	cluster,err := ctrl.DependencyLocator.PrismaClient.Cluster.FindUnique(unique).Exec(c)
	if errors.Is(err,db.ErrNotFound) { 
		log.WithError(err).WithField("params", params).Info("Cluster not found")
		c.JSON(http.StatusNotFound, false)
		return 
	} else if err != nil { 
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cluster)

}

func (ctrl *clusterControllerDefaultImp) UpdateCluster(c *gin.Context) {

}

func (ctrl *clusterControllerDefaultImp) parseFindCluster(c *gin.Context) (params *FindClusterByHost, err error) {
	var param FindClusterByHost
	err = c.BindQuery(&param)
	if err != nil {
		log.WithError(err).Error("Error trying to parse query params")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil,err
	}
	return &param,nil
}