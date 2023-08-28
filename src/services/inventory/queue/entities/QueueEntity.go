package entities

import (
	"github.com/productivityeng/orabbit/cluster/entities"
	"gorm.io/gorm"
)

type QueueEntity struct {
	gorm.Model
	ClusterId uint                   `json:"ClusterId" gorm:"index:idx_unique_queuename_by_host,unique"`
	Cluster   entities.ClusterEntity `json:"Cluster" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE,foreignKey:ClusterId"`
	Name      string                 `json:"Name" gorm:"index:idx_unique_queuename_by_host,unique"`
	Type      string                 `json:"Type"`
	Durable   bool                   `json:"Durable"`
	Arguments map[string]interface{} `gorm:"serializer:json"`
}

func (QueueEntity) TableName() string {
	return "Queue"
}
