package contracts

import "github.com/productivityeng/orabbit/rabbitmq/common"






type GetUserByNameRequest struct { 
	common.RabbitAccess
	Username string
}

