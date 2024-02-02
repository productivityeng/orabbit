package core

import (
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/pkg/rabbitmq/exchange"
	"github.com/productivityeng/orabbit/pkg/rabbitmq/queue"
	"github.com/productivityeng/orabbit/pkg/rabbitmq/user"
	"github.com/productivityeng/orabbit/pkg/rabbitmq/virtualhost"
)


type DependencyLocator struct { 
	PrismaClient *db.PrismaClient
	ExchangeManagement contracts.ExchangeManagement
	VirtualHostManagement virtualhost.VirtualHostManagement
	UserManagement contracts.UserManagement
	QueueManagement contracts.QueueManagement
}

func NewDependencyLocator() (*DependencyLocator) { 
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	
	return &DependencyLocator{ PrismaClient: client, ExchangeManagement: exchange.NewExchangeManagement(),
	VirtualHostManagement: virtualhost.NewirtualHostManagement(),UserManagement: user.NewUserManagement(),QueueManagement: queue.NewQueueManagement(),}

}