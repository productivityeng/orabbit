package exchange

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core"
	"github.com/productivityeng/orabbit/pkg/controllers/exchange/resources"
)


var exchangeController *resources.ExchangeController
func Routes(routes *gin.Engine, dependencyLocator *core.DependencyLocator) *gin.RouterGroup{

	exchangeController = resources.NewExchangeController(dependencyLocator)
	userRouter := routes.Group("/:clusterId/exchange")
	userRouter.GET("/", exchangeController.ListAllExchanges)
	userRouter.POST("/", exchangeController.CreateExchange)
	userRouter.POST("/import", exchangeController.ImportExchange)
	userRouter.DELETE("/:exchangeId", exchangeController.DeleteExchange)
	userRouter.POST("/:exchangeId/syncronize", exchangeController.SyncronizeExchange)

	return userRouter
}