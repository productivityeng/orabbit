package entities

import (
	"github.com/productivityeng/orabbit/cluster/entities"
	"gorm.io/gorm"
)

type QueueEntity struct {
	gorm.Model
	ClusterID int32                  `json:"ClusterID"`
	Cluster   entities.ClusterEntity `json:"Cluster"`
	Name      string                 `json:"Name"`
	Type      string                 `json:"Type"`
	Durable   bool                   `json:"Durable"`
	Arguments map[string]string      `gorm:"serializer:json"`
}

func (QueueEntity) TableName() string {
	return "Queue"
}
