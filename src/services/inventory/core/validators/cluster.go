package validators

import (
	"context"

	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/core/core"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq"
)

type ClusterValidator interface {
	ValidateCreateRequest(request contracts.CreateClusterRequest, ctx context.Context) error
}

type clusterValidatorDefault struct {
	OverviewManagement rabbitmq.OverviewManagement
	DependencyLocator  *core.DependencyLocator
}

func NewClusterValidatorDefault(DependencyLocator *core.DependencyLocator, OverviewManagement rabbitmq.OverviewManagement) *clusterValidatorDefault {
	return &clusterValidatorDefault{
		OverviewManagement: OverviewManagement,
		DependencyLocator: DependencyLocator,
	}
}

func (val *clusterValidatorDefault) ValidateCreateRequest(request contracts.CreateClusterRequest, ctx context.Context) error {

	var err error
	err = val.validateIfClusterWithThisHostnameExists(request, ctx)
	if err != nil {
		return err
	}
	err = val.OverviewManagement.CheckCredentials(rabbitmq.CheckCredentialsRequest{
		Host:     request.Host,
		Username: request.User,
		Password: request.Password,
		Port:     request.Port,
	})

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
