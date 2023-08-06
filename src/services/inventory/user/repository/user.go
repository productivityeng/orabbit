package repository

import (
	"context"
	"github.com/productivityeng/orabbit/contracts"
	userEntities "github.com/productivityeng/orabbit/user/entities"
)

type UserRepository interface {
	///CreateUser store a new broker in storage with provided parameter
	CreateUser(broker *userEntities.UserEntity) (*userEntities.UserEntity, error)
	//ListUsers retrieve a lista of broker with paginated options
	ListUsers(pageSize int, pageNumber int) (*contracts.PaginatedResult[userEntities.UserEntity], error)
	//DeleteUser soft delete a broker with a provided brokerId
	DeleteUser(userId int32, ctx context.Context) error
	//GetUser retrieve a broker with a provided brokerId
	GetUser(userId int32, ctx context.Context) (*userEntities.UserEntity, error)
}
