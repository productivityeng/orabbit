package validators

import (
	"context"
	"errors"
	"github.com/productivityeng/orabbit/broker/repository"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq"
)

type BrokerValidator interface {
	ValidateCreateRequest(request contracts.CreateBrokerRequest, ctx context.Context) error
}

type brokerValidatorDefault struct {
	BrokerRepository   repository.BrokerRepositoryInterface
	OverviewManagement rabbitmq.OverviewManagement
}

func NewBrokerValidatorDefault(brokerRepository repository.BrokerRepositoryInterface, OverviewManagement rabbitmq.OverviewManagement) *brokerValidatorDefault {
	return &brokerValidatorDefault{
		BrokerRepository:   brokerRepository,
		OverviewManagement: OverviewManagement,
	}
}

func (val *brokerValidatorDefault) ValidateCreateRequest(request contracts.CreateBrokerRequest, ctx context.Context) error {

	var err error
	err = val.validateIfBrokerWithThisHostnameExists(request, ctx)
	err = val.OverviewManagement.CheckCredentials(rabbitmq.CheckCredentialsRequest{
		Host:     request.Host,
		Username: request.User,
		Password: request.Password,
		Port:     request.Port,
	})

	return err
}

func (val *brokerValidatorDefault) validateIfBrokerWithThisHostnameExists(request contracts.CreateBrokerRequest, ctx context.Context) error {
	exists := val.BrokerRepository.CheckIfHostIsAlreadyRegisted(request.Host, request.Port, ctx)
	if exists {
		return errors.New("[BROKER_ALREADY_EXISTS]")
	}
	return nil
}
