package broker

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/broker/entities"
	"github.com/productivityeng/orabbit/broker/repository"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/core/validators"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type PageParam struct {
	PageSize   int `json:"PageSize" binding:"required,gt=0"`
	PageNumber int `json:"PageNumber" binding:"required,gt=0"`
}
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

	entityToCreate := &entities.BrokerEntity{Name: request.Name, Host: request.Host, User: request.User, Password: request.Password, Port: request.Port,
		Description: request.Description}

	resp, err := ctrl.BrokerRepository.CreateBroker(entityToCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// PingExample godoc
// @Summary Soft delete a broker
// @Schemes
// @Description Soft delete a broker will not completly erase from database, but will not show up anymore in the
// system. All queues,bindings,shovels and related artifacts will be soft delete to
// @Tags Broker
// @Accept json
// @Produce json
// @Success 204 {object} bool
// @Param brokerId path int true "Id of a broker to be soft deleted"
// @Router /broker/{brokerId} [delete]
func (ctrl *brokerControllerDefaultImp) DeleteBroker(c *gin.Context) {
	brokerIdParam := c.Param("brokerId")
	brokerId, err := strconv.ParseInt(brokerIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("brokerId", brokerIdParam).Error("Fail to parse brokerId Param")
		c.JSON(http.StatusBadRequest, "Error parsing brokerId from url route")
		return
	}

	log.WithField("brokerId", brokerId).Info("Broker will be deleted")
	err = ctrl.BrokerRepository.DeleteBroker(int32(brokerId), c)
	if err != nil {
		errorMsg := "Fail to delete brokerId"
		log.WithError(err).WithField("brokerId", brokerId).Error(errorMsg)
		c.JSON(http.StatusInternalServerError, errorMsg)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("broker with id %d successufly deleted", brokerId)})
}

// PingExample godoc
// @Summary Retrieve a list of rabbitmq clusters registered
// @Schemes
// @Description Retrieve a paginated list of cluster that the user has access
// @Tags Broker
// @Accept json
// @Produce json
// @Success 200 {object} contracts.PaginatedResult[entities.BrokerEntity]
// @Param params query PageParam true "Number of items in one page"
// @Router /broker [get]
func (ctrl *brokerControllerDefaultImp) ListBrokers(c *gin.Context) {

	var param PageParam
	err := c.BindQuery(&param)
	if err != nil {
		log.WithError(err).Error("Error trying to parse query params")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paginatedBrokers, err := ctrl.BrokerRepository.ListBroker(param.PageSize, param.PageNumber)
	if err != nil {
		log.WithError(err).Error("Error retrieving brokers from repository")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, paginatedBrokers)

}

func (ctrl *brokerControllerDefaultImp) UpdateBroker(c *gin.Context) {

}
