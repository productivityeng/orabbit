package repository

import "github.com/productivityeng/orabbit/queue/entities"

func (repo QueueRepositoryMysql) Save(entity *entities.QueueEntity) error {

	tx := repo.Db.Save(entity)
	return tx.Error
}
