package models

import (
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/rabbitmq/common"
)


func GetRabbitMqAccess(cluster *db.ClusterModel) common.RabbitAccess {
	return common.RabbitAccess{
		Host:     cluster.Host,
		Port:     cluster.Port,
		Username: cluster.User,
		Password: cluster.Password,
	}
}
