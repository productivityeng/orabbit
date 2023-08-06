package dto

type ImportUserRequest struct {
	BrokerId int32  `json:"BrokerId" binding:"required"`
	Username string `json:"Username" binding:"required"`
}
