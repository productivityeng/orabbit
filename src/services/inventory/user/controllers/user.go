package controllers

import (
	"github.com/gin-gonic/gin"
	repository "github.com/productivityeng/orabbit/user/repository"
)

type UserController interface {
	GetUser(c *gin.Context)
	ListUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	FindUser(c *gin.Context)
}

type UserControllerImpl struct {
	UserRepository repository.UserRepository
}

func NewUserController(userRepository repository.UserRepository) *UserControllerImpl {
	return &UserControllerImpl{UserRepository: userRepository}
}

func (entity *UserControllerImpl) GetEntity(c *gin.Context) {

}
func (entity *UserControllerImpl) ListUsers(c *gin.Context)  {}
func (entity *UserControllerImpl) CreateUser(c *gin.Context) {}
func (entity *UserControllerImpl) UpdateUser(c *gin.Context) {}
func (entity *UserControllerImpl) DeleteUser(c *gin.Context) {}
func (entity *UserControllerImpl) FindUser(c *gin.Context)   {}
func (entity *UserControllerImpl) GetUser(c *gin.Context)    {}
