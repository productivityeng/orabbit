package validators

import (
	"github.com/productivityeng/orabbit/contracts"
	"github.com/stretchr/testify/mock"
)

type BrokerValidatorMockedObject struct {
	mock.Mock
}

func (blmo *BrokerValidatorMockedObject) ValidateCreateRequest(request contracts.CreateBrokerRequest) error {
	args := blmo.Mock.Called()
	return args.Error(0)
}
