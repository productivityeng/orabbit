package contracts




type GetUserHashRequest struct {
	RabbitAccess

	UserToRetrieveHash string
}

type CreateNewUserRequest struct {
	RabbitAccess
	UserToCreate            string
	PasswordOfUsertToCreate string
}

type CreateNewUserWithHashPasswordRequest struct {
	RabbitAccess
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
	RabbitAccess
}

type DeleteUserRequest struct {
	RabbitAccess
	Username string
}
