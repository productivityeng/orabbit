package resources

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/controllers/user/dto"
	"github.com/productivityeng/orabbit/db"
	log "github.com/sirupsen/logrus"
)

// FindUser godoc
// @Summary Retrieve a mirror user from broker
// @Schemes
// @Description Recovery the details of a specific mirror user that is already imported from the cluster
// @Tags User
// @Accept json
// @Produce json
// @Param userId path int true "User id registered"
// @Param clusterId path int true "Cluster from where the user is"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /{clusterId}/user/{userId} [get]
func (controller *UserControllerImpl) FindUser(c *gin.Context) {
	clusterId,userId,err := controller.parseFindUserRequest(c)
	if err != nil {
		return
	}

	result,err  := controller.DependencyLocator.PrismaClient.User.FindUnique(db.User.UniqueIDClusterid(db.User.ID.Equals(userId), db.User.ClusterID.Equals(clusterId))).Exec(c)

	if errors.Is(err,db.ErrNotFound) {
		log.WithError(err).WithContext(c).Error("Fail to retrieve user by id")
		c.JSON(http.StatusNotFound, gin.H{"error": "[USER_NOT_FOUND]"})
		return
	}
	
	c.JSON(http.StatusOK, dto.GetUserResponseFromUserEntity(result))
	return
}

func (controller *UserControllerImpl) parseFindUserRequest(c *gin.Context) (clusterId int, userId int,err error) {
	userIdParam := c.Param("userId")
	userIdConv, err := strconv.ParseInt(userIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("brokerId", userIdParam).Error("Fail to parse brokerId Param")
		c.JSON(http.StatusBadRequest, "Error parsing brokerId from url route")
		return 0,0,err
	}

	clusterIdParam := c.Param("clusterId")
	clusterIdConv, err := strconv.ParseInt(userIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("brokerId", clusterIdParam).Error("Fail to parse clusterId Param")
		c.JSON(http.StatusBadRequest, "Error parsing clusterId from url route")
		return 0,0,err
	}

	return int(clusterIdConv),int(userIdConv),nil
}
