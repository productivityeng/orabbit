package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/cluster/repository"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/virtualhost"
)

type VirtualHostController interface {
	ListVirtualHost(c *gin.Context)
	ImportOrCreateVirtualHost(c *gin.Context)
}

type VirtualHostControllerImpl struct {
	VirtualHostManagement virtualhost.VirtualHostManagement
	ClusterRepository     repository.ClusterRepositoryInterface
}

func NewVirtualHostControllerImpl(VirtualHostManagement virtualhost.VirtualHostManagement, ClusterRepository repository.ClusterRepositoryInterface) VirtualHostControllerImpl {
	return VirtualHostControllerImpl{
		VirtualHostManagement: VirtualHostManagement,
		ClusterRepository:     ClusterRepository,
	}
}
