package exchange

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/pkg/exchange/dto"
	"github.com/sirupsen/logrus"
)





type ExchangeManagementImpl struct { 

}

func NewExchangeManagement() contracts.ExchangeManagement { 
	return ExchangeManagementImpl{}
}

func (management ExchangeManagementImpl) GetAllExchangesFromCluster(request contracts.ListExchangeRequest,c *gin.Context) ([]dto.GetExchangeDto, error) {
	
	rmqc, err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.Username, request.Password)
	if err != nil {
		logrus.WithContext(c).WithError(err).Error("Error trying to connect to cluster")
		return nil, errors.New("[CLUSTER_CONNECT_FAIL]")
	}

	exchanges, err :=rmqc.ListExchanges()
	if err != nil { 
		logrus.WithContext(c).WithError(err).Error("Error trying to list exchanges")
		return nil, errors.New("[CLUSTER_EXCHANGE_LIST_FAIL]")
	}

	exchangesResponse := make([]dto.GetExchangeDto,0)
	for _, exchange := range exchanges {
		exchangesResponse = append(exchangesResponse, dto.GetExchangeDto{ 
			Name: exchange.Name,
			VHost: exchange.Vhost,
			Type: exchange.Type,
			Durable: exchange.Durable,
			Internal: exchange.Internal,
			Arguments: exchange.Arguments,
			ClusterId: 0,
			IsInCluster: true,
		})
	 }

	return exchangesResponse,nil
}

func (management ExchangeManagementImpl) CreateExchange(request contracts.CreateExchangeRequest) (error) {
	rmqc,err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.Username, request.Password)
	if err != nil { 
		return err
	}
	settings := rabbithole.ExchangeSettings{
		Type: request.Type,
		Durable: request.Durable,
		Arguments: request.Arguments,
	}
	response,err := rmqc.DeclareExchange("/",request.Name,settings)
	if err != nil {
		logrus.WithError(err).Error("Error trying to create exchange")
		return err
	 }else {
		logrus.WithField("response", response).Info("Exchange created")
	 }
	return nil
}

func (management ExchangeManagementImpl) DeleteExchange(request contracts.DeleteExchangeRequest,c *gin.Context) (error) {
	rmqc,err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.Username, request.Password)
	if err != nil { 
		return err
	}
	response,err := rmqc.DeleteExchange("/",request.Name)
	if err != nil { 
		logrus.WithContext(c).WithError(err).Error("Error trying to delete exchange")
		return err
	 }else {
		logrus.WithContext(c).WithField("response", response).Info("Exchange deleted")
	 }
	 return nil
}

func (management ExchangeManagementImpl) GetExchangeByName(request contracts.GetExchangeRequest,c *gin.Context) (*dto.GetExchangeDto, error) {
	rmqc,err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.Username, request.Password)
	if err != nil { 
		return nil,err
	}
	exchange,err := rmqc.GetExchange(request.VirtualHostName,request.Name)
	if err != nil { 
		logrus.WithContext(c).WithField("exchange",exchange).WithError(err).Error("Error trying to get exchange")
		return nil,err
	 }else {
		logrus.WithContext(c).WithField("response", exchange).Info("Exchange retrieved")
		return &dto.GetExchangeDto{ 
			Name: exchange.Name,
			VHost: exchange.Vhost,
			Type: exchange.Type,
			Durable: exchange.Durable,
			Internal: exchange.Internal,
			Arguments: exchange.Arguments,
			IsInCluster: true,
		},nil
	}
}