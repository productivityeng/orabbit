package dto

import "time"

type LockUserDto struct {
	UserId     int64      `json:"userId"`
	Reason     string     `json:"reason"`
	Enabledt   time.Time  `json:"enabledt"`
	DisabledAt *time.Time `json:"disabledAt"`
}
