package main

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/cluster"
	brokerEntities "github.com/productivityeng/orabbit/cluster/entities"
	database_mysql "github.com/productivityeng/orabbit/database"
	"github.com/productivityeng/orabbit/docs"
	lockerEntities "github.com/productivityeng/orabbit/locker/entities"
	"github.com/productivityeng/orabbit/queue"
	queueEntities "github.com/productivityeng/orabbit/queue/entities"
	"github.com/productivityeng/orabbit/user"
	userEntities "github.com/productivityeng/orabbit/user/entities"
	"github.com/productivityeng/orabbit/virtualhost"
	virtualHostEntities "github.com/productivityeng/orabbit/virtualhost/entities"
	log "github.com/sirupsen/logrus"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	database_mysql.Db.AutoMigrate(&brokerEntities.ClusterEntity{},
		&userEntities.UserEntity{}, &queueEntities.QueueEntity{}, &lockerEntities.LockerEntity{},
		&virtualHostEntities.VirtualHost{})

	gin.ForceConsoleColor()

	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	cluster.Routes(r, database_mysql.Db)
	user.Routes(r, database_mysql.Db)
	queue.Routes(r, database_mysql.Db)
	virtualhost.Routes(r, database_mysql.Db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	log.Fatal(r.Run(":8082"))
}
