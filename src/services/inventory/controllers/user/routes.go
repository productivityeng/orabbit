package user

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/controllers/user/resources"
	"github.com/productivityeng/orabbit/core/core"
	"github.com/productivityeng/orabbit/pkg/rabbitmq/user"
)

var userController resources.UserController
var userManagement contracts.UserManagement

func Routes(routes *gin.Engine, DependencyLocator *core.DependencyLocator) *gin.RouterGroup {
	userManagement = user.NewUserManagement()
	userController = resources.NewUserController(DependencyLocator)

	userRouter := routes.Group("/:clusterId/user")
	userRouter.GET("/", userController.ListUsersFromCluster)
	userRouter.POST("/", userController.CreateUser)
	userRouter.DELETE("/:userId", userController.DeleteUser)
	userRouter.GET("/:userId", userController.FindUser)
	userRouter.GET("/find", userController.FindUser)
	userRouter.POST("/:userId/syncronize", userController.SyncronizeUser)

	return userRouter
}
