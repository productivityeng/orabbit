package resources

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/pkg/controllers/locker/dto"
	log "github.com/sirupsen/logrus"
)


func (ctrl *LockerController) parseCreateLockerBody(c *gin.Context) (*dto.CreateLockerRequest,error){
	var createLockerRequest dto.CreateLockerRequest
	err := c.BindJSON(&createLockerRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"error parsing request body"})
		return nil,err
	}
	return &createLockerRequest,nil
}

func (ctrl *LockerController) parseDisableLockerBody(c *gin.Context) (*dto.DisableLockerRequest,error){
	var disableLockerRequest dto.DisableLockerRequest
	err := c.BindJSON(&disableLockerRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"error parsing request body"})
		return nil,err
	}
	return &disableLockerRequest,nil
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