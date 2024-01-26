package validators

import (
	"context"

	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/core/core"
	"github.com/productivityeng/orabbit/db"
)

type ClusterValidator interface {
	ValidateCreateRequest(request contracts.CreateClusterRequest, ctx context.Context) error
}

type clusterValidatorDefault struct {
	DependencyLocator  *core.DependencyLocator
}

func NewClusterValidatorDefault(DependencyLocator *core.DependencyLocator, ) *clusterValidatorDefault {
	return &clusterValidatorDefault{
		DependencyLocator: DependencyLocator,
	}
}

func (val *clusterValidatorDefault) ValidateCreateRequest(request contracts.CreateClusterRequest, ctx context.Context) error {

	var err error


	err = val.validateIfClusterWithThisHostnameExists(request, ctx)
	if err != nil {
		return err
	}


	return err
}

func (val *clusterValidatorDefault) validateIfClusterWithThisHostnameExists(request contracts.CreateClusterRequest, ctx context.Context) error {
	unique := db.Cluster.UniqueNameHostPort(
		db.Cluster.Host.Equals(request.Host),
		db.Cluster.Port.Equals(request.Port),
	)
	_,err := val.DependencyLocator.PrismaClient.Cluster.FindUnique(unique).Exec(ctx)
	if err != nil { 
		return err
	}
	return nil
}
