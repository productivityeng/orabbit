package repository

import (
	"context"
	userEntities "github.com/productivityeng/orabbit/user/entities"
	"github.com/productivityeng/orabbit/virtualhost/entities"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type VirtualHostRepositoryMysql struct {
	Db *gorm.DB
}

func NewVirtualHostRepositoryMysql(db *gorm.DB) *VirtualHostRepositoryMysql {
	return &VirtualHostRepositoryMysql{Db: db}
}

func (repo VirtualHostRepositoryMysql) ListVirtualHosts(clusterId uint, pageSize int, pageNumber int, ctx context.Context) ([]entities.VirtualHost, error) {
	entryFields := log.Fields{"pageSize": pageSize, "pageNumber": pageNumber}

	offset := (pageNumber - 1) * pageSize

	var resultsFromDb []entities.VirtualHost

	err := repo.Db.WithContext(ctx).Where(userEntities.UserEntity{ClusterId: clusterId}).Offset(offset).Limit(pageSize).Find(&resultsFromDb).Error

	if err != nil {
		log.WithError(err).WithFields(entryFields).Error("error trying to query items for virtual hosts")
		return nil, err
	}

	return resultsFromDb, nil
}
