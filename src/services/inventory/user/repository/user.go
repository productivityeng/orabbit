package repository

import (
	"context"

	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/user/dto"
	userEntities "github.com/productivityeng/orabbit/user/entities"
)

type UserRepository interface {
	// CreateUser store a new broker in storage with provided parameter
	CreateUser(broker *userEntities.UserEntity) (*userEntities.UserEntity, error)
	//ListUsers retrieve a lista of broker with paginated options
	ListUsers(clusterId uint, pageSize int, pageNumber int, ctx context.Context) (*contracts.PaginatedResult[dto.GetUserResponse], error)
	//DeleteUser soft delete a broker with a provided brokerId
	DeleteUser(clusterId uint, userId uint, ctx context.Context) error
	//GetUser retrieve a broker with a provided brokerId
	GetUser(clusterId uint, userId uint, ctx context.Context) (*userEntities.UserEntity, error)
	GetUserLock(userId int64, ctx context.Context) (*dto.LockUserDto, error)

	CheckIfUserExistsForCluster(brokerId uint, username string, ctx context.Context) (bool, error)
	ListAllRegisteredUsers(clusterId uint, ctx context.Context) (userEntities.UserEntityList, error)
}
