package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/db"
	log "github.com/sirupsen/logrus"
)


func (ctrl VirtualHostControllerImpl) parseClusterIdParams(c *gin.Context) (int, error) {
	clusterIdParam := c.Param("clusterId")
	clusterId, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse brokerId Param")
		c.JSON(http.StatusBadRequest, "Error parsing brokerId from url route")
		return 0, err
	}
	return int(clusterId), nil
}



func (ctrl VirtualHostControllerImpl) parseVirtualHostIdParams(c *gin.Context) (int, error) {
	virtualHostIdParam := c.Param("virtualHostId")
	virtualHostId, err := strconv.ParseInt(virtualHostIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("clusterId", virtualHostIdParam).Error("Fail to parse virtualHostId Param")
		c.JSON(http.StatusBadRequest, "Error parsing virtualHostId from url route")
		return 0, err
	}
	return int(virtualHostId), nil
}

func (ctrl VirtualHostControllerImpl) getClusterById(clusterId int, c *gin.Context) (*db.ClusterModel, error) {
	cluster,err := ctrl.DependencyLocator.PrismaClient.Cluster.FindUnique(db.Cluster.ID.Equals(clusterId)).Exec(c)
	if errors.Is(err, db.ErrNotFound) {
		log.WithError(err).WithField("clusterId", clusterId).Error("Cluster not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Cluster not found"})
		return nil, err
	}else if err != nil {
		log.WithError(err).WithField("clusterId", clusterId).Error("Fail to retrieve cluster")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}
	return cluster, nil
}