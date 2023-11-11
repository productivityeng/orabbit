package validators

import (
	"context"
	"errors"
	"github.com/productivityeng/orabbit/cluster/repository"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq"
)

type ClusterValidator interface {
	ValidateCreateRequest(request contracts.CreateClusterRequest, ctx context.Context) error
}

type clusterValidatorDefault struct {
	ClusterRepository  repository.ClusterRepositoryInterface
	OverviewManagement rabbitmq.OverviewManagement
}

func NewClusterValidatorDefault(brokerRepository repository.ClusterRepositoryInterface, OverviewManagement rabbitmq.OverviewManagement) *clusterValidatorDefault {
	return &clusterValidatorDefault{
		ClusterRepository:  brokerRepository,
		OverviewManagement: OverviewManagement,
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
	exists := val.ClusterRepository.CheckIfHostIsAlreadyRegisted(request.Host, request.Port, ctx)
	if exists {
		return errors.New("[BROKER_ALREADY_EXISTS]")
	}
	return nil
}
