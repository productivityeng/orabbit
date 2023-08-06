package entities

import "gorm.io/gorm"

type BrokerEntity struct {
	gorm.Model
	Id          int32  `gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Host        string `json:"host" gorm:"index:idx_unique_host,unique"`
	Port        int32  `json:"port" gorm:"index:idx_unique_host,unique"`
	User        string `json:"user"`
	Password    string `json:"password"`
}

func (BrokerEntity) TableName() string {
	return "broker"
}
