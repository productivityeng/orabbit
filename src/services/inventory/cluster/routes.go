package cluster

import (
	"github.com/gin-gonic/gin"
	broker_controller "github.com/productivityeng/orabbit/cluster/controllers"
	"github.com/productivityeng/orabbit/cluster/repository"
	"github.com/productivityeng/orabbit/core/context"
	"github.com/productivityeng/orabbit/core/validators"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq"
)

var clusterController broker_controller.ClusterController
var brokerRepository repository.ClusterRepositoryInterface
var clusterValidator validators.ClusterValidator
var overviewManagement rabbitmq.OverviewManagement

func Routes(routes *gin.Engine, DependencyLocator *context.DependencyLocator) {
	overviewManagement = rabbitmq.NewOverviewManagementImpl()
	clusterValidator = validators.NewClusterValidatorDefault(brokerRepository, overviewManagement)
	clusterController = broker_controller.NewClusterController(brokerRepository, clusterValidator)

	brokerResourcePath := "/:clusterId"
	clusterRoutes := routes.Group("/cluster")

	clusterRoutes.GET("/", clusterController.ListClusters)
	clusterRoutes.GET("/exists", clusterController.FindCluster)
	clusterRoutes.GET(brokerResourcePath, clusterController.GetCluster)
	clusterRoutes.PUT(brokerResourcePath, clusterController.UpdateCluster)
	clusterRoutes.POST("/", clusterController.CreateCluster)
	clusterRoutes.DELETE(brokerResourcePath, clusterController.DeleteCluster)
}
