package virtualhost

import "github.com/productivityeng/orabbit/pkg/rabbitmq/common"



type CreateVirtualHostRequest struct {
	common.RabbitAccess
	Name        string
	Description string
	DefaultQueueType string
	Tags []string
}

type ListVirtualHostRequest struct {
	common.RabbitAccess
}

type GetVirtualHostRequest struct { 
	common.RabbitAccess
	Name string
}

type DeleteVirtualHostRequest struct { 
	common.RabbitAccess
	Name string
}