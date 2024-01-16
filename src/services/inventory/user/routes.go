package user

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core/context"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/user"
	"github.com/productivityeng/orabbit/user/controllers"
)

var userController controllers.UserController
var userManagement user.UserManagement

func Routes(routes *gin.Engine, DependencyLocator *context.DependencyLocator) *gin.RouterGroup {
	userManagement = user.NewUserManagement()
	userController = controllers.NewUserController(DependencyLocator, userManagement)

	userRouter := routes.Group("/:clusterId/user")
	userRouter.GET("/", userController.ListUsersFromCluster)
	userRouter.POST("/", userController.CreateUser)
	userRouter.POST("/:userId/lock", userController.LockUser)
	userRouter.DELETE("/:userId", userController.DeleteUser)
	userRouter.GET("/:userId", userController.FindUser)
	userRouter.GET("/find", userController.FindUser)
	userRouter.POST("/syncronize", userController.SyncronizeUser)

	return userRouter
}
