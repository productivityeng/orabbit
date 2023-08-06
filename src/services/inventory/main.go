package main

import (
	"github.com/gin-gonic/gin"
	broker_controller "github.com/productivityeng/orabbit/broker/controllers"
	brokerEntities "github.com/productivityeng/orabbit/broker/entities"
	"github.com/productivityeng/orabbit/broker/repository"
	"github.com/productivityeng/orabbit/core/validators"
	database_mysql "github.com/productivityeng/orabbit/database"
	"github.com/productivityeng/orabbit/docs"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq"
	userEntities "github.com/productivityeng/orabbit/user/entities"
	log "github.com/sirupsen/logrus"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var brokerController broker_controller.BrokerController
var brokerRepository repository.BrokerRepositoryInterface
var brokerValidator validators.BrokerValidator
var overviewManagement rabbitmq.OverviewManagement

func main() {
	database_mysql.Db.AutoMigrate(&brokerEntities.BrokerEntity{}, &userEntities.UserEntity{})

	gin.ForceConsoleColor()

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	overviewManagement = rabbitmq.NewOverviewManagementImpl()
	brokerRepository = repository.NewBrokerMysqlImpl(database_mysql.Db)
	brokerValidator = validators.NewBrokerValidatorDefault(brokerRepository, overviewManagement)
	brokerController = broker_controller.NewBrokerController(brokerRepository, brokerValidator)
	brokerResourcePath := "/broker/:brokerId"
	r.GET("/broker", brokerController.ListBrokers)
	r.GET("/broker/exists", brokerController.FindBroker)
	r.GET(brokerResourcePath, brokerController.GetBroker)
	r.PUT(brokerResourcePath, brokerController.UpdateBroker)
	r.POST("/broker", brokerController.CreateBroker)
	r.DELETE(brokerResourcePath, brokerController.DeleteBroker)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	log.Fatal(r.Run(":8082"))
}
