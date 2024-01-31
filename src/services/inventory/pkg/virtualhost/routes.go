package virtualhost

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core"
	"github.com/productivityeng/orabbit/pkg/virtualhost/resources"
)

var virtualHostController resources.VirtualHostController

func Routes(routes *gin.Engine, DependencyLocator *core.DependencyLocator) *gin.RouterGroup {
	virtualHostController = resources.NewVirtualHostControllerImpl(DependencyLocator)

	userRouter := routes.Group("/:clusterId/virtualhost")
	userRouter.GET("/", virtualHostController.ListVirtualHost)
	userRouter.POST("/import", virtualHostController.Import)
	userRouter.POST("/:virtualHostId/syncronize", virtualHostController.Syncronize)
	userRouter.DELETE("/:virtualHostId", virtualHostController.DeleteVirtualHost)

	return userRouter
}
