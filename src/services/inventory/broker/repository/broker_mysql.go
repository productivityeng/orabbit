package repository

import (
	"github.com/productivityeng/orabbit/broker/entities"
	"github.com/productivityeng/orabbit/contracts"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BrokerRepositoryMysqlImpl struct {
	Db *gorm.DB
}

func NewBrokerMysqlImpl(Db *gorm.DB) *BrokerRepositoryMysqlImpl {

	return &BrokerRepositoryMysqlImpl{
		Db: Db,
	}
}

func (repo *BrokerRepositoryMysqlImpl) CreateBroker(request contracts.CreateBrokerRequest) (*entities.BrokerEntity, error) {
	entityToCreate := &entities.BrokerEntity{Name: request.Name, Host: request.Host, User: request.User, Password: request.Password, Port: request.Port,
		Description: request.Description}
	tx := repo.Db.Save(entityToCreate)
	if tx.Error != nil {
		log.WithError(tx.Error).WithField("request", request).Error("Erro when trying save")
		return nil, tx.Error
	}
	return entityToCreate, nil
}

func (repo *BrokerRepositoryMysqlImpl) ListBroker(pageSize int, pageNumber int) ([]*entities.BrokerEntity, error) {
	return nil, nil
}
