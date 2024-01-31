package queue

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core"
	"github.com/productivityeng/orabbit/pkg/queue/resources"
	"github.com/productivityeng/orabbit/pkg/rabbitmq/queue"
)

var queueController resources.QueueController

func Routes(routes *gin.Engine, dependencyLocator *core.DependencyLocator) *gin.RouterGroup {

	queueController = resources.NewQueueController(dependencyLocator,queue.NewQueueManagement())

	userRouter := routes.Group("/:clusterId/queue")
	userRouter.GET("/queuesfromcluster", queueController.ListQueuesFromCluster)
	userRouter.POST("/import", queueController.ImportQueueFromCluster)
	userRouter.POST("/syncronize", queueController.SyncronizeQueue)
	userRouter.DELETE("/remove", queueController.RemoveQueueFromCluster)
	userRouter.GET("/:queueId", queueController.FindQueue)

	return userRouter
}
