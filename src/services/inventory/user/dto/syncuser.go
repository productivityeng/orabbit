package dto

type ImportOrCreateUserRequest struct {
	ClusterId uint   `json:"ClusterId" binding:"required"`
	Username  string `json:"Username" binding:"required"`
	Password  string `json:"Password" `
	Create    bool   `json:"Create" `
}

type GetUserResponse struct {
	Id           uint   `json:"Id"`
	Username     string `json:"Username"`
	PasswordHash string `json:"PasswordHash"`
	ClusterId    uint   `json:"ClusterId"`
}
