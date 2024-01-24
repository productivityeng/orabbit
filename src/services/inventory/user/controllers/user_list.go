package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/db"
	common_rabbit "github.com/productivityeng/orabbit/src/packages/rabbitmq/common"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/user"
	log "github.com/sirupsen/logrus"
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

	clusterId,err := userCtrl.parseUserParams(c)
	if err != nil { return }

	fields := log.Fields{"clusterId": clusterId}

	log.WithFields(fields).Info("Looking for rabbitmq cluster")

	cluster,err := userCtrl.getCluster(c, int(clusterId))
	
	if err != nil { 
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

	log.WithField("qtdUsersFromCluster", len(usersFromCluster)).Info("Users founded in cluster")

	log.WithFields(fields).Info("Looking for rabbitmq users from ostern")

	usersFromDb, err := userCtrl.DependencyLocator.PrismaClient.User.FindMany(db.User.ClusterID.Equals(cluster.ID)).Exec(c)

	if err != nil {
		log.WithError(err).WithField("clusterId", cluster).Error("Fail to retrieve users for the cluster")
		c.JSON(http.StatusBadRequest, "Fail to retrieve users for the cluster")
		return
	}

	log.WithFields(fields).Info("Preparing response")

	response := make(contracts.GetUserResponseList,0)

	for _, userFromCluster := range usersFromCluster {
		userResponse := contracts.GetUserResponse{
			Username:     userFromCluster.Name,
			PasswordHash: userFromCluster.PasswordHash,
			Id:           0,
			IsInCluster:  true,
			IsInDatabase: false,
		}

		userEqual,_ := userCtrl.DependencyLocator.PrismaClient.User.FindUnique(db.User.UniqueUsernameClusterid(db.User.Username.Equals(userFromCluster.Name),db.User.ClusterID.Equals(cluster.ID))).Exec(c)
		
		if userEqual != nil {
			userResponse.Id = userEqual.ID
			userResponse.IsInDatabase = true
		}

		response = append(response, userResponse)

		
		
		
	}

	for _, userFromDb := range usersFromDb {
		if !response.UserInListByName(userFromDb.Username) {
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


func (controller *UserControllerImpl) parseUserParams(c *gin.Context) (clusterId int,  err error) {
	clusterIdParam := c.Param("clusterId")
	clusterIdConv, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse clusterId Param")
		c.JSON(http.StatusBadRequest, "Error parsing clusterId from url route")
		return  0, err
	}

	return int(clusterIdConv), nil
	
}	
