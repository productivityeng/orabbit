package repository

import (
	"github.com/productivityeng/orabbit/broker/entities"
)

type BrokerRepositoryInterface interface {
	CreateBroker(broker *entities.BrokerEntity) (*entities.BrokerEntity, error)
	ListBroker(pageSize int, pageNumber int) ([]entities.BrokerEntity, error)
}
