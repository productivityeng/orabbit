package controllers

import (
	"github.com/productivityeng/orabbit/core/core"
)




type ExchangeController struct { 
	DependencyLocator *core.DependencyLocator
}

func NewExchangeController(DependencyLocator *core.DependencyLocator) *ExchangeController {
	return &ExchangeController{ DependencyLocator: DependencyLocator}
}