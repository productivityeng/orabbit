package core

import (
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/rabbitmq/exchange"
)


type DependencyLocator struct { 
	PrismaClient *db.PrismaClient
	ExchangeManagement contracts.ExchangeManagement

}

func NewDependencyLocator() (*DependencyLocator) { 
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	
	return &DependencyLocator{ PrismaClient: client, ExchangeManagement: exchange.NewExchangeManagement()}

}