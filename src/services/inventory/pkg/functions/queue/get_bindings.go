package functions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core"
	"github.com/productivityeng/orabbit/db"
	log "github.com/sirupsen/logrus"
)

// GetBindingsWhereQueueIsDest looks for bindings where the queueId is the destination
func GetBindingsWhereQueueIsDest(DependencyLocator *core.DependencyLocator, c *gin.Context,queueId int) ([]db.BindingExchangeToQueueModel, error) {
	bindings,err := DependencyLocator.PrismaClient.BindingExchangeToQueue.FindMany(
		db.BindingExchangeToQueue.DestinationQueueID.Equals(queueId),
	).With(
		db.BindingExchangeToQueue.SourceExchange.Fetch(),
		
	).Exec(c)
	if err != nil {
		log.WithContext(c).WithError(err).Error("Fail to find bindings")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil,err
	}
	return bindings,nil
}