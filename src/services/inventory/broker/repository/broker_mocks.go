package repository

import (
	"github.com/productivityeng/orabbit/broker/entities"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/stretchr/testify/mock"
)

type BrokerRepositoryMockedObject struct {
	mock.Mock
}

func (brmo *BrokerRepositoryMockedObject) CreateBroker(entityToCreate *entities.BrokerEntity) (*entities.BrokerEntity, error) {
	args := brmo.Called(entityToCreate)
	return args.Get(0).(*entities.BrokerEntity), args.Error(1)
}

func (brmo *BrokerRepositoryMockedObject) ListBroker(pageSize int, pageNumber int) (*contracts.PaginatedResult[entities.BrokerEntity], error) {
	args := brmo.Called(pageSize, pageNumber)
	return args.Get(0).(*contracts.PaginatedResult[entities.BrokerEntity]), args.Error(1)
}
