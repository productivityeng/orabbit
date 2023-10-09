package entities

import "gorm.io/gorm"

type LockerEntity struct {
	gorm.Model
	Reason string `json:"Reason"`
}

func (LockerEntity) TableName() string {
	return "Locker"
}
