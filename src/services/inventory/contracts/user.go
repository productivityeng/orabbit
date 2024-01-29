package contracts

import (
	"context"

	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	"github.com/productivityeng/orabbit/rabbitmq/common"
)


type UserManagement interface {
	GetUserHash(request GetUserHashRequest, ctx context.Context) (string, error)
	GetUserByName(request GetUserByNameRequest, ctx context.Context) (*rabbithole.UserInfo, error)
	CreateNewUser(request CreateNewUserRequest, ctx context.Context) (*CreateNewUserResult, error)
	CreateNewUserWithHashPassword(request CreateNewUserWithHashPasswordRequest, ctx context.Context) (*CreateNewUserResult, error)
	ListAllUser(request ListAllUsersRequest) ([]rabbithole.UserInfo, error)
	DeleteUser(request DeleteUserRequest, ctx context.Context) error
}


type GetUserHashRequest struct {
	common.RabbitAccess

	UserToRetrieveHash string
}

type CreateNewUserRequest struct {
	common.RabbitAccess
	UserToCreate            string
	PasswordOfUsertToCreate string
}

type CreateNewUserWithHashPasswordRequest struct {
	common.RabbitAccess
	UsernameToCreate string
	PasswordHash     string
}

type CreateNewUserResult struct {
	PasswordHash string
}

type ListUserResult struct {
	PasswordHash string
	Name         string
}

type ListAllUsersRequest struct {
	common.RabbitAccess
}

type DeleteUserRequest struct {
	common.RabbitAccess
	Username string
}
