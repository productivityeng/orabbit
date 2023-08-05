//go:build !mock
// +build !mock

package repository

import (
	"context"
	"github.com/productivityeng/orabbit/broker/entities"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/stretchr/testify/mock"
)

type BrokerRepositoryMockedObject struct {
	mock.Mock
}

func (brmo *BrokerRepositoryMockedObject) CreateBroker(entityToCreate *entities.BrokerEntity) (*entities.BrokerEntity, error) {
	args := brmo.Called(entityToCreate)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.BrokerEntity), args.Error(1)
}

func (brmo *BrokerRepositoryMockedObject) ListBroker(pageSize int, pageNumber int) (*contracts.PaginatedResult[entities.BrokerEntity], error) {
	args := brmo.Called(pageSize, pageNumber)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*contracts.PaginatedResult[entities.BrokerEntity]), args.Error(1)
}

func (brmo *BrokerRepositoryMockedObject) DeleteBroker(brokerId int32, ctx context.Context) error {

	args := brmo.Called(brokerId, ctx)
	return args.Error(0)
}

func (brmo *BrokerRepositoryMockedObject) GetBroker(brokerId int32, ctx context.Context) (*entities.BrokerEntity, error) {
	args := brmo.Called(brokerId, ctx)
	return args.Get(0).(*entities.BrokerEntity), args.Error(1)
}

func (brmo *BrokerRepositoryMockedObject) CheckIfHostIsAlreadyRegisted(host string, port int32, ctx context.Context) bool {
	args := brmo.Called(host, port, ctx)
	return args.Get(0).(bool)
}
