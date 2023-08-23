package queue

import (
	"errors"
	"fmt"
	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/common"
	"github.com/sirupsen/logrus"
)

type ListQueuesRequest struct {
	common.RabbitAccess
}

func NewQueueManagement() QueueManagement {
	return &QueueManagementImpl{}
}

type QueueManagement interface {
	GetAllQueuesFromCluster(request ListQueuesRequest) ([]rabbithole.QueueInfo, error)
}

type QueueManagementImpl struct {
}

func (q QueueManagementImpl) GetAllQueuesFromCluster(request ListQueuesRequest) ([]rabbithole.QueueInfo, error) {
	rmqc, err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.Username, request.Password)
	if err != nil {
		logrus.WithError(err).Error("Error trying to connect to cluster")
		return nil, errors.New("[CLUSTER_CONNECT_FAIL]")
	}

	queues, err := rmqc.ListQueues()

	if err != nil {
		logrus.WithError(err).Error("Error trying to list queues from cluster")
		return nil, errors.New("[CLUSTER_LISTQUEUE_FAIL]")
	}
	return queues, nil

}
