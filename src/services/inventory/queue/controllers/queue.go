package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core/context"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/queue"
)

func NewQueueController(DependencyLocator *context.DependencyLocator, management queue.QueueManagement,
	) QueueControllerImpl {
	return QueueControllerImpl{ QueueManagement: management, DependencyLocator: DependencyLocator}
}

type QueueController interface {
	ListQueuesFromCluster(c *gin.Context)
	ImportQueueFromCluster(c *gin.Context)
	SyncronizeQueue(c *gin.Context)
	RemoveQueueFromCluster(c *gin.Context)
}

type QueueControllerImpl struct {
	DependencyLocator *context.DependencyLocator
	QueueManagement   queue.QueueManagement
}
