package dto

type ImportOrCreateUserRequest struct {
	ClusterId int32  `json:"ClusterId" binding:"required"`
	Username  string `json:"Username" binding:"required"`
	Password  string `json:"Password" `
	Create    bool   `json:"Create" `
}

type GetUserResponse struct {
	Id           int32  `json:"Id"`
	Username     string `json:"Username"`
	PasswordHash string `json:"PasswordHash"`
	BrokerId     int32  `json:"ClusterId"`
}
