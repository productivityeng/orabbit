package locker

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core/core"
	locker "github.com/productivityeng/orabbit/locker/controllers"
)


var LockerController *locker.LockerController


func Routes(routes *gin.Engine, dependencyLocator *core.DependencyLocator) *gin.RouterGroup {

	LockerController = locker.NewLockerController(dependencyLocator)
	
	userRouter := routes.Group("/:clusterId/locker/:lockerType/:lockerId")
	userRouter.GET("/", LockerController.FindLocker)
	userRouter.POST("/", LockerController.CreateLocker)
	userRouter.DELETE("/", LockerController.DisableLocker)

	return userRouter
}
