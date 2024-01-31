package main

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core"
	"github.com/productivityeng/orabbit/docs"
	"github.com/productivityeng/orabbit/pkg/cluster"
	"github.com/productivityeng/orabbit/pkg/exchange"
	"github.com/productivityeng/orabbit/pkg/locker"
	"github.com/productivityeng/orabbit/pkg/queue"
	"github.com/productivityeng/orabbit/pkg/user"
	"github.com/productivityeng/orabbit/pkg/virtualhost"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	DependencyLocator := core.NewDependencyLocator()
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
		ForceQuote: true,
	})
	gin.ForceConsoleColor()

	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	cluster.Routes(r, DependencyLocator)
	user.Routes(r, DependencyLocator)
	queue.Routes(r, DependencyLocator)
	virtualhost.Routes(r, DependencyLocator)
	locker.Routes(r, DependencyLocator)
	exchange.Routes(r, DependencyLocator)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	log.Fatal(r.Run(":8082"))
}
