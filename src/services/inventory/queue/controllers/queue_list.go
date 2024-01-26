package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/queue/dto"
	"github.com/productivityeng/orabbit/rabbitmq/common"
	"github.com/productivityeng/orabbit/rabbitmq/queue"
	log "github.com/sirupsen/logrus"
)

// ListQueuesFromCluster
// @Summary Retrieve all users from rabbitmq cluster
// @Schemes
// @Description Retrieve all users that exist on rabbit cluster. Event if it its registered in ostern
// @Tags Queue
// @Accept json
// @Produce json
// @Param clusterId path int true "Cluster id from where retrieve users"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /{clusterId}/queue/queuesfromcluster [get]
func (q QueueControllerImpl) ListQueuesFromCluster(c *gin.Context) {
	
	clusterId, err := q.parseListQueuesFromClusterParams(c)
	if err != nil { 
		return
	}
	log.WithField("clusterId", clusterId).Info("ListQueuesFromCluster")
	
	cluster,err := q.getClusterByid(c,clusterId)
	if err != nil { return }

	log.WithField("cluster", cluster).Info("ListQueuesFromCluster")

	queuesFromCluster, err := q.QueueManagement.GetAllQueuesFromCluster(queue.ListQueuesRequest{
		RabbitAccess: common.RabbitAccess{
			Host:     cluster.Host,
			Port:     cluster.Port,
			Username: cluster.User,
			Password: cluster.Password,
		},
	})

	if err != nil {
		log.WithError(err).Error("Fail to retrieve all queuesFromCluster from cluster")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.WithField("qtdQueues", len(queuesFromCluster)).Info("queues retrived from the cluster ")

	log.WithField("clusterId", clusterId).Info("getting queues from database")

	queuesFromDatabase,err := q.getAllQueuesFromDatabase(clusterId,c)

	if err != nil {
		return
	 }

	log.WithField("qtdQueues", len(queuesFromDatabase)).Info("queues retrived from the database ")

	getQueueResponse := make(dto.GetQueueResponseList, 0)

	for _, queueFromCluster := range queuesFromCluster {

		var queueFromClusterExistsInDatabase = false
		var queueIdFromDatabase = int(0)
		var locker []db.LockerQueueModel
		if queueFromBd := queuesFromDatabase.GetQueueFromListByName(queueFromCluster.Name); queueFromBd != nil {
			queueFromClusterExistsInDatabase = true
			queueIdFromDatabase = queueFromBd.ID
			locker = queueFromBd.LockerQueues()
			
		}

		getQueueResponse = append(getQueueResponse, dto.GetQueueResponse{
			ID:           queueIdFromDatabase,
			ClusterID:    clusterId,
			Name:         queueFromCluster.Name,
			VHost:        queueFromCluster.Vhost,
			Type:         queueFromCluster.Type,
			IsInCluster:  true,
			IsInDatabase: queueFromClusterExistsInDatabase,
			Arguments:    queueFromCluster.Arguments,
			Durable:      queueFromCluster.Durable,
			Lockers: 	  locker,
		})

	}

	for _, queueFromDb := range queuesFromDatabase {
		if queueFromResponse := getQueueResponse.GetByName(queueFromDb.Name); queueFromResponse == nil {

			getQueueResponse = append(getQueueResponse, dto.GetQueueResponse{
				ID:           queueFromDb.ID,
				ClusterID:    clusterId,
				Name:         queueFromDb.Name,
				VHost:        "/",
				Type:         queueFromDb.Type.String(),
				IsInCluster:  false,
				IsInDatabase: true,
				Durable:      queueFromDb.Durable,
				Lockers: 	  queueFromDb.LockerQueues(),
			})
			json.Unmarshal(queueFromDb.Arguments, &getQueueResponse[len(getQueueResponse)-1].Arguments)
		}
	}

	c.JSON(http.StatusOK, getQueueResponse)
}
func (controller *QueueControllerImpl) parseListQueuesFromClusterParams(c *gin.Context) (clusterId int, err error) {
	clusterIdParam := c.Param("clusterId")
	clusterId, err = strconv.Atoi(clusterIdParam)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse clusterId Param")
		c.JSON(http.StatusBadRequest, "Error parsing clusterId from url route")
		return 0, err
	}
	return clusterId, nil
}

func (controller *QueueControllerImpl) getAllQueuesFromDatabase(clusterId int, c *gin.Context) (queuesFromDatabase db.QueueList, err error) {
	queuesFromDatabase, err = controller.DependencyLocator.PrismaClient.Queue.FindMany(db.Queue.ClusterID.Equals(clusterId)).With(
		db.Queue.LockerQueues.Fetch(),
	).Exec(c)
	if err != nil {
		log.WithContext(c).WithError(err).Error("Fail to retrieve all queues from database")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}
	return queuesFromDatabase, nil
}