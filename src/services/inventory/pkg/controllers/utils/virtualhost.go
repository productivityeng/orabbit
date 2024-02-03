package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/db"
)

// VerifyIfVirtualHostIsLockedById verifies if a virtualhost is locked.
// returns bad request response if the virtualhost is locked.
// returns nil if the virtualhost is not locked.
func VerifyIfVirtualHostIsLockedById(prismaClient *db.PrismaClient, virtualHostId int,c *gin.Context) error {
	result,err := prismaClient.LockerVirtualHost.FindFirst(
		db.LockerVirtualHost.VirtualHostID.Equals(virtualHostId),
		db.LockerVirtualHost.Enabled.Equals(true),
	).With(
		db.LockerVirtualHost.VirtualHost.Fetch(),
	).Exec(c)

	if errors.Is(err, db.ErrNotFound) { 
		return nil
	}else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("VirtualHost of this queue is locked by %s",result.UserResponsibleEmail)})
		return errors.New("VirtualHost is locked")
	}

	return nil
}



