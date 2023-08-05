package validators

import (
	"context"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/stretchr/testify/mock"
)

type BrokerValidatorMockedObject struct {
	mock.Mock
}

func (blmo *BrokerValidatorMockedObject) ValidateCreateRequest(request contracts.CreateBrokerRequest, ctx context.Context) error {
	args := blmo.Mock.Called(request, ctx)
	return args.Error(0)
}
