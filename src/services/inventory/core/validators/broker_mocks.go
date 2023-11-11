package validators

import (
	"context"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/stretchr/testify/mock"
)

type ClusterValidatorMockedObject struct {
	mock.Mock
}

func (blmo *ClusterValidatorMockedObject) ValidateCreateRequest(request contracts.CreateClusterRequest, ctx context.Context) error {
	args := blmo.Mock.Called(request, ctx)
	return args.Error(0)
}
