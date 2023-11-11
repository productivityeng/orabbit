package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/productivityeng/orabbit/queue/entities"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (repo QueueRepositoryMysql) Get(clusterId uint, queueId uint, ctx context.Context) (*entities.QueueEntity, error) {

	var foundedQueue entities.QueueEntity
	result := repo.Db.Where(&entities.QueueEntity{ClusterId: clusterId, Model: gorm.Model{ID: queueId}}).Find(&foundedQueue)
	if result.Error != nil {
		log.WithError(result.Error).WithFields(log.Fields{"clusterId": clusterId,
			"queueId": queueId}).Error("Fail to find queue")
		return nil, result.Error
	}

	if foundedQueue.ID == 0 {
		return nil, errors.New(fmt.Sprintf("Queue with id %d not found for cluster with id %d", queueId, clusterId))
	}

	return &foundedQueue, nil
}
