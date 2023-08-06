package user

import (
	"github.com/gin-gonic/gin"
	repository2 "github.com/productivityeng/orabbit/broker/repository"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/user"
	"github.com/productivityeng/orabbit/user/controllers"
	"github.com/productivityeng/orabbit/user/repository"
	"gorm.io/gorm"
)

var userController controllers.UserController
var userRepository repository.UserRepository
var brokerRepository repository2.BrokerRepositoryInterface
var userManagement user.UserManagement

func Routes(routes *gin.Engine, db *gorm.DB) *gin.RouterGroup {
	brokerRepository = repository2.NewBrokerMysqlImpl(db)
	userRepository = repository.NewUserRepositoryMySql(db)
	userManagement = user.NewUserManagement()

	userController = controllers.NewUserController(userRepository, brokerRepository, userManagement)

	userRouter := routes.Group("/user")
	userRouter.GET("/", userController.ListUsers)
	userRouter.POST("/", userController.CreateUser)
	userRouter.DELETE("/:userId", userController.DeleteUser)
	userRouter.GET("/:userId", userController.GetUser)
	userRouter.GET("/find", userController.FindUser)

	return userRouter
}
