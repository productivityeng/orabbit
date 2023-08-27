package repository

import (
	"context"
	"errors"
	"github.com/productivityeng/orabbit/queue/entities"
	log "github.com/sirupsen/logrus"
)

func (repo *QueueRepositoryMysql) Delete(queueId uint, ctx context.Context) error {

	fields := log.Fields{"queueId": queueId}
	var queue = entities.QueueEntity{}
	queue.ID = queueId
	err := repo.Db.WithContext(ctx).First(&queue)
	if err.Error != nil {
		errorMsg := "cluster id cound't not be found"
		log.WithFields(fields).WithError(err.Error).Error(errorMsg)
		return errors.New(errorMsg)
	}
	log.WithFields(fields).Infof("queue founded, trying delete")
	err = repo.Db.Delete(&queue)
	if err.Error != nil {
		errorMsg := "fail to delete queue"
		log.WithFields(fields).WithError(err.Error).Error(errorMsg)

		return errors.New(errorMsg)
	}
	log.WithFields(fields).Info("queue deleted successfully")
	return nil
}
