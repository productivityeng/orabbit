//go:build !mock
// +build !mock

package repository

import (
	"context"
	"github.com/productivityeng/orabbit/cluster/entities"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/stretchr/testify/mock"
)

type ClusterRepositoryMockedObject struct {
	mock.Mock
}

func (brmo *ClusterRepositoryMockedObject) CreateCluster(entityToCreate *entities.ClusterEntity) (*entities.ClusterEntity, error) {
	args := brmo.Called(entityToCreate)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.ClusterEntity), args.Error(1)
}

func (brmo *ClusterRepositoryMockedObject) ListCluster(pageSize int, pageNumber int) (*contracts.PaginatedResult[entities.ClusterEntity], error) {
	args := brmo.Called(pageSize, pageNumber)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*contracts.PaginatedResult[entities.ClusterEntity]), args.Error(1)
}

func (brmo *ClusterRepositoryMockedObject) DeleteCluster(brokerId uint, ctx context.Context) error {

	args := brmo.Called(brokerId, ctx)
	return args.Error(0)
}

func (brmo *ClusterRepositoryMockedObject) GetCluster(brokerId uint, ctx context.Context) (*entities.ClusterEntity, error) {
	args := brmo.Called(brokerId, ctx)
	return args.Get(0).(*entities.ClusterEntity), args.Error(1)
}

func (brmo *ClusterRepositoryMockedObject) CheckIfHostIsAlreadyRegisted(host string, port int32, ctx context.Context) bool {
	args := brmo.Called(host, port, ctx)
	return args.Get(0).(bool)
}
