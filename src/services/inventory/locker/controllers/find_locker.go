package locker

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// FindLocker
// @Summary Retrieve a specific locker from a artificat in a cluster based on lockerType and artifactId
// @Schemes
// @Description Retrieve a specific locker from a artificat in a cluster based on lockerType and artifactId
// @Tags Locker
// @Accept json
// @Produce json
// @Param clusterId path int true "Cluster id from where retrieve users"
// @Param lockerType path string true "Artifact name from where retrieve users"
// @Param artifactId path int true "Artifact id from where retrieve users"
// @Success 200
// @Failure 404
// @Failure 400 Some input parameter was wrong provided
// @Failure 500
// @Router /{clusterId}/locker/{lockerType}/{artifactId} [get]
func (ctrl *LockerController) FindLocker(c *gin.Context) {
	clusterId,lockerType,artifactId,err := ctrl.parseRouteParams(c)
	if err != nil { return}
	
	logrus.WithFields(logrus.Fields{"clusterId":clusterId,"lockerType":lockerType,"artifactId":artifactId}).Info("Parsed route params")

	switch lockerType {
		case "queue": { 
			enabledLocker,err := ctrl.getEnabledLockerQueue(clusterId,lockerType,artifactId,c)
			if err != nil { 
				c.JSON(http.StatusNotFound,gin.H{"message":"enabled locker not found"})
				return
	 		}

			c.JSON(http.StatusOK,enabledLocker)
			return
		}

		default: {
			c.JSON(http.StatusBadRequest,gin.H{"message":"locker type not found"})
		 }
	 }
	
	
}



