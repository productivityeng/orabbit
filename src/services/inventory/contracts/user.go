package contracts

type GetUserResponse struct {
	Id           int32  `json:"Id"`
	Username     string `json:"Username"`
	PasswordHash string `json:"PasswordHash"`
	IsRegistered bool   `json:"IsRegistered"`
}
