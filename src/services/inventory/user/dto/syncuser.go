package dto

type ImportUserRequest struct {
	BrokerId int32  `json:"BrokerId" binding:"required"`
	Username string `json:"Username" binding:"required"`
}

type GetUserResponse struct {
	Id           int32  `json:"Id"`
	Username     string `json:"Username"`
	PasswordHash string `json:"PasswordHash"`
	BrokerId     int32  `json:"BrokerId"`
}
