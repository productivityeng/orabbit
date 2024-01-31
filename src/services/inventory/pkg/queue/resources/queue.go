package resources

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/core"
	"github.com/productivityeng/orabbit/db"
)

func NewQueueController(DependencyLocator *core.DependencyLocator, management contracts.QueueManagement,
	) QueueControllerImpl {
	return QueueControllerImpl{ QueueManagement: management, DependencyLocator: DependencyLocator}
}

type QueueController interface {
	ListQueuesFromCluster(c *gin.Context)
	ImportQueueFromCluster(c *gin.Context)
	SyncronizeQueue(c *gin.Context)
	RemoveQueueFromCluster(c *gin.Context)
	FindQueue(c *gin.Context) 
}

type QueueControllerImpl struct {
	DependencyLocator *core.DependencyLocator
	QueueManagement   contracts.QueueManagement
}

func(ctrl QueueControllerImpl) verifyIfQueueIsLocked(queueId int,c *gin.Context) error {
	result,err :=ctrl.DependencyLocator.PrismaClient.LockerQueue.FindFirst(
		db.LockerQueue.QueueID.Equals(queueId),
		db.LockerQueue.Enabled.Equals(true),
	).Exec(c)

	if errors.Is(err, db.ErrNotFound) { 
		return nil
	}else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Queue is locked"})
		return errors.New("queue is locked")
	}

	return nil
}