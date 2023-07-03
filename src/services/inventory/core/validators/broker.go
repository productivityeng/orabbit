package validators

import "github.com/productivityeng/orabbit/contracts"

type BrokerValidator interface {
	ValidateCreateRequest(request contracts.CreateBrokerRequest) error
}

type brokerValidatorDefault struct {
}

func NewBrokerValidatorDefault() *brokerValidatorDefault {
	return &brokerValidatorDefault{}
}

func (val *brokerValidatorDefault) ValidateCreateRequest(request contracts.CreateBrokerRequest) error {
	return nil
}
