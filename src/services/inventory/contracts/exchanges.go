package contracts

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/pkg/exchange/dto"
	"github.com/productivityeng/orabbit/pkg/rabbitmq/common"
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

type GetExchangeRequest struct {
	common.RabbitAccess
	Name string `json:"Name"`
	VirtualHostName string `json:"VirtualHostName"`
 }

type DeleteExchangeRequest struct {
	common.RabbitAccess
	Name string `json:"Name"`
}

type ExchangeManagement interface {
	GetAllExchangesFromCluster(request ListExchangeRequest,c *gin.Context) ([]dto.GetExchangeDto, error)
	CreateExchange(request CreateExchangeRequest) (error) 
	DeleteExchange(request DeleteExchangeRequest,c *gin.Context) (error)
	GetExchangeByName(request GetExchangeRequest,c *gin.Context) (*dto.GetExchangeDto, error)
}