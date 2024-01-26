package cluster

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/core/core"
	"github.com/productivityeng/orabbit/core/validators"
	"github.com/productivityeng/orabbit/db"
	log "github.com/sirupsen/logrus"
)

type ClusterController interface {
	GetCluster(c *gin.Context)
	ListClusters(c *gin.Context)
	CreateCluster(c *gin.Context)
	UpdateCluster(c *gin.Context)
	DeleteCluster(c *gin.Context)
	FindCluster(c *gin.Context)
}

type clusterControllerDefaultImp struct {
	DependencyLocator *core.DependencyLocator
	ClusterValidator  validators.ClusterValidator
}

func NewClusterController(DependencyLocator *core.DependencyLocator, ClusterValidator validators.ClusterValidator) *clusterControllerDefaultImp {
	return &clusterControllerDefaultImp{DependencyLocator,  ClusterValidator}
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

	request,err := ctrl.parseCreateClusterParams(c)
	log.WithField("request", request).Info("Received request")
	if err != nil { return }

	// log.WithField("request", request).Info("Validating request")
	// if err := ctrl.ClusterValidator.ValidateCreateRequest(*request, c); err != nil {
	// 	log.WithError(err).Error("Error validating request")
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	log.WithField("request", request).Info("Creating cluster")

	result,err := ctrl.DependencyLocator.PrismaClient.Cluster.CreateOne(
		db.Cluster.Name.Set(request.Name),
		db.Cluster.Description.Set(request.Description),
		db.Cluster.Host.Set(request.Host),
		db.Cluster.Port.Set(request.Port),
		db.Cluster.User.Set(request.User),
		db.Cluster.Password.Set(request.Password),
	).Exec(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cluster already exists"})
		return
	}
	log.WithField("result", result).Info("Created cluster")

	c.JSON(http.StatusCreated, result)
}

func (ctrl *clusterControllerDefaultImp) parseCreateClusterParams(c *gin.Context) (params *contracts.CreateClusterRequest, err error) {
	var param contracts.CreateClusterRequest
	err = c.BindJSON(&param)
	if err != nil {
		log.WithError(err).Error("Error trying to parse query params")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}
	return &param, nil
}
