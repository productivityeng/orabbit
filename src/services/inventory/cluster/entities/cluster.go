package entities

import (
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/common"
	"gorm.io/gorm"
)

type ClusterEntity struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Host        string `json:"host" gorm:"index:idx_unique_host,unique"`
	Port        int32  `json:"port" gorm:"index:idx_unique_host,unique"`
	User        string `json:"user"`
	Password    string `json:"password"`
}

func (ClusterEntity) TableName() string {
	return "Cluster"
}

func (cluster ClusterEntity) GetRabbitMqAccess() common.RabbitAccess {
	return common.RabbitAccess{
		Host:     cluster.Host,
		Port:     cluster.Port,
		Username: cluster.User,
		Password: cluster.Password,
	}
}
