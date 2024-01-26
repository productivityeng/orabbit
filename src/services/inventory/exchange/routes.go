package exchange

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core/core"
	controller "github.com/productivityeng/orabbit/exchange/controllers"
)


var exchangeController *controller.ExchangeController
func Routes(routes *gin.Engine, dependencyLocator *core.DependencyLocator) *gin.RouterGroup{

	exchangeController = controller.NewExchangeController(dependencyLocator)
	userRouter := routes.Group("/:clusterId/exchange")
	userRouter.GET("/", exchangeController.ListAllExchanges)
	userRouter.POST("/", exchangeController.CreateExchange)
	userRouter.POST("/import", exchangeController.ImportExchange)
	userRouter.DELETE("/:exchangeId", exchangeController.DeleteExchange)

	return userRouter
}