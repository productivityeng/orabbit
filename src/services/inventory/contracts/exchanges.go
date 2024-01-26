package contracts

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/exchange/dto"
	"github.com/productivityeng/orabbit/rabbitmq/common"
)


type ListExchangeRequest struct {
	common.RabbitAccess
}

type CreateExchangeRequest struct { 
	common.RabbitAccess
	Name string `json:"Name"`
	Type string `json:"Type"`
	ClusterId int `json:"ClusterId"`
	Internal bool `json:"Internal"`
	Durable bool `json:"Durable"`
	Arguments map[string]interface{} `json:"Arguments"`
}

type ExchangeManagement interface {
	GetAllExchangesFromCluster(request ListExchangeRequest,c *gin.Context) ([]dto.GetExchangeDto, error)
	CreateExchange(request CreateExchangeRequest) (error) 
}