package user

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/user/controllers"
	"github.com/productivityeng/orabbit/user/repository"
	"gorm.io/gorm"
)

var userController controllers.UserController
var userRepository repository.UserRepository

func Routes(routes *gin.Engine, db *gorm.DB) *gin.RouterGroup {

	userRepository = repository.NewUserRepositoryMySql(db)
	userController = controllers.NewUserController(userRepository)

	userRouter := routes.Group("/user")
	userRouter.GET("/", userController.ListUsers)
	userRouter.POST("/", userController.CreateUser)
	userRouter.DELETE("/:userId", userController.DeleteUser)
	userRouter.GET("/:userId", userController.GetUser)
	userRouter.GET("/find", userController.FindUser)

	return userRouter
}
