package cluster

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core"
	broker_controller "github.com/productivityeng/orabbit/pkg/cluster/resources"
)

var clusterController broker_controller.ClusterController

func Routes(routes *gin.Engine, DependencyLocator *core.DependencyLocator) {
	clusterController = broker_controller.NewClusterController(DependencyLocator)

	brokerResourcePath := "/:clusterId"
	clusterRoutes := routes.Group("/cluster")

	clusterRoutes.GET("/", clusterController.ListClusters)
	clusterRoutes.GET("/exists", clusterController.FindCluster)
	clusterRoutes.GET(brokerResourcePath, clusterController.GetCluster)
	clusterRoutes.PUT(brokerResourcePath, clusterController.UpdateCluster)
	clusterRoutes.POST("/", clusterController.CreateCluster)
	clusterRoutes.DELETE(brokerResourcePath, clusterController.DeleteCluster)
}
