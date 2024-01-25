package dto

type ImportOrCreateUserRequest struct {
	ClusterId int   `json:"ClusterId" binding:"required"`
	Username  string `json:"Username" binding:"required"`
	Password  string `json:"Password" `
	Create    bool   `json:"Create" `
}

