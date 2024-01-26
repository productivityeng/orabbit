package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/cluster/models"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/queue/dto"
	"github.com/productivityeng/orabbit/rabbitmq/queue"
	log "github.com/sirupsen/logrus"
)

// RemoveQueueFromCluster
// @Summary Remove a fila do cluster
// @Schemes
// @Description Remove a fila do cluster mas nao altera o cadastro no ostern
// @Tags Queue
// @Accept json
// @Produce json
// @Param clusterId path int true "Cluster id"
// @Success 200
// @Failure 404
// @Failure 500
// @Param QueueImportRequest body dto.QueueRemoveRequest true "Request"
// @Router /{clusterId}/queue/remove [delete]
func (q QueueControllerImpl) RemoveQueueFromCluster(c *gin.Context) {
	clusterId,queueRemoveRequest,err := q.parseRemoveQueueFromClusterParam(c)
	if err != nil { 
		return
	}

	queueFromDb,err := q.getQueueById(c,queueRemoveRequest.QueueId)
	if err != nil {
		return
	 }

	cluster,err := q.getClusterByid(c,clusterId)

	if err != nil {
		return
	}

	err = q.verifyIfQueueIsLocked(queueFromDb.ID,c)
	if err != nil {
		return
	
	}

	err = q.deleteQueue(c,cluster,queueFromDb)
	if err != nil { return }

	

	c.JSON(http.StatusOK, gin.H{"info": "Fila removida do  cluster"})

}

func (controller QueueControllerImpl) deleteQueue(c *gin.Context,cluster *db.ClusterModel,queueFromCluster *db.QueueModel) error {
	deleteQueueRequest := queue.DeleteQueueRequest{
		RabbitAccess: models.GetRabbitMqAccess(cluster),
		Queue:        queueFromCluster.Name,
	}

	err := controller.QueueManagement.DeleteQueue(deleteQueueRequest)

	if err != nil {
		log.WithError(err).WithField("request", deleteQueueRequest).Error("Erro ao tentar deletar a fila no cluster")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	return nil
}

func (controller QueueControllerImpl) parseRemoveQueueFromClusterParam(c *gin.Context)(clusterId int,request *dto.QueueRemoveRequest, err error){
	clusterIdParam := c.Param("clusterId")
	clusterIdConv, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse clusterId Param")
		c.JSON(http.StatusBadRequest, "Error parsing clusterId from url route")
		return 0,nil,err
	}

	var queueRemoveRequest dto.QueueRemoveRequest

	err = c.BindJSON(&queueRemoveRequest)
	if err != nil {
		log.WithContext(c).WithError(err).Error("Falha ao interpretar a requisicao")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0,nil,err
	}
	return int(clusterIdConv),&queueRemoveRequest,nil
}