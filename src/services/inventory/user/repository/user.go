package repository

import (
	"context"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/user/dto"
	userEntities "github.com/productivityeng/orabbit/user/entities"
)

type UserRepository interface {
	///CreateUser store a new broker in storage with provided parameter
	CreateUser(broker *userEntities.UserEntity) (*userEntities.UserEntity, error)
	//ListUsers retrieve a lista of broker with paginated options
	ListUsers(brokerId int32, pageSize int, pageNumber int, ctx context.Context) (*contracts.PaginatedResult[dto.GetUserResponse], error)
	//DeleteUser soft delete a broker with a provided brokerId
	DeleteUser(userId int32, ctx context.Context) error
	//GetUser retrieve a broker with a provided brokerId
	GetUser(clusterId int32, userId int32, ctx context.Context) (*userEntities.UserEntity, error)

	CheckIfUserExistsForCluster(brokerId int32, username string, ctx context.Context) (bool, error)
}
