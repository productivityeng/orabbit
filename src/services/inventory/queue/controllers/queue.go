package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/cluster/repository"
	queue_repository "github.com/productivityeng/orabbit/queue/repository"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/queue"
)

func NewQueueController(repositoryInterface repository.ClusterRepositoryInterface, management queue.QueueManagement,
	queueRepository queue_repository.QueueRepository) QueueControllerImpl {
	return QueueControllerImpl{ClusterRepository: repositoryInterface, QueueManagement: management,
		QueueRepository: queueRepository}
}

type QueueController interface {
	ListQueuesFromCluster(c *gin.Context)
	ImportQueueFromCluster(c *gin.Context)
}

type QueueControllerImpl struct {
	ClusterRepository repository.ClusterRepositoryInterface
	QueueManagement   queue.QueueManagement
	QueueRepository   queue_repository.QueueRepository
}
