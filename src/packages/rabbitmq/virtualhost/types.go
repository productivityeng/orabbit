package virtualhost

import "github.com/productivityeng/orabbit/src/packages/rabbitmq/common"

type CreateVirtualHostRequest struct {
	common.RabbitAccess
	Name        string
	Description string
}

type ListVirtualHostRequest struct {
	common.RabbitAccess
}
