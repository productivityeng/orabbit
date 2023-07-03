package entities

import (
	"gorm.io/gorm"
	"time"
)

type BrokerEntity struct {
	Id          int32          `gorm:primaryKey`
	CreatedAt   time.Time      `json:createdAt`
	UpdatedAt   time.Time      `json:updatedAt`
	DeletedAt   gorm.DeletedAt `gorm:Index,json:deletedAt`
	Name        string         `json:name`
	Description string         `json:description`
	Host        string         `json:host`
	Port        int32          `json:"port"`
	User        string         `json:user`
	Password    string         `json:password`
}

func (BrokerEntity) TableName() string {
	return "broker"
}
