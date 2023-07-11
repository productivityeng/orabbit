package repository

import (
	"github.com/productivityeng/orabbit/broker/entities"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BrokerRepositoryMysqlImpl struct {
	Db *gorm.DB
}

func NewBrokerMysqlImpl(Db *gorm.DB) *BrokerRepositoryMysqlImpl {
	Db.AutoMigrate(&entities.BrokerEntity{})
	return &BrokerRepositoryMysqlImpl{
		Db: Db,
	}
}

func (repo *BrokerRepositoryMysqlImpl) CreateBroker(entityToCreate *entities.BrokerEntity) (*entities.BrokerEntity, error) {
	tx := repo.Db.Save(entityToCreate)
	if tx.Error != nil {
		//log.WithError(tx.Error).WithField("request", entityToCreate).Error("Erro when trying save")
		return nil, tx.Error
	}
	return entityToCreate, nil
}

func (repo *BrokerRepositoryMysqlImpl) ListBroker(pageSize int, pageNumber int) ([]entities.BrokerEntity, error) {
	var brokers []entities.BrokerEntity
	offset := (pageNumber - 1) * pageSize

	err := repo.Db.Offset(offset).Limit(pageSize).Find(&brokers).Error

	if err != nil {
		log.WithError(err).WithFields(log.Fields{"pageSize": pageSize, "pageNumber": pageNumber})
		return nil, err
	}

	repo.Db.Scan(&brokers)
	return brokers, nil
}
