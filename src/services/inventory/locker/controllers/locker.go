package locker

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core/core"
	"github.com/productivityeng/orabbit/db"
	log "github.com/sirupsen/logrus"
)



type LockerController struct {
	DependencyLocator *core.DependencyLocator
}

func NewLockerController(DependencyLocator *core.DependencyLocator) *LockerController {
	return &LockerController{DependencyLocator: DependencyLocator}
}


func (ctrl *LockerController) parseRouteParams(c *gin.Context) (clusterId int,lockerType string,lockerId int,err error){
	clusterIdParam := c.Param("clusterId")
	clusterIdConv, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse clusterId Param")
		c.JSON(http.StatusBadRequest, "Error parsing clusterId from url route")
		return 0,"",0,err
	}

	lockerIdParam := c.Param("lockerId")
	lockerIdConv, err := strconv.ParseInt(lockerIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("lockerId", lockerIdParam).Error("Fail to parse clusterId Param")
		c.JSON(http.StatusBadRequest, "Error parsing clusterId from url route")
		return 0,"",0,err
	}

	lockerType = c.Param("lockerType")
	if lockerType == "" {
		log.WithError(err).WithField("lockerType", lockerType).Error("Fail to parse lockerType Param")
		c.JSON(http.StatusBadRequest, "Error parsing lockerType from url route")
		return 0,"",0,errors.New("Error parsing lockerType from url route")
	 }


	return int(clusterIdConv),lockerType,int(lockerIdConv),nil
}

func (ctrl *LockerController) getEnabledLockerQueue(clusterId int,lockerType string,artifactId int,c *gin.Context) (*db.LockerQueueModel,error){
	locker,err := ctrl.DependencyLocator.PrismaClient.LockerQueue.FindFirst(
		db.LockerQueue.QueueID.Equals(artifactId),
		db.LockerQueue.Enabled.Equals(true),
	).Exec(c)
	return locker,err
}

func (ctrl *LockerController) getLockerQueue(lockerId int,c *gin.Context) (*db.LockerQueueModel,error){
	locker,err := ctrl.DependencyLocator.PrismaClient.LockerQueue.FindUnique(
		db.LockerQueue.ID.Equals(lockerId),
	).Exec(c)
	return locker,err
}

func (ctrl *LockerController) disableLockerQueue(lockerId int,responsible string,c *gin.Context) (*db.LockerQueueModel,error){
	log.WithFields(log.Fields{"lockerId":lockerId,"responsible":responsible}).Info("Disabling locker")
	updated,err := ctrl.DependencyLocator.PrismaClient.LockerQueue.FindUnique(
		db.LockerQueue.ID.Equals(lockerId),
		
	).Update(
		db.LockerQueue.UserDisabled.Set(responsible),
		db.LockerQueue.Enabled.Set(false),
	).Exec(c)
	if err != nil { return nil,err}
	return updated,nil
}
