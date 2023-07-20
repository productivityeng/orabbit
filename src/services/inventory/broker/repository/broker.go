package repository

import (
	"context"
	"github.com/productivityeng/orabbit/broker/entities"
	"github.com/productivityeng/orabbit/contracts"
)

type BrokerRepositoryInterface interface {
	///CreateBroker store a new broker in storage with provided parameter
	CreateBroker(broker *entities.BrokerEntity) (*entities.BrokerEntity, error)
	//ListBroker retrieve a lista of broker with paginated options
	ListBroker(pageSize int, pageNumber int) (*contracts.PaginatedResult[entities.BrokerEntity], error)
	//DeleteBroker soft delete a broker with a provided brokerId
	DeleteBroker(brokerId int32, ctx context.Context) error
	//GetBroker retrieve a broker with a provided brokerId
	GetBroker(brokerId int32, ctx context.Context) (*entities.BrokerEntity, error)
	CheckIfHostIsAlreadyRegisted(host string, port int32, ctx context.Context) bool
}
