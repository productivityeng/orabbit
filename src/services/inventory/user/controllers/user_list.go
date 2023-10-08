package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/contracts"
	common_rabbit "github.com/productivityeng/orabbit/src/packages/rabbitmq/common"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/user"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// ListUsersFromCluster
// @Summary Retrieve all users from rabbitmq cluster
// @Schemes
// @Description Retrieve all users that exist on rabbit cluster. Event if it its registered in ostern
// @Tags User
// @Accept json
// @Produce json
// @Param clusterId path int true "Cluster id from where retrieve users"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /{clusterId}/user [get]
func (userCtrl *UserControllerImpl) ListUsersFromCluster(c *gin.Context) {

	clusterIdParam := c.Param("clusterId")
	clusterId, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse brokerId Param")
		c.JSON(http.StatusBadRequest, "Error parsing brokerId from url route")
		return
	}

	fields := log.Fields{"clusterId": clusterId}

	log.WithFields(fields).Info("Looking for rabbitmq cluster")
	cluster, err := userCtrl.ClusterRepository.GetCluster(uint(clusterId), c)

	if err != nil {
		log.WithError(err).WithFields(fields).Error("Fail to retrieve cluster")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.WithFields(fields).Info("Looking for rabbitmq users from cluster")

	usersFromCluster, err := userCtrl.UserManagement.ListAllUser(user.ListAllUsersRequest{RabbitAccess: common_rabbit.RabbitAccess{
		Host:     cluster.Host,
		Port:     cluster.Port,
		Username: cluster.User,
		Password: cluster.Password,
	}})

	if err != nil {
		usersFromClusterError := errors.New("fail to retrieve users from rabbitmq cluster")
		log.WithError(err).WithField("clusterId", clusterId).Error(usersFromClusterError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": usersFromClusterError.Error()})
	}
	log.WithFields(fields).Info("Looking for rabbitmq users from ostern")

	usersFromDb, err := userCtrl.UserRepository.ListAllRegisteredUsers(cluster.ID, c)

	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to retrieve users for the cluster")
		c.JSON(http.StatusBadRequest, "Fail to retrieve users for the cluster")
		return
	}

	log.WithFields(fields).Info("Preparing response")

	var response contracts.GetUserResponseList

	for _, userFromCluster := range usersFromCluster {
		userResponse := contracts.GetUserResponse{
			Username:     userFromCluster.Name,
			PasswordHash: userFromCluster.PasswordHash,
			Id:           0,
			IsInCluster:  true,
			IsInDatabase: false,
		}
		if userEqual := usersFromDb.UserInListByName(userFromCluster.Name); userEqual != nil {
			userResponse.Id = userEqual.ID
			userResponse.IsInDatabase = true
		}
		response = append(response, userResponse)
	}

	for _, userFromDb := range usersFromDb {
		if response.UserInListByName(userFromDb.Username) == false {
			response = append(response, contracts.GetUserResponse{
				Id:           userFromDb.ID,
				Username:     userFromDb.Username,
				PasswordHash: userFromDb.PasswordHash,
				IsInCluster:  false,
				IsInDatabase: true,
			})
		}
	}

	c.JSON(http.StatusOK, response)
	return
}
