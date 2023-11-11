package repository

import (
	"context"
	"github.com/productivityeng/orabbit/virtualhost/entities"
)

type VirtualHostRepository interface {
	ListVirtualHosts(clusterId uint, pageSize int, pageNumber int, ctx context.Context) ([]entities.VirtualHost, error)
}
