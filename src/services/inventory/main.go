package main

import (
	"github.com/productivityeng/orabbit/broker/repository"
	"github.com/productivityeng/orabbit/core/validators"
	database_mysql "github.com/productivityeng/orabbit/database"
	"github.com/productivityeng/orabbit/docs"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq"
	"net/http"

	"github.com/gin-gonic/gin"
	broker_controller "github.com/productivityeng/orabbit/broker/controllers"
	"github.com/productivityeng/orabbit/broker/entities"
	log "github.com/sirupsen/logrus"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var brokerController broker_controller.BrokerController
var brokerRepository repository.BrokerRepositoryInterface
var brokerValidator validators.BrokerValidator
var overviewManagement rabbitmq.OverviewManagement

func main() {
	database_mysql.Db.AutoMigrate(&entities.BrokerEntity{})

	gin.ForceConsoleColor()

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	overviewManagement = rabbitmq.NewOverviewManagementImpl()
	brokerRepository = repository.NewBrokerMysqlImpl(database_mysql.Db)
	brokerValidator = validators.NewBrokerValidatorDefault(brokerRepository, overviewManagement)
	brokerController = broker_controller.NewBrokerController(brokerRepository, brokerValidator)

	r.GET("/broker", brokerController.ListBrokers)
	r.GET("/broker/exists", brokerController.FindBroker)
	r.GET("/broker/:brokerId", brokerController.GetBroker)
	r.PUT("/broker/:brokerId", brokerController.UpdateBroker)
	r.POST("/broker", brokerController.CreateBroker)
	r.DELETE("/broker/:brokerId", brokerController.DeleteBroker)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	log.Fatal(r.Run(":8082"))
}
