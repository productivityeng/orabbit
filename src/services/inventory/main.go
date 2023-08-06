package main

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/broker"
	brokerEntities "github.com/productivityeng/orabbit/broker/entities"
	database_mysql "github.com/productivityeng/orabbit/database"
	"github.com/productivityeng/orabbit/docs"
	"github.com/productivityeng/orabbit/user"
	userEntities "github.com/productivityeng/orabbit/user/entities"
	log "github.com/sirupsen/logrus"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	database_mysql.Db.AutoMigrate(&brokerEntities.BrokerEntity{}, &userEntities.UserEntity{})

	gin.ForceConsoleColor()

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	broker.Routes(r, database_mysql.Db)
	user.Routes(r, database_mysql.Db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	log.Fatal(r.Run(":8082"))
}
