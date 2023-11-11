package user

import "github.com/productivityeng/orabbit/src/packages/rabbitmq/common"

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
