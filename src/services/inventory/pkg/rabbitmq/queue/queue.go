package queue

import (
	"errors"
	"fmt"

	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/sirupsen/logrus"
)

type QueueManagementImpl struct {
}

func NewQueueManagement() contracts.QueueManagement {
	return &QueueManagementImpl{}
}

func (q QueueManagementImpl) GetAllQueuesFromCluster(request contracts.ListQueuesRequest) ([]rabbithole.QueueInfo, error) {
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

func (q QueueManagementImpl) GetQueueFromCluster(request contracts.GetQueueRequest) (*rabbithole.DetailedQueueInfo, error) {
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

func (q QueueManagementImpl) CreateQueue(request contracts.CreateQueueRequest) (*rabbithole.DetailedQueueInfo, error) {

	logrus.WithField("request", request).Info("Declaring queue ...")
	rmqc, err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.Username, request.Password)
	if err != nil {
		logrus.WithError(err).Error("Error trying to connect to cluster")
		return nil, errors.New("[CLUSTER_CONNECT_FAIL]")
	}
	_, err = rmqc.DeclareQueue(request.Vhost, request.Queue, rabbithole.QueueSettings{
		Type:       request.Type,
		Durable:    request.Durable,
		AutoDelete: false,
		Arguments:  request.Arguments,
	})
	if err != nil {
		logrus.WithError(err).WithField("request", request).Error("Fail to declare queue")
		return nil,err
	}

	return q.GetQueueFromCluster(contracts.GetQueueRequest{
		RabbitAccess: request.RabbitAccess,
		Queue:        request.Queue,
	})
}

func (q QueueManagementImpl) DeleteQueue(request contracts.DeleteQueueRequest) error {
	logrus.WithField("request", request).Info("Declaring queue ...")
	rmqc, err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.Username, request.Password)
	if err != nil {
		logrus.WithError(err).Error("Error trying to connect to cluster")
		return errors.New("[CLUSTER_CONNECT_FAIL]")
	}

	_, err = rmqc.DeleteQueue("/", request.Queue)
	return err
}

func (q QueueManagementImpl) GetQueueBindingsFromCluster(request contracts.GetQueueBindingsRequest) ([]rabbithole.BindingInfo, error) {
	rmqc, err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.Username, request.Password)
	if err != nil {
		logrus.WithError(err).Error("Error trying to connect to cluster")
		return nil, errors.New("[CLUSTER_CONNECT_FAIL]")
	}

	bindings, err := rmqc.ListQueueBindings(request.VirtualHostName, request.Name)
	if err != nil {
		logrus.WithError(err).Error("Error trying to list queue bindings from cluster")
		return nil, errors.New("[CLUSTER_LISTQUEUEBINDINGS_FAIL]")
	}
	return bindings, nil
}

func (q QueueManagementImpl) CreateQueueBinding(request contracts.CreateQueueBindingRequest) error {
	rmqc, err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.Username, request.Password)
	if err != nil { 
		logrus.WithError(err).Error("Error trying to connect to cluster")
		return errors.New("[CLUSTER_CONNECT_FAIL]")
	}
	bindingInfo := rabbithole.BindingInfo{
		Source: request.ExchangeName,
		Vhost:  request.VHost,
		Destination: request.QueueName,
		DestinationType: "queue",
		RoutingKey: request.RoutingKey,
		Arguments: request.Arguments,
	}
	result,err :=rmqc.DeclareBinding(request.VHost,bindingInfo)
	if err != nil { 
		logrus.WithError(err).Error("Error trying to create queue binding")
		return errors.New("[CLUSTER_CREATEQUEUEBINDING_FAIL]")
	}

	logrus.WithField("result", result).Info("Queue binding created")
	return nil
}