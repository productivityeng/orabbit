package resources

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/pkg/controllers/locker/dto"
	log "github.com/sirupsen/logrus"
)

// DisableLocker
// @Summary Disable a specific locker from a artificat in a cluster based on lockerType and artifactId
// @Schemes
// @Description Disable a specific locker from a artificat in a cluster based on lockerType and artifactId
// @Tags Locker
// @Accept json
// @Produce json
// @Param DisableLockerRequest body dto.DisableLockerRequest true "Request"
// @Param clusterId path int true "Cluster id from where retrieve users"
// @Param lockerType path string true "Artifact name from where retrieve users"
// @Param lockerId path int true "Id of the locker to be disable"
// @Success 200
// @Failure 404
// @Failure 400 Some input parameter was wrong provided
// @Failure 500
// @Router /{clusterId}/locker/{lockerType}/{lockerId} [delete]
func (ctrl *LockerController) DisableLocker(c *gin.Context) {
	clusterId,lockerType,lockerId,err := ctrl.parseRouteParams(c)
	if err != nil { return}
	lockerRequest,err := ctrl.parseDisableLockerBody(c)
	if err != nil { return}
	log.WithFields(log.Fields{"clusterId":clusterId,"lockerType":lockerType,"lockerId":lockerId,"userResponsible":lockerRequest.Responsible}).Info("Parsed route params")

	switch lockerType {
		case "queue": { 
			ctrl.handleDisableLockerQueue(lockerId,*lockerRequest,c)
			return
		}
		case "user": {
			ctrl.handleDisableLockerUser(lockerId,*lockerRequest,c)
			return
		}

		case "exchange": {
			ctrl.handleDisableLockerExchange(lockerId,*lockerRequest,c)
			return
		}

		case "virtualhost": { 
			ctrl.handleDisableLockerVirtualHost(lockerId,*lockerRequest,c)
			return
		}

		default: {
			c.JSON(http.StatusBadRequest,gin.H{"message":"locker type not found"})
			return
		 }
	 }
}

func (ctrl *LockerController) handleDisableLockerQueue(lockerId int,lockerRequest dto.DisableLockerRequest,c *gin.Context) {
	_,err := ctrl.getLockerQueue(lockerId,c)
	if errors.Is(err,db.ErrNotFound){ 
		c.JSON(http.StatusNotFound,gin.H{"message":"locker not found"})
		return
	}else if err != nil { 
		c.JSON(http.StatusInternalServerError,gin.H{"message":"error retrieving locker"})
		return
	}

	locker,err := ctrl.disableLockerQueue(lockerId,lockerRequest.Responsible,c)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message":"error disabling locker","error":err.Error()})
	}
	log.WithFields(log.Fields{"locker":locker}).Info("Disabled locker")
	c.JSON(http.StatusOK,locker)
}


func (ctrl *LockerController) handleDisableLockerUser(lockerId int,lockerRequest dto.DisableLockerRequest,c *gin.Context) {
	_,err := ctrl.getLockerUser(lockerId,c)
	if errors.Is(err,db.ErrNotFound){ 
		c.JSON(http.StatusNotFound,gin.H{"message":"locker not found"})
		return
	}else if err != nil { 
		c.JSON(http.StatusInternalServerError,gin.H{"message":"error retrieving locker"})
		return
	}

	locker,err := ctrl.disableLockerUser(lockerId,lockerRequest.Responsible,c)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message":"error disabling locker","error":err.Error()})
	}
	log.WithFields(log.Fields{"locker":locker}).Info("Disabled locker")
	c.JSON(http.StatusOK,locker)
}


func (ctrl *LockerController) handleDisableLockerExchange(lockerId int,lockerRequest dto.DisableLockerRequest,c *gin.Context) {
	_,err := ctrl.getLockerExchange(lockerId,c)
	if errors.Is(err,db.ErrNotFound){ 
		c.JSON(http.StatusNotFound,gin.H{"message":"locker not found"})
		return
	}else if err != nil { 
		c.JSON(http.StatusInternalServerError,gin.H{"message":"error retrieving locker"})
		return
	}

	locker,err := ctrl.disableLockerExchange(lockerId,lockerRequest.Responsible,c)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message":"error disabling locker","error":err.Error()})
	}
	log.WithFields(log.Fields{"locker":locker}).Info("Disabled locker")
	c.JSON(http.StatusOK,locker)
}


func (ctrl *LockerController) handleDisableLockerVirtualHost(lockerId int,lockerRequest dto.DisableLockerRequest,c *gin.Context) {
	locker,err := ctrl.getLockerVirtualHost(lockerId,c)

	if errors.Is(err,db.ErrNotFound){ 
		c.JSON(http.StatusNotFound,gin.H{"message":"locker not found"})
		return
	}else if err != nil { 
		c.JSON(http.StatusInternalServerError,gin.H{"message":"error retrieving locker"})
		return
	}

	if locker.Enabled == false { 
		c.JSON(http.StatusConflict,gin.H{"message":"locker already disabled"})
		return
	}

	disabledLocker,err := ctrl.disableLockerVirtualHost(lockerId,lockerRequest.Responsible,c)
	if err != nil { 
		c.JSON(http.StatusInternalServerError,gin.H{"message":"error disabling locker","error":err.Error()})
		return
	}

	c.JSON(http.StatusNoContent,disabledLocker)

}





