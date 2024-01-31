package locker

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core"
	"github.com/productivityeng/orabbit/pkg/locker/resources"
)


var LockerController *resources.LockerController


func Routes(routes *gin.Engine, dependencyLocator *core.DependencyLocator) *gin.RouterGroup {

	LockerController = resources.NewLockerController(dependencyLocator)
	
	userRouter := routes.Group("/:clusterId/locker/:lockerType/:lockerId")
	userRouter.GET("/", LockerController.FindLocker)
	userRouter.POST("/", LockerController.CreateLocker)
	userRouter.DELETE("/", LockerController.DisableLocker)

	return userRouter
}
