package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core/core"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/virtualhost"
)

type VirtualHostController interface {
	ListVirtualHost(c *gin.Context)
	ImportOrCreateVirtualHost(c *gin.Context)
}

type VirtualHostControllerImpl struct {
	VirtualHostManagement virtualhost.VirtualHostManagement
	DependencyLocator     *core.DependencyLocator	
}

func NewVirtualHostControllerImpl(VirtualHostManagement virtualhost.VirtualHostManagement,
	DependencyLocator *core.DependencyLocator) VirtualHostControllerImpl {
	return VirtualHostControllerImpl{
		VirtualHostManagement: VirtualHostManagement,
	}
}
