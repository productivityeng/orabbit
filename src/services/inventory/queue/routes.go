package queue

import (
	"github.com/gin-gonic/gin"
	repository2 "github.com/productivityeng/orabbit/cluster/repository"
	queue_controller "github.com/productivityeng/orabbit/queue/controllers"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/queue"
	"gorm.io/gorm"
)

var queueController queue_controller.QueueController
var clusterRepository repository2.ClusterRepositoryInterface

func Routes(routes *gin.Engine, db *gorm.DB) *gin.RouterGroup {
	clusterRepository = repository2.NewClusterMysqlRepositoryImpl(db)

	queueController = queue_controller.NewQueueController(clusterRepository, queue.NewQueueManagement())
	userRouter := routes.Group("/:clusterId/queue")
	userRouter.GET("/queuesfromcluster", queueController.ListQueuesFromCluster)

	return userRouter
}
