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

type GetQueueRequest struct {
	common.RabbitAccess
	Queue string
}

func NewQueueManagement() QueueManagement {
	return &QueueManagementImpl{}
}

type QueueManagement interface {
	GetAllQueuesFromCluster(request ListQueuesRequest) ([]rabbithole.QueueInfo, error)
	GetQueueFromCluster(request GetQueueRequest) (*rabbithole.DetailedQueueInfo, error)
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

func (q QueueManagementImpl) GetQueueFromCluster(request GetQueueRequest) (*rabbithole.DetailedQueueInfo, error) {
	rmqc, err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.Username, request.Password)
	if err != nil {
		logrus.WithError(err).Error("Error trying to connect to cluster")
		return nil, errors.New("[CLUSTER_CONNECT_FAIL]")
	}

	queue, err := rmqc.GetQueue("/", request.Queue)
	if err != nil {
		if (err.(rabbithole.ErrorResponse)).StatusCode == 404 {
			logrus.WithError(err).Warn("Queue not found")
			return nil, nil
		}
		logrus.WithError(err).Error("Error trying to get queue from cluster")
		return nil, errors.New("[CLUSTER_GETQUEUE_FAIL]")
	}
	return queue, nil
}
