package rabbitmq

import (
	"errors"
	"fmt"
	"github.com/michaelklishin/rabbit-hole/v2"
	"github.com/sirupsen/logrus"
)

type CheckCredentialsRequest struct {
	Host     string
	Username string
	Password string
	Port     int32
}
type OverviewManagement interface {
	CheckCredentials(request CheckCredentialsRequest) error
}

type OverviewManagementImpl struct {
}

func NewOverviewManagementImpl() *OverviewManagementImpl {
	return &OverviewManagementImpl{}
}

func (receiver OverviewManagementImpl) CheckCredentials(request CheckCredentialsRequest) error {
	rmqc, err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.Username, request.Password)
	if err != nil {
		logrus.WithError(err).Error("Error trying to connect to cluster")
		return errors.New("[CLUSTER_CONNECT_FAIL]")
	}

	_, err = rmqc.Overview()
	if err != nil {
		logrus.WithError(err).Error("Error trying to retrieve overview of cluster")
		return errors.New("[CLUSTER_RETRIEVEOVERVIEW_FAIL]")
	}
	return nil
}
