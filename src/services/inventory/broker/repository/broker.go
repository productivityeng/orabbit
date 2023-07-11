package repository

import (
	"github.com/productivityeng/orabbit/broker/entities"
	"github.com/productivityeng/orabbit/contracts"
)

type BrokerRepositoryInterface interface {
	CreateBroker(request contracts.CreateBrokerRequest) (*entities.BrokerEntity, error)
	ListBroker(pageSize int, pageNumber int) ([]entities.BrokerEntity, error)
}
