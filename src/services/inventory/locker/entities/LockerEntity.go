package entities

import (
	"time"

	"gorm.io/gorm"
)

type LockerEntity struct {
	gorm.Model
	Reason     string     `json:"Reason"`
	EnabledAt  time.Time  `json:"EnabledAt"`
	DisabledAt *time.Time `json:"DisabledAt"`
}

func (LockerEntity) TableName() string {
	return "Locker"
}
