package repository

import (
	"context"
	"github.com/productivityeng/orabbit/broker/entities"
	"github.com/productivityeng/orabbit/contracts"
)

type BrokerRepositoryInterface interface {
	CreateBroker(broker *entities.BrokerEntity) (*entities.BrokerEntity, error)
	ListBroker(pageSize int, pageNumber int) (*contracts.PaginatedResult[entities.BrokerEntity], error)
	DeleteBroker(brokerId int32, ctx context.Context) error
}
