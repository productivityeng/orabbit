package broker

import (
	"github.com/gin-gonic/gin"
	broker_controller "github.com/productivityeng/orabbit/broker/controllers"
	"github.com/productivityeng/orabbit/broker/repository"
	"github.com/productivityeng/orabbit/core/validators"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq"
	"gorm.io/gorm"
)

var brokerController broker_controller.BrokerController
var brokerRepository repository.BrokerRepositoryInterface
var brokerValidator validators.BrokerValidator
var overviewManagement rabbitmq.OverviewManagement

func Routes(routes *gin.Engine, db *gorm.DB) {
	overviewManagement = rabbitmq.NewOverviewManagementImpl()
	brokerRepository = repository.NewBrokerMysqlImpl(db)
	brokerValidator = validators.NewBrokerValidatorDefault(brokerRepository, overviewManagement)
	brokerController = broker_controller.NewBrokerController(brokerRepository, brokerValidator)

	brokerResourcePath := "/:brokerId"
	brokerRoutes := routes.Group("/broker")

	brokerRoutes.GET("/", brokerController.ListBrokers)
	brokerRoutes.GET("/exists", brokerController.FindBroker)
	brokerRoutes.GET(brokerResourcePath, brokerController.GetBroker)
	brokerRoutes.PUT(brokerResourcePath, brokerController.UpdateBroker)
	brokerRoutes.POST("/", brokerController.CreateBroker)
	brokerRoutes.DELETE(brokerResourcePath, brokerController.DeleteBroker)
}
