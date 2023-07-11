package repository

import (
	"github.com/productivityeng/orabbit/broker/entities"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/stretchr/testify/mock"
)

type BrokerRepositoryMockedObject struct {
	mock.Mock
}

func (brmo *BrokerRepositoryMockedObject) CreateBroker(request contracts.CreateBrokerRequest) (*entities.BrokerEntity, error) {
	args := brmo.Called(request)
	return args.Get(0).(*entities.BrokerEntity), args.Error(1)
}

func (brmo *BrokerRepositoryMockedObject) ListBroker(pageSize int, pageNumber int) ([]entities.BrokerEntity, error) {
	args := brmo.Called(pageSize, pageNumber)
	return args.Get(0).([]entities.BrokerEntity), args.Error(1)
}
