package contracts

import (
	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	"github.com/productivityeng/orabbit/pkg/rabbitmq/common"
)


type ListQueuesRequest struct {
	common.RabbitAccess
}

type GetQueueRequest struct {
	common.RabbitAccess
	Queue string
	VirtualHostName string
}

type DeleteQueueRequest struct {
	common.RabbitAccess
	Queue string
}

type CreateQueueRequest struct {
	common.RabbitAccess
	Queue     string
	Vhost     string
	Type      string
	Durable   bool
	Arguments map[string]interface{}
}

type GetQueueBindingsRequest struct { 
	common.RabbitAccess
	Name string
	VirtualHostName string
}

type CreateQueueBindingRequest struct { 
	common.RabbitAccess
	QueueName string
	ExchangeName string
	RoutingKey string
	Arguments map[string]interface{}
	VHost string
}



type QueueManagement interface {
	GetAllQueuesFromCluster(request ListQueuesRequest) ([]rabbithole.QueueInfo, error)
	GetQueueFromCluster(request GetQueueRequest) (*rabbithole.DetailedQueueInfo, error)
	CreateQueue(request CreateQueueRequest) (*rabbithole.DetailedQueueInfo, error)
	DeleteQueue(request DeleteQueueRequest) error
	GetQueueBindingsFromCluster(request GetQueueBindingsRequest) ([]rabbithole.BindingInfo, error)
	CreateQueueBinding(request CreateQueueBindingRequest) error
}
