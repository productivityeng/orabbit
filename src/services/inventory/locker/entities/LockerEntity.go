package entities

import (
	"time"
)

type LockerEntity struct {
	ID         uint       `json:"ID" gorm:"primaryKey"`
	Reason     string     `json:"Reason"`
	EnabledAt  time.Time  `json:"EnabledAt"`
	DisabledAt *time.Time `json:"DisabledAt"`
}

func (LockerEntity) TableName() string {
	return "Locker"
}
