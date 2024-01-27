package virtualhost

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core/core"
	"github.com/productivityeng/orabbit/virtualhost/controllers"
)

var virtualHostController controllers.VirtualHostController

func Routes(routes *gin.Engine, DependencyLocator *core.DependencyLocator) *gin.RouterGroup {

	virtualHostController = controllers.NewVirtualHostControllerImpl(DependencyLocator)

	userRouter := routes.Group("/:clusterId/virtualhost")
	userRouter.GET("/", virtualHostController.ListVirtualHost)
	userRouter.POST("/import", virtualHostController.Import)

	return userRouter
}
