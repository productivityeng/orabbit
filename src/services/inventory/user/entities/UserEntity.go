package entities

import (
	"github.com/productivityeng/orabbit/broker/entities"
	"gorm.io/gorm"
)

type UserEntity struct {
	gorm.Model
	Id           int32                 `gorm:"primaryKey"`
	Username     string                `json:"Username"`
	PasswordHash string                `json:"PasswordHash"`
	BrokerId     int32                 `json:"BrokerId" gorm:""`
	Broker       entities.BrokerEntity `json:"Broker" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (UserEntity) TableName() string {
	return "entity"
}
