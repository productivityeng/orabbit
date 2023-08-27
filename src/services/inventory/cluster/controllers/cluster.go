package cluster

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/productivityeng/orabbit/cluster/entities"
	"github.com/productivityeng/orabbit/cluster/repository"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/core/validators"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type FindClusterByHost struct {
	Host string `json:"Host" binding:"required"`
	Port int32  `json:"Port" binding:"required"`
}
type ClusterController interface {
	GetCluster(c *gin.Context)
	ListClusters(c *gin.Context)
	CreateCluster(c *gin.Context)
	UpdateCluster(c *gin.Context)
	DeleteCluster(c *gin.Context)
	FindCluster(c *gin.Context)
}

type clusterControllerDefaultImp struct {
	ClusterRepository repository.ClusterRepositoryInterface
	ClusterValidator  validators.ClusterValidator
}

func NewClusterController(ClusterRepository repository.ClusterRepositoryInterface, ClusterValidator validators.ClusterValidator) *clusterControllerDefaultImp {
	return &clusterControllerDefaultImp{ClusterRepository: ClusterRepository, ClusterValidator: ClusterValidator}
}

// @BasePath /

// CreateCluster
// @Summary Register a new RabbitMQ Cluster
// @Schemes
// @Description Create a new <b>RabbitMQ</b> cluster. The credential provider must be valid and the cluster operational
// @Tags Cluster
// @Accept json
// @Produce json
// @Param request body contracts.CreateClusterRequest true "Request"
// @Success 201 {string} Helloworld
// @Router /cluster [post]
func (ctrl *clusterControllerDefaultImp) CreateCluster(c *gin.Context) {

	var request contracts.CreateClusterRequest
	if err := c.BindJSON(&request); err != nil {
		log.WithError(err).Error("Error parsing request")
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := ctrl.ClusterValidator.ValidateCreateRequest(request, c); err != nil {
		log.WithError(err).Error("Error validating request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entityToCreate := &entities.ClusterEntity{Name: request.Name, Host: request.Host, User: request.User, Password: request.Password, Port: request.Port,
		Description: request.Description}

	resp, err := ctrl.ClusterRepository.CreateCluster(entityToCreate)

	if err != nil {
		if _, ok := err.(*mysql.MySQLError); ok {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "Host already registred", "field": "host"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// DeleteCluster
// @Summary Soft delete a cluster
// @Schemes
// @Description Soft delete a cluster will not completly erase from database, but will not show up anymore in the
// system. All queues,bindings,shovels and related artifacts will be soft delete to
// @Tags Cluster
// @Accept json
// @Produce json
// @Success 204 {object} bool
// @Param clusterId path int true "Id of a cluster to be soft deleted"
// @Router /cluster/{clusterId} [delete]
func (ctrl *clusterControllerDefaultImp) DeleteCluster(c *gin.Context) {
	clusterIdParam := c.Param("clusterId")
	clusterId, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse clusterId Param")
		c.JSON(http.StatusBadRequest, "Error parsing clusterId from url route")
		return
	}

	log.WithField("clusterId", clusterId).Info("Cluster will be deleted")
	err = ctrl.ClusterRepository.DeleteCluster(uint(clusterId), c)
	if err != nil {
		errorMsg := "Fail to delete clusterId"
		log.WithError(err).WithField("clusterId", clusterId).Error(errorMsg)
		c.JSON(http.StatusInternalServerError, errorMsg)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": fmt.Sprintf("cluster with id %d successufly deleted", clusterId)})
}

// FindCluster
// @Summary Verify if exists a rabbitmqcluster
// @Schemes
// @Description Check if exists an rabbitmq cluster with host es
// @Tags Cluster
// @Accept json
// @Produce json
// @Success 200
// @Param params query FindClusterByHost true "Number of items in one page"
// @Router /cluster/exists [get]
func (ctrl *clusterControllerDefaultImp) FindCluster(c *gin.Context) {

	var param FindClusterByHost
	err := c.BindQuery(&param)
	if err != nil {
		log.WithError(err).Error("Error trying to parse query params")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exists := ctrl.ClusterRepository.CheckIfHostIsAlreadyRegisted(param.Host, param.Port, c)

	c.JSON(http.StatusOK, exists)

}

func (ctrl *clusterControllerDefaultImp) UpdateCluster(c *gin.Context) {

}

// GetCluster
// @Summary Retrieve a single rabbitmq cluster
// @Schemes
// @Description Retrieve a single rabbitmq cluster
// @Tags Cluster
// @Accept json
// @Produce json
// @Success 200 {object} entities.ClusterEntity
// @NotFound 404 {object} bool
// @Param clusterId path int true "Id of a cluster to be retrived"
// @Router /cluster/{clusterId} [get]
func (ctrl *clusterControllerDefaultImp) GetCluster(c *gin.Context) {
	clusterIdParam := c.Param("clusterId")
	clusterId, err := strconv.ParseInt(clusterIdParam, 10, 32)
	if err != nil {
		log.WithError(err).WithField("clusterId", clusterIdParam).Error("Fail to parse clusterId Param")
		c.JSON(http.StatusBadRequest, "Error parsing clusterId from url route")
		return
	}

	log.WithField("clusterId", clusterId).Info("Cluster will be deleted")
	cluster, err := ctrl.ClusterRepository.GetCluster(uint(clusterId), c)
	if err != nil {
		errorMsg := "Fail to retrive Cluster with clusterId"
		log.WithError(err).WithField("clusterId", clusterId).Error(errorMsg)
		c.JSON(http.StatusInternalServerError, errorMsg)
		return
	}
	c.JSON(http.StatusOK, cluster)
	return
}
