package repository

import (
	"context"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/queue/entities"
	log "github.com/sirupsen/logrus"
)

func (repo *QueueRepositoryMysql) List(clusterId uint, pageSize int, pageNumber int, ctx context.Context) (*contracts.PaginatedResult[entities.QueueEntity], error) {
	entryFields := log.Fields{"pageSize": pageSize, "pageNumber": pageNumber}
	var result contracts.PaginatedResult[entities.QueueEntity]
	result.PageSize = pageSize
	result.PageNumber = pageNumber

	offset := (pageNumber - 1) * pageSize

	var resultsFromDb []*entities.QueueEntity

	err := repo.Db.WithContext(ctx).Where(entities.QueueEntity{ClusterID: clusterId}).Offset(offset).Limit(pageSize).Find(&resultsFromDb).Error
	result.Result = resultsFromDb
	if err != nil {
		log.WithError(err).WithFields(entryFields).Error("error trying to query items for users")
		return nil, err
	}

	tx := repo.Db.WithContext(ctx).Model(&entities.QueueEntity{ClusterID: clusterId}).Count(&result.TotalItems)
	if tx.Error != nil {
		log.WithError(tx.Error).WithFields(entryFields).Error("error trying to get count items for users")
	}

	return &result, nil
}
