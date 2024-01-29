package resources

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// RemoveQueueFromCluster
// @Summary Retriave a queue from cluster
// @Schemes
// @Description Retrieve a queue from cluster
// @Tags Queue
// @Accept json
// @Produce json
// @Param clusterId path int true "Cluster id"
// @Param queueId path int true "Queue id"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /{clusterId}/queue/{queueId} [get]
func (q QueueControllerImpl) FindQueue(c *gin.Context) {
	_,queueId,err := q.parseFindQueueFromClusterParam(c)
	if err != nil { 
		return
	}

	queueFromDb,err := q.getQueueById(c,queueId)
	if err != nil { 
		return
	}

	c.JSON(http.StatusOK,queueFromDb)

}



func (controller QueueControllerImpl) parseFindQueueFromClusterParam(c *gin.Context)(clusterId int,queueId int, err error){
	clusterIdParam := c.Param("clusterId")
	clusterIdConv, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse clusterId Param")
		c.JSON(http.StatusBadRequest, "Error parsing clusterId from url route")
		return 0,0,err
	}


	queueIdParam := c.Param("queueId")
	queueIdConv, err := strconv.ParseInt(queueIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("queueId", queueIdParam).Error("Fail to parse queueId Param")
		c.JSON(http.StatusBadRequest, "Error parsing queueId from url route")
		 return 0,0,err
	}

	return int(clusterIdConv),int(queueIdConv),nil
}