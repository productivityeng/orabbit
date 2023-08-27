package repository

import (
	"context"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/queue/entities"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type QueueRepository interface {
	// Save  store a new queue in storage with provided parameter
	Save(entity *entities.QueueEntity) error
	// List retrieve a list of broker with paginated options
	List(brokerId uint, pageSize int, pageNumber int, ctx context.Context) (*contracts.PaginatedResult[entities.QueueEntity], error)
	// Delete soft delete a broker with a provided clusterId
	Delete(queueId uint, ctx context.Context) error
	// Get retrieve a queue with a provided clusterId
	Get(clusterId uint, userId uint, ctx context.Context) (*entities.QueueEntity, error)
}

type QueueRepositoryMysql struct {
	Db *gorm.DB
}

func NewQueueRepositoryMySql(Db *gorm.DB) *QueueRepositoryMysql {
	err := Db.AutoMigrate(entities.QueueEntity{})
	if err != nil {
		log.WithError(err).Panic("erro migrating entity queue")
	}
	return &QueueRepositoryMysql{
		Db: Db,
	}
}
