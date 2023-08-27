package repository

import (
	"context"
	"github.com/productivityeng/orabbit/queue/entities"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (repo QueueRepositoryMysql) Get(clusterId uint, queueId uint, ctx context.Context) (*entities.QueueEntity, error) {

	foundedQueue := &entities.QueueEntity{}
	result := repo.Db.Where(&entities.QueueEntity{ClusterID: clusterId, Model: gorm.Model{ID: queueId}}).Find(foundedQueue)
	if result.Error != nil {
		log.WithError(result.Error).WithFields(log.Fields{"clusterId": clusterId,
			"queueId": queueId}).Error("Fail to find queue")
		return nil, result.Error
	}

	return foundedQueue, nil
}
