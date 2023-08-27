package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/src/packages/common"
	common_rabbit "github.com/productivityeng/orabbit/src/packages/rabbitmq/common"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/user"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// ListUsers
// @Summary Retrieve a mirror user from broker
// @Schemes
// @Description Recovery the details of a specific mirror user that is already imported from the cluster
// @Tags User
// @Accept json
// @Produce json
// @Param clusterId path int true "Cluster id from where retrieve users"
// @Param params query common.PageParam true "Number of items in one page"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /{clusterId}/user [get]
func (userCtrl *UserControllerImpl) ListUsers(c *gin.Context) {

	var param common.PageParam

	err := c.BindQuery(&param)
	if err != nil {
		log.WithError(err).Error("Error trying to parse query params")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clusterIdParam := c.Param("clusterId")
	clusterId, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse brokerId Param")
		c.JSON(http.StatusBadRequest, "Error parsing brokerId from url route")
		return
	}
	log.WithField("parameter", clusterIdParam).Info("Looking for list of users")

	result, err := userCtrl.UserRepository.ListUsers(uint(clusterId), param.PageSize, param.PageNumber, c)

	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to retrieve users for the cluster")
		c.JSON(http.StatusBadRequest, "Fail to retrieve users for the cluster")
		return
	}

	c.JSON(http.StatusOK, result)
	return
}

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
// @Router /{clusterId}/user/usersfromcluster [get]
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

	var response []contracts.GetUserResponse

	for _, userFromCluster := range usersFromCluster {
		userResponse := contracts.GetUserResponse{
			Username:     userFromCluster.Name,
			PasswordHash: userFromCluster.PasswordHash,
			Id:           0,
		}

	loopUsersFromdb:
		for _, userFromDb := range usersFromDb {
			if userFromDb.Username == userResponse.Username {
				userResponse.Id = userFromDb.ID
				userResponse.IsRegistered = true
				break loopUsersFromdb
			}
		}

		response = append(response, userResponse)
	}

	c.JSON(http.StatusOK, response)
	return
}
