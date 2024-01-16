package queue

import (
	"github.com/gin-gonic/gin"
	repository2 "github.com/productivityeng/orabbit/cluster/repository"
	"github.com/productivityeng/orabbit/core/context"
	queue_controller "github.com/productivityeng/orabbit/queue/controllers"
	"github.com/productivityeng/orabbit/queue/repository"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/queue"
)

var queueController queue_controller.QueueController
var clusterRepository repository2.ClusterRepositoryInterface
var queueRepository repository.QueueRepository

func Routes(routes *gin.Engine, dependencyLocator *context.DependencyLocator) *gin.RouterGroup {

	queueController = queue_controller.NewQueueController(dependencyLocator,queue.NewQueueManagement())

	userRouter := routes.Group("/:clusterId/queue")
	userRouter.GET("/queuesfromcluster", queueController.ListQueuesFromCluster)
	userRouter.POST("/import", queueController.ImportQueueFromCluster)
	userRouter.POST("/syncronize", queueController.SyncronizeQueue)
	userRouter.DELETE("/remove", queueController.RemoveQueueFromCluster)

	return userRouter
}
