package broker

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/broker/repository"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/core/validators"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type BrokerController interface {
	ListBrokers(c *gin.Context)
	CreateBroker(c *gin.Context)
	UpdateBroker(c *gin.Context)
	DeleteBroker(c *gin.Context)
}

type brokerControllerDefaultImp struct {
	BrokerRepository repository.BrokerRepositoryInterface
	BrokerValidator  validators.BrokerValidator
}

func NewBrokerController(BrokerRepository repository.BrokerRepositoryInterface, BrokerValidator validators.BrokerValidator) *brokerControllerDefaultImp {
	return &brokerControllerDefaultImp{BrokerRepository: BrokerRepository, BrokerValidator: BrokerValidator}
}

// @BasePath /

// PingExample godoc
// @Summary Register a new RabbitMQ Broker
// @Schemes
// @Description Create a new <b>RabbitMQ</b> broker. The credential provider must be valid and the cluster operational
// @Tags Broker
// @Accept json
// @Produce json
// @Param request body contracts.CreateBrokerRequest true "Request"
// @Success 201 {string} Helloworld
// @Router /broker [post]
func (ctrl *brokerControllerDefaultImp) CreateBroker(c *gin.Context) {
	var request contracts.CreateBrokerRequest
	if err := c.BindJSON(&request); err != nil {
		log.WithError(err).Error("Error parsing request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.BrokerValidator.ValidateCreateRequest(request); err != nil {
		log.WithError(err).Error("Error validating request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := ctrl.BrokerRepository.CreateBroker(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (ctrl *brokerControllerDefaultImp) DeleteBroker(c *gin.Context) {

}

func (ctrl *brokerControllerDefaultImp) ListBrokers(c *gin.Context) {
	type PageParam struct {
		PageSize   int `json:"pageSize" binding:"required,gt=0"`
		PageNumber int `json:"pageNumber" binding:"required,gt=0"`
	}
	var param PageParam
	err := c.BindQuery(&param)
	if err != nil {
		log.WithError(err).Error("Error trying to parse query params")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, nil)

}

// PingExample godoc
// @Summary Retrieve a list of rabbitmq clusters registered
// @Schemes
// @Description Retrieve a paginated list of cluster that the user has access
// @Tags Broker
// @Accept json
// @Produce json
// @Param pageSize query int true "Number of items in one page"
// @Param pageNumber query int true "Page number starting from 1"
// @Router /broker [get]
func (ctrl *brokerControllerDefaultImp) UpdateBroker(c *gin.Context) {

}
