package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core/core"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/user"
)

type UserController interface {
	ListUsersFromCluster(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	FindUser(c *gin.Context)
	SyncronizeUser(c *gin.Context)
}

type UserControllerImpl struct {
	DependencyLocator *core.DependencyLocator 
	UserManagement    user.UserManagement
}

func NewUserController(DependencyLocator *core.DependencyLocator,
	userManagement user.UserManagement) *UserControllerImpl {
	return &UserControllerImpl{DependencyLocator: DependencyLocator,
		UserManagement:    userManagement}
}

func (entity *UserControllerImpl) GetEntity(c *gin.Context) {

}

func (entity *UserControllerImpl) UpdateUser(c *gin.Context) {}
