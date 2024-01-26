package queue

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core/core"
	queue_controller "github.com/productivityeng/orabbit/queue/controllers"
	"github.com/productivityeng/orabbit/rabbitmq/queue"
)

var queueController queue_controller.QueueController

func Routes(routes *gin.Engine, dependencyLocator *core.DependencyLocator) *gin.RouterGroup {

	queueController = queue_controller.NewQueueController(dependencyLocator,queue.NewQueueManagement())

	userRouter := routes.Group("/:clusterId/queue")
	userRouter.GET("/queuesfromcluster", queueController.ListQueuesFromCluster)
	userRouter.POST("/import", queueController.ImportQueueFromCluster)
	userRouter.POST("/syncronize", queueController.SyncronizeQueue)
	userRouter.DELETE("/remove", queueController.RemoveQueueFromCluster)
	userRouter.GET("/:queueId", queueController.FindQueue)

	return userRouter
}
