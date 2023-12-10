package dto

type LockUserDto struct {
	UserId int64  `json:"userId"`
	Reason string `json:"reason"`
}
