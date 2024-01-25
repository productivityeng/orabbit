package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/db"
	common_rabbit "github.com/productivityeng/orabbit/src/packages/rabbitmq/common"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/user"
	"github.com/productivityeng/orabbit/user/dto"
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

	response := userCtrl.buildUserListResponse(usersFromCluster,usersFromDb,clusterId,c)

	c.JSON(http.StatusOK, response)
}

// ListUsersFromCluster
//Merge users from rabbitmq cluster and ostern database in a single response
func (controller *UserControllerImpl) buildUserListResponse(usersFromCluster []user.ListUserResult,usersFromDb []db.UserModel,clusterId int,c *gin.Context) dto.GetUserResponseList {
	response := make(dto.GetUserResponseList,0)

	for _, userFromCluster := range usersFromCluster {
		userResponse := dto.GetUserResponse{
			ClusterId:   clusterId,
			Username:     userFromCluster.Name,
			PasswordHash: userFromCluster.PasswordHash,
			Id:           0,
			IsInCluster:  true,
			IsInDatabase: false,
			Lockers: 	make([]db.LockerUserModel,0),
		}

		userEqual,_ := controller.DependencyLocator.PrismaClient.User.FindUnique(db.User.UniqueUsernameClusterid(db.User.Username.Equals(userFromCluster.Name),db.User.ClusterID.Equals(clusterId))).
		With(db.User.LockerUser.Fetch()).
		Exec(c)
		
		if userEqual != nil {
			userResponse.Id = userEqual.ID
			userResponse.IsInDatabase = true
			userResponse.Lockers = userEqual.LockerUser()
		}
		
		response = append(response, userResponse)
	}

	for _, userFromDb := range usersFromDb {
		if !response.UserInListByName(userFromDb.Username) {
			response = append(response, dto.GetUserResponse{
				Id:           userFromDb.ID,
				ClusterId: clusterId,
				Username:     userFromDb.Username,
				PasswordHash: userFromDb.PasswordHash,
				IsInCluster:  false,
				IsInDatabase: true,
			})
		}
	}

	return response
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
