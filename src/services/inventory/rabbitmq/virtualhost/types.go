package virtualhost

import "github.com/productivityeng/orabbit/rabbitmq/common"


type CreateVirtualHostRequest struct {
	common.RabbitAccess
	Name        string
	Description string
}

type ListVirtualHostRequest struct {
	common.RabbitAccess
}
