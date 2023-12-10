package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/cluster/repository"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/virtualhost"
	repository2 "github.com/productivityeng/orabbit/virtualhost/repository"
)

type VirtualHostController interface {
	ListVirtualHost(c *gin.Context)
	ImportOrCreateVirtualHost(c *gin.Context)
}

type VirtualHostControllerImpl struct {
	VirtualHostManagement virtualhost.VirtualHostManagement
	ClusterRepository     repository.ClusterRepositoryInterface
	VirtualHostRepository repository2.VirtualHostRepository
}

func NewVirtualHostControllerImpl(VirtualHostManagement virtualhost.VirtualHostManagement,
	ClusterRepository repository.ClusterRepositoryInterface,
	hostRepository repository2.VirtualHostRepository) VirtualHostControllerImpl {
	return VirtualHostControllerImpl{
		VirtualHostManagement: VirtualHostManagement,
		ClusterRepository:     ClusterRepository,
		VirtualHostRepository: hostRepository,
	}
}
