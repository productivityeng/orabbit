package cluster

import (
	"github.com/gin-gonic/gin"
	broker_controller "github.com/productivityeng/orabbit/cluster/controllers"
	"github.com/productivityeng/orabbit/core/core"
	"github.com/productivityeng/orabbit/core/validators"
)

var clusterController broker_controller.ClusterController
var clusterValidator validators.ClusterValidator

func Routes(routes *gin.Engine, DependencyLocator *core.DependencyLocator) {
	clusterValidator = validators.NewClusterValidatorDefault(DependencyLocator)
	clusterController = broker_controller.NewClusterController(DependencyLocator, clusterValidator)

	brokerResourcePath := "/:clusterId"
	clusterRoutes := routes.Group("/cluster")

	clusterRoutes.GET("/", clusterController.ListClusters)
	clusterRoutes.GET("/exists", clusterController.FindCluster)
	clusterRoutes.GET(brokerResourcePath, clusterController.GetCluster)
	clusterRoutes.PUT(brokerResourcePath, clusterController.UpdateCluster)
	clusterRoutes.POST("/", clusterController.CreateCluster)
	clusterRoutes.DELETE(brokerResourcePath, clusterController.DeleteCluster)
}
