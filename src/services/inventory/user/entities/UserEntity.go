package entities

import (
	"github.com/productivityeng/orabbit/cluster/entities"
	"gorm.io/gorm"
)

type UserEntity struct {
	gorm.Model
	Id           int32                  `gorm:"primaryKey"`
	Username     string                 `json:"Username" gorm:"index:idx_unique_username_by_host,unique"`
	PasswordHash string                 `json:"PasswordHash"`
	ClusterId    int32                  `json:"ClusterId" gorm:"index:idx_unique_username_by_host,unique"`
	Cluster      entities.ClusterEntity `json:"Cluster" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (UserEntity) TableName() string {
	return "User"
}
