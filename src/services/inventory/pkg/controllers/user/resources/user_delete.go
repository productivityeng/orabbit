package resources

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/models"
	log "github.com/sirupsen/logrus"
)

// DeleteUser PingExample godoc
// @Summary Delete a mirror user
// @Schemes
// @Description Delete a mirrored user from the registry, the user will not be deleted from the cluster
// @Tags User
// @Accept json
// @Produce json
// @Param userId path int true "User id registered"
// @Param clusterId path int true "Cluster id from where delete user"
// @Success 204
// @Failure 404
// @Failure 500
// @Router /{clusterId}/user/{userId} [delete]
func (ctrl *UserControllerImpl) DeleteUser(c *gin.Context) {
	
	
	clusterId, userId, err := ctrl.parseDeleteUserRequest(c)
	if err != nil { 
		return
	}
	log.WithField("userId", userId).Info("request parsed")

	userFromDb,err := ctrl.getUser(c, userId)
	if err != nil { 
		return
	}
	
	log.WithField("userFromDb", userFromDb).Info("userFromDb")

	err = ctrl.verifyIfUserIsLocked(c,userFromDb)
	if err != nil { 
		return 
	}

	log.WithField("userId", userId).Info("user is not locked")
	
	cluster, err := ctrl.getCluster(c, clusterId)

	if err != nil { 
		return
	}

	log.WithField("cluster", cluster).Info("cluster found")

	deleteUserRequest := contracts.DeleteUserRequest{
		RabbitAccess: models.GetRabbitMqAccess(cluster),
		Username:     userFromDb.Username,
	}
	err = ctrl.DependencyLocator.UserManagement.DeleteUser(deleteUserRequest, c)

	if err != nil {
		log.WithError(err).WithField("request", deleteUserRequest).Error("Erro ao deletar usuario no rabbit")
		c.JSON(http.StatusInternalServerError, "Erro ao deletar usuario na base")
		return
	}
	log.WithField("userId", userId).Info("user deleted")
	c.JSON(http.StatusNoContent, "Deleted")
}

func (controller *UserControllerImpl) parseDeleteUserRequest(c *gin.Context) (clusterId int, userId int, err error) {
	userIdParam := c.Param("userId")

	userIdConv, err := strconv.ParseInt(userIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("brokerId", userIdParam).Error("Fail to parse brokerId Param")
		c.JSON(http.StatusBadRequest, "Error parsing cluster from url route")
		return 0, 0, err
	}

	clusterIdParam := c.Param("clusterId")
	clusterIdConv, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("brokerId", clusterIdParam).Error("Fail to parse clusterId Param")
		c.JSON(http.StatusBadRequest, "Error parsing clusterId from url route")
		return 0, 0, err
	}
	return int(clusterIdConv), int(userIdConv), nil
}