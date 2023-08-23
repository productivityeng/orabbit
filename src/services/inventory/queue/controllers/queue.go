package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/cluster/repository"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/queue"
)

func NewQueueController(repositoryInterface repository.ClusterRepositoryInterface, management queue.QueueManagement) QueueControllerImpl {
	return QueueControllerImpl{ClusterRepository: repositoryInterface, QueueManagement: management}
}

type QueueController interface {
	ListQueuesFromCluster(c *gin.Context)
	ImportQueueFromCluster(c *gin.Context)
}

type QueueControllerImpl struct {
	ClusterRepository repository.ClusterRepositoryInterface
	QueueManagement   queue.QueueManagement
}
