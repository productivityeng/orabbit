package repository

import (
	"context"
	"github.com/productivityeng/orabbit/cluster/entities"
	"github.com/productivityeng/orabbit/contracts"
)

type ClusterRepositoryInterface interface {
	///CreateCluster store a new broker in storage with provided parameter
	CreateCluster(broker *entities.ClusterEntity) (*entities.ClusterEntity, error)
	//ListCluster retrieve a lista of broker with paginated options
	ListCluster(pageSize int, pageNumber int) (*contracts.PaginatedResult[entities.ClusterEntity], error)
	//DeleteCluster soft delete a broker with a provided brokerId
	DeleteCluster(clusterId uint, ctx context.Context) error
	//GetCluster retrieve a broker with a provided brokerId
	GetCluster(clusterId uint, ctx context.Context) (*entities.ClusterEntity, error)
	CheckIfHostIsAlreadyRegisted(host string, port int32, ctx context.Context) bool
}
