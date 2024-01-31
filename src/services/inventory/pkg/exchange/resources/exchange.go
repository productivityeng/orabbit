package resources

import (
	"github.com/productivityeng/orabbit/core"
)




type ExchangeController struct { 
	DependencyLocator *core.DependencyLocator
}

func NewExchangeController(DependencyLocator *core.DependencyLocator) *ExchangeController {
	return &ExchangeController{ DependencyLocator: DependencyLocator}
}