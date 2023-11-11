package virtualhost

import (
	"fmt"
	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	"github.com/sirupsen/logrus"
)

type VirtualHostManagement interface {
	CreateVirtualHost(request CreateVirtualHostRequest) error
	ListVirtualHosts(request ListVirtualHostRequest) ([]rabbithole.VhostInfo, error)
}

type VirtualHostManagementImpl struct {
}

func NewirtualHostManagement() VirtualHostManagement {
	return &VirtualHostManagementImpl{}
}

func (management VirtualHostManagementImpl) ListVirtualHosts(request ListVirtualHostRequest) ([]rabbithole.VhostInfo, error) {
	rmqc, err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.RabbitAccess.Username, request.RabbitAccess.Password)
	if err != nil {
		logrus.WithError(err).Error(fmt.Sprintf("Ao criar client para o cluster %s com ususario %s", request.RabbitAccess.Host, request.RabbitAccess.Username))
		return nil, err
	}

	vhosts, err := rmqc.ListVhosts()

	if err != nil {
		logrus.WithError(err).Error(fmt.Sprintf("Erro ao listar os Vhosts com usuario %s do cluster %s", request.Username, request.Host))
		return nil, err
	}

	return vhosts, nil
}

func (management VirtualHostManagementImpl) CreateVirtualHost(request CreateVirtualHostRequest) error {
	rmqc, err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.RabbitAccess.Username, request.RabbitAccess.Password)
	if err != nil {
		logrus.WithError(err).Error(fmt.Sprintf("Ao criar client para o cluster %s com ususario %s", request.RabbitAccess.Host, request.RabbitAccess.Username))
		return err
	}

	resp, err := rmqc.PutVhost(request.Name, rabbithole.VhostSettings{
		Description:      request.Description,
		DefaultQueueType: "classic",
	})

	if err != nil {
		logrus.WithError(err).Error(fmt.Sprintf("Erro ao remover usuario %s do cluster %s", request.Username, request.Host))
		return err
	}

	logrus.WithField("resp", resp).Info("Virtual Host criado com sucesso")
	return nil
}
