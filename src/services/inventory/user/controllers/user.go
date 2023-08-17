package controllers

import (
	"github.com/gin-gonic/gin"
	repository2 "github.com/productivityeng/orabbit/cluster/repository"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/user"
	repository "github.com/productivityeng/orabbit/user/repository"
)

type UserController interface {
	ListUsers(c *gin.Context)
	ListUsersFromCluster(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	FindUser(c *gin.Context)
}

type UserControllerImpl struct {
	UserRepository    repository.UserRepository
	ClusterRepository repository2.ClusterRepositoryInterface
	UserManagement    user.UserManagement
}

func NewUserController(userRepository repository.UserRepository, BrokerRepository repository2.ClusterRepositoryInterface,
	userManagement user.UserManagement) *UserControllerImpl {
	return &UserControllerImpl{UserRepository: userRepository,
		ClusterRepository: BrokerRepository,
		UserManagement:    userManagement}
}

func (entity *UserControllerImpl) GetEntity(c *gin.Context) {

}

func (entity *UserControllerImpl) UpdateUser(c *gin.Context) {}

// PingExample godoc
// @Summary Delete a mirror user
// @Schemes
// @Description Delete a mirrored user from the registry, the user will not be deleted from the cluster
// @Tags User
// @Accept json
// @Produce json
// @Param userId path int true "User id registered"
// @Success 204
// @Failure 404
// @Failure 500
// @Router /{clusterId}/user/{userId} [delete]
func (entity *UserControllerImpl) DeleteUser(c *gin.Context) {}
