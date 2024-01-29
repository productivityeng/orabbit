package contracts

import "github.com/productivityeng/orabbit/pkg/rabbitmq/common"







type GetUserByNameRequest struct { 
	common.RabbitAccess
	Username string
}

