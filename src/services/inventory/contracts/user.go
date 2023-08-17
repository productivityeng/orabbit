package contracts

type GetUserResponse struct {
	Id           int32  `json:"Id"`
	Name         string `json:"Name"`
	PasswordHash string `json:"PasswordHash"`
	IsRegistered bool   `json:"IsRegistered"`
}
