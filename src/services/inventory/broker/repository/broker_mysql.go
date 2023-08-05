package repository

import (
	"context"
	"errors"
	"github.com/productivityeng/orabbit/broker/entities"
	"github.com/productivityeng/orabbit/contracts"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BrokerRepositoryMysqlImpl struct {
	Db *gorm.DB
}

func NewBrokerMysqlImpl(Db *gorm.DB) *BrokerRepositoryMysqlImpl {
	Db.AutoMigrate(&entities.BrokerEntity{})
	/* if err != nil {
		log.WithError(err).Fatal("erro migrating entity broker")
	} */
	return &BrokerRepositoryMysqlImpl{
		Db: Db,
	}
}

func (repo *BrokerRepositoryMysqlImpl) CreateBroker(entityToCreate *entities.BrokerEntity) (*entities.BrokerEntity, error) {

	tx := repo.Db.Save(entityToCreate)
	if tx.Error != nil {
		log.WithError(tx.Error).WithField("request", entityToCreate).Error("Erro when trying save")
		tx.Rollback()
		return nil, tx.Error
	}
	return entityToCreate, nil
}

// ListBroker retrieve paginated result of brokers in mysql store
func (repo *BrokerRepositoryMysqlImpl) ListBroker(pageSize int, pageNumber int) (*contracts.PaginatedResult[entities.BrokerEntity], error) {

	entryFields := log.Fields{"pageSize": pageSize, "pageNumber": pageNumber}
	var result contracts.PaginatedResult[entities.BrokerEntity]
	result.PageSize = pageSize
	result.PageNumber = pageNumber

	offset := (pageNumber - 1) * pageSize

	err := repo.Db.Offset(offset).Limit(pageSize).Find(&result.Result).Error

	if err != nil {
		log.WithError(err).WithFields(entryFields).Error("error trying to query items for brokers")
		return nil, err
	}

	tx := repo.Db.Model(&entities.BrokerEntity{}).Count(&result.TotalItems)
	if tx.Error != nil {
		log.WithError(tx.Error).WithFields(entryFields).Error("error trying to get count items for brokers")
	}
	return &result, nil
}

func (repo *BrokerRepositoryMysqlImpl) DeleteBroker(brokerId int32, ctx context.Context) error {
	fields := log.Fields{"brokerId": brokerId}
	var broker = entities.BrokerEntity{Id: brokerId}
	err := repo.Db.First(&broker)
	if err.Error != nil {
		errorMsg := "broker id cound't not be found"
		log.WithFields(fields).WithError(err.Error).Error(errorMsg)
		return errors.New(errorMsg)
	}
	log.WithFields(fields).Infof("broker founded, trying delete")
	err = repo.Db.Delete(&broker)
	if err.Error != nil {
		errorMsg := "fail to delete broker"
		log.WithFields(fields).WithError(err.Error).Error(errorMsg)

		return errors.New(errorMsg)
	}
	log.WithFields(fields).Info("broker deleted successfully")
	return nil
}

func (repo *BrokerRepositoryMysqlImpl) GetBroker(brokerId int32, ctx context.Context) (*entities.BrokerEntity, error) {
	fields := log.Fields{"brokerId": brokerId}
	var broker = entities.BrokerEntity{Id: brokerId}
	err := repo.Db.WithContext(ctx).First(&broker)
	if err.Error != nil {
		errorMsg := "broker id cound't not be found"
		log.WithFields(fields).WithError(err.Error).Error(errorMsg)
		return nil, errors.New(errorMsg)
	}

	return &broker, nil
}

func (repo *BrokerRepositoryMysqlImpl) CheckIfHostIsAlreadyRegisted(host string, port int32, ctx context.Context) bool {
	fields := log.Fields{"host": host}
	log.WithFields(fields).Info("Checking if host already registered")
	count := int64(0)
	repo.Db.WithContext(ctx).Model(&entities.BrokerEntity{}).Where("Host = ? and Port = ?", host, port).Count(&count).Limit(1)
	return count > 0
}
