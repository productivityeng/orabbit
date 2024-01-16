package cluster

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/db"
	log "github.com/sirupsen/logrus"
)


func (controller clusterControllerDefaultImp) parseGetClusterParams(c *gin.Context)(clusterId int, err error) { 
	clusterIdParam := c.Param("clusterId")
	clusterIdConv, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse clusterId Param")
		c.JSON(http.StatusBadRequest, "Error parsing clusterId from url route")
		return 0,err
	}
	return int(clusterIdConv),nil
}

func(controller clusterControllerDefaultImp) getClusterByid(c *gin.Context,clusterId int) (cluster *db.ClusterModel,err error) { 
	cluster,err = controller.DependencyLocator.PrismaClient.Cluster.FindUnique(db.Cluster.ID.Equals(clusterId)).Exec(c)
	if err != nil {
		errorMsg := "Fail to retrive Cluster with clusterId"
		log.WithError(err).WithField("clusterId", clusterId).Error(errorMsg)
		c.JSON(http.StatusInternalServerError, errorMsg)
		return nil,err
	}
	return cluster,nil
}
