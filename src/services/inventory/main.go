package main

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/cluster"
	"github.com/productivityeng/orabbit/core/core"
	"github.com/productivityeng/orabbit/docs"
	"github.com/productivityeng/orabbit/locker"
	"github.com/productivityeng/orabbit/queue"
	"github.com/productivityeng/orabbit/user"
	"github.com/productivityeng/orabbit/virtualhost"
	log "github.com/sirupsen/logrus"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	DependencyLocator := core.NewDependencyLocator()

	gin.ForceConsoleColor()

	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	cluster.Routes(r, DependencyLocator)
	user.Routes(r, DependencyLocator)
	queue.Routes(r, DependencyLocator)
	virtualhost.Routes(r, DependencyLocator)
	locker.Routes(r, DependencyLocator)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	log.Fatal(r.Run(":8082"))
}
