package repository

import "github.com/productivityeng/orabbit/queue/entities"

func (repo *QueueRepositoryMysql) Save(entity *entities.QueueEntity) error {

	return repo.Save(entity)
}
