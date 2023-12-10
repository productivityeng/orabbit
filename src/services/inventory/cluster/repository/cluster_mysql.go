package repository

import (
	"context"
	"errors"
	"github.com/productivityeng/orabbit/cluster/entities"
	"github.com/productivityeng/orabbit/contracts"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ClusterRepositoryMysqlImpl struct {
	Db *gorm.DB
}

func NewClusterMysqlRepositoryImpl(Db *gorm.DB) *ClusterRepositoryMysqlImpl {
	Db.AutoMigrate(&entities.ClusterEntity{})
	/* if err != nil {
		log.WithError(err).Fatal("erro migrating entity broker")
	} */
	return &ClusterRepositoryMysqlImpl{
		Db: Db,
	}
}

// CreateCluster create a registry for a new cluster in the database
func (repo *ClusterRepositoryMysqlImpl) CreateCluster(entityToCreate *entities.ClusterEntity) (*entities.ClusterEntity, error) {

	tx := repo.Db.Save(entityToCreate)
	if tx.Error != nil {
		log.WithError(tx.Error).WithField("request", entityToCreate).Error("Erro when trying save")
		tx.Rollback()
		return nil, tx.Error
	}
	return entityToCreate, nil
}

// ListCluster retrieve paginated result of brokers in mysql store
func (repo *ClusterRepositoryMysqlImpl) ListCluster(pageSize int, pageNumber int) (*contracts.PaginatedResult[entities.ClusterEntity], error) {

	entryFields := log.Fields{"pageSize": pageSize, "pageNumber": pageNumber}
	var result contracts.PaginatedResult[entities.ClusterEntity]
	result.PageSize = pageSize
	result.PageNumber = pageNumber

	offset := (pageNumber - 1) * pageSize

	err := repo.Db.Where("").Offset(offset).Limit(pageSize).Find(&result.Result).Error

	if err != nil {
		log.WithError(err).WithFields(entryFields).Error("error trying to query items for brokers")
		return nil, err
	}

	tx := repo.Db.Model(&entities.ClusterEntity{}).Count(&result.TotalItems)
	if tx.Error != nil {
		log.WithError(tx.Error).WithFields(entryFields).Error("error trying to get count items for brokers")
	}
	return &result, nil
}

func (repo *ClusterRepositoryMysqlImpl) DeleteCluster(clusterId uint, ctx context.Context) error {
	fields := log.Fields{"clusterId": clusterId}
	var broker = entities.ClusterEntity{Model: gorm.Model{ID: clusterId}}
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

func (repo *ClusterRepositoryMysqlImpl) GetCluster(clusterId uint, ctx context.Context) (*entities.ClusterEntity, error) {
	fields := log.Fields{"clusterId": clusterId}
	var broker = entities.ClusterEntity{Model: gorm.Model{ID: clusterId}}
	err := repo.Db.WithContext(ctx).Unscoped().First(&broker)
	if err.Error != nil {
		errorMsg := "broker id cound't not be found"
		log.WithFields(fields).WithError(err.Error).Error(errorMsg)
		return nil, errors.New(errorMsg)
	}

	return &broker, nil
}

func (repo *ClusterRepositoryMysqlImpl) CheckIfHostIsAlreadyRegisted(host string, port int32, ctx context.Context) bool {
	fields := log.Fields{"host": host}
	log.WithFields(fields).Info("Checking if host already registered")
	count := int64(0)
	repo.Db.WithContext(ctx).Model(&entities.ClusterEntity{}).Unscoped().Where("Host = ? and Port = ?", host, port).Count(&count)
	return count > 0
}
