package locker

import (
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




func (ctrl *LockerController) getEnabledLockerQueue(clusterId int,lockerType string,artifactId int,c *gin.Context) (*db.LockerQueueModel,error){
	locker,err := ctrl.DependencyLocator.PrismaClient.LockerQueue.FindFirst(
		db.LockerQueue.QueueID.Equals(artifactId),
		db.LockerQueue.Enabled.Equals(true),
	).Exec(c)
	return locker,err
}

func (ctrl *LockerController) getEnabledLockerUser(clusterId int,lockerType string,artifactId int,c *gin.Context) (*db.LockerUserModel,error){
	locker,err := ctrl.DependencyLocator.PrismaClient.LockerUser.FindFirst(
		db.LockerUser.UserID.Equals(artifactId),
		db.LockerUser.Enabled.Equals(true),
	).Exec(c)
	if err != nil {
		log.WithContext(c).WithError(err).Error("Error retrieving locker")
	}
	return locker,err
}

func (ctrl *LockerController) getLockerQueue(lockerId int,c *gin.Context) (*db.LockerQueueModel,error){
	locker,err := ctrl.DependencyLocator.PrismaClient.LockerQueue.FindUnique(
		db.LockerQueue.ID.Equals(lockerId),
	).Exec(c)
	return locker,err
}


func (ctrl *LockerController) getLockerUser(lockerId int,c *gin.Context) (*db.LockerUserModel,error){
	locker,err := ctrl.DependencyLocator.PrismaClient.LockerUser.FindUnique(
		db.LockerUser.ID.Equals(lockerId),
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

func (ctrl *LockerController) disableLockerUser(lockerId int,responsible string,c *gin.Context) (*db.LockerUserModel,error){
	log.WithFields(log.Fields{"lockerId":lockerId,"responsible":responsible}).Info("Disabling locker")
	updated,err := ctrl.DependencyLocator.PrismaClient.LockerUser.FindUnique(
		db.LockerUser.ID.Equals(lockerId),
		
	).Update(
		db.LockerUser.UserDisabled.Set(responsible),
		db.LockerUser.Enabled.Set(false),
	).Exec(c)
	if err != nil { return nil,err}
	return updated,nil
}
