package repository

import (
	"context"
	"github.com/productivityeng/orabbit/queue/entities"
	log "github.com/sirupsen/logrus"
)

func (repo *QueueRepositoryMysql) List(clusterId uint, ctx context.Context) (entities.QueueEntityList, error) {

	var resultsFromDb entities.QueueEntityList

	err := repo.Db.WithContext(ctx).Where(entities.QueueEntity{ClusterId: clusterId}).Find(&resultsFromDb).Error
	if err != nil {
		log.WithError(err).Error("error trying to query items for users")
		return nil, err
	}

	return resultsFromDb, nil
}
