package entities

import (
	"github.com/productivityeng/orabbit/cluster/entities"
	"gorm.io/gorm"
)

type UserEntity struct {
	gorm.Model
	Username     string                 `json:"Username" gorm:"index:idx_unique_username_by_host,unique"`
	PasswordHash string                 `json:"PasswordHash"`
	ClusterId    uint                   `json:"ClusterId" gorm:"index:idx_unique_username_by_host,unique"`
	Cluster      entities.ClusterEntity `json:"Cluster" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE,foreignKey:ClusterId"`
}

func (UserEntity) TableName() string {
	return "User"
}
