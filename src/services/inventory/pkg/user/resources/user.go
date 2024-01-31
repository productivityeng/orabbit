package resources

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core"
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
}

func NewUserController(DependencyLocator *core.DependencyLocator) *UserControllerImpl {
	return &UserControllerImpl{DependencyLocator: DependencyLocator}
}

func (entity *UserControllerImpl) GetEntity(c *gin.Context) {

}

func (entity *UserControllerImpl) UpdateUser(c *gin.Context) {}
