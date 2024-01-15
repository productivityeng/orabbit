package dto

type ImportOrCreateUserRequest struct {
	ClusterId int   `json:"ClusterId" binding:"required"`
	Username  string `json:"Username" binding:"required"`
	Password  string `json:"Password" `
	Create    bool   `json:"Create" `
}

type GetUserResponse struct {
	Id           int   `json:"Id"`
	Username     string `json:"Username"`
	PasswordHash string `json:"PasswordHash"`
	ClusterId    int   `json:"ClusterId"`
	LockedReason string `json:"LockedReason"`
}
