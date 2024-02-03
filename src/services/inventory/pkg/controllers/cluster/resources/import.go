package resources

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core"
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
}

func NewClusterController(DependencyLocator *core.DependencyLocator) *clusterControllerDefaultImp {
	return &clusterControllerDefaultImp{DependencyLocator}
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

	log.WithField("request", request).Info("Creating cluster")
	createClusterResult,err := ctrl.DependencyLocator.PrismaClient.Cluster.CreateOne(
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

	virtualHost,err :=ctrl.importDefaultVirtualHost(createClusterResult,c)
	if err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error importing default virtual host"})
		return
	}

	err = ctrl.putDefaultLockerOnVirtualHost(virtualHost.ID,c)
	if err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error putting default locker on virtual host"})
		return
	}


	defaultUser,err := ctrl.importDefaultUser(createClusterResult,request.User,c)
	if err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error importing default user"})
		return
	}

	err = ctrl.putDefaultLockerOnUser(defaultUser.ID,c)
	if err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error putting default locker on user"})
		return
	}
	
	exchanges,err := ctrl.importDefaultExchanges(createClusterResult,virtualHost,c)
	if err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error importing default exchanges"})
		return
	} 

	for _,exchange := range exchanges { 
		err = ctrl.putDefaultLockerOnExchange(exchange.ID,c)
		if err != nil { 
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error putting default locker on exchange"})
			return
		}
	}

	log.WithField("result", createClusterResult).Info("Created cluster")

	c.JSON(http.StatusCreated, createClusterResult)
}
