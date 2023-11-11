package user

import (
	"github.com/gin-gonic/gin"
	repository2 "github.com/productivityeng/orabbit/cluster/repository"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/user"
	"github.com/productivityeng/orabbit/user/controllers"
	"github.com/productivityeng/orabbit/user/repository"
	"gorm.io/gorm"
)

var userController controllers.UserController
var userRepository repository.UserRepository
var brokerRepository repository2.ClusterRepositoryInterface
var userManagement user.UserManagement

func Routes(routes *gin.Engine, db *gorm.DB) *gin.RouterGroup {
	brokerRepository = repository2.NewClusterMysqlRepositoryImpl(db)
	userRepository = repository.NewUserRepositoryMySql(db)
	userManagement = user.NewUserManagement()

	userController = controllers.NewUserController(userRepository, brokerRepository, userManagement)

	userRouter := routes.Group("/:clusterId/user")
	userRouter.GET("/", userController.ListUsersFromCluster)
	userRouter.POST("/", userController.CreateUser)
	userRouter.DELETE("/:userId", userController.DeleteUser)
	userRouter.GET("/:userId", userController.FindUser)
	userRouter.GET("/find", userController.FindUser)
	userRouter.POST("/syncronize", userController.SyncronizeUser)

	return userRouter
}
