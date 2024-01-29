package cluster

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/cluster/models"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/core/core"
	"github.com/productivityeng/orabbit/core/validators"
	"github.com/productivityeng/orabbit/db"
	"github.com/productivityeng/orabbit/rabbitmq/virtualhost"
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


// importDefaultVirtualHost is a function that imports the default virtual host from a cluster.
// It takes a cluster model and a gin context as parameters.
// It returns a virtual host model and an error.
func (ctrl *clusterControllerDefaultImp) importDefaultVirtualHost(cluster *db.ClusterModel,c *gin.Context) (virtualHost *db.VirtualHostModel,err error) {
	access := models.GetRabbitMqAccess(cluster)
	virtualHostFromCluster,err := ctrl.DependencyLocator.VirtualHostManagement.GetVirtualHost(virtualhost.GetVirtualHostRequest{
		RabbitAccess: access,
		Name: "/",
	})
	if err != nil { 
		return nil,err
	}

	_tags,_ :=json.Marshal(virtualHostFromCluster.Tags)

	virtualHostCreated,err := ctrl.DependencyLocator.PrismaClient.VirtualHost.CreateOne(
		db.VirtualHost.Name.Set(virtualHostFromCluster.Name),
		db.VirtualHost.Description.Set(virtualHostFromCluster.Description),
		db.VirtualHost.DefaultQueueType.Set(db.ParseQueueType(virtualHostFromCluster.DefaultQueueType)),
		db.VirtualHost.Tags.Set(_tags),
		db.VirtualHost.Cluster.Link(db.Cluster.ID.Equals(cluster.ID)),
		).Exec(c)
	if err != nil { 
		return nil,err
	}
	return virtualHostCreated,nil
}


// putDefaultLockerOnVirtualHost puts a default locker on a virtual host.
// It creates a new locker for the specified virtual host ID and sets the necessary fields.
// Parameters:
// - virtualHostId: The ID of the virtual host.
// - c: The gin context.
// Returns an error if there was a problem creating the locker.
func (ctrl *clusterControllerDefaultImp) putDefaultLockerOnVirtualHost(virtualHostId int,c *gin.Context) error {
	_,err := ctrl.DependencyLocator.PrismaClient.LockerVirtualHost.CreateOne(
		db.LockerVirtualHost.VirtualHost.Link(db.VirtualHost.ID.Equals(virtualHostId)),
		db.LockerVirtualHost.Enabled.Set(true),
		db.LockerVirtualHost.Reason.Set("Default virtual host"),
		db.LockerVirtualHost.UserResponsibleEmail.Set("system@system"),
	).Exec(c)
	return err
}



func (ctrl *clusterControllerDefaultImp) importDefaultExchanges(cluster *db.ClusterModel,virtualHost *db.VirtualHostModel,c *gin.Context)(exchanges []db.ExchangeModel,err error) {
	defaultExchangesList := []string{"amq.direct","amq.fanout","amq.headers","amq.match","amq.rabbitmq.log","amq.topic"}	
	access := models.GetRabbitMqAccess(cluster)


	exchangesCreated := make([]db.ExchangeModel,0)
	for _,exchangeName := range defaultExchangesList { 
		exchangeFromCluster,err := ctrl.DependencyLocator.ExchangeManagement.GetExchangeByName(contracts.GetExchangeRequest{ 
			RabbitAccess: access,
			Name: exchangeName,
		},c)

		if err != nil { 
			continue
		}

		_arguments,_ := json.Marshal(exchangeFromCluster.Arguments)

		exchangeFromDatabase,err := ctrl.DependencyLocator.PrismaClient.Exchange.CreateOne(
			db.Exchange.Cluster.Link(db.Cluster.ID.Equals(cluster.ID)),
			db.Exchange.VirtualHost.Link(db.VirtualHost.ID.Equals(virtualHost.ID)),
			db.Exchange.Name.Set(exchangeFromCluster.Name),
			db.Exchange.Internal.Set(exchangeFromCluster.Internal),
			db.Exchange.Durable.Set(exchangeFromCluster.Durable),
			db.Exchange.Arguments.Set(_arguments),
			db.Exchange.Type.Set(exchangeFromCluster.Type),
			).Exec(c)
		if err != nil {
			log.WithError(err).Error("Error creating exchange")
		}else {
			exchangesCreated = append(exchangesCreated,*exchangeFromDatabase)
		}
	}
	return exchangesCreated,nil
}

func (ctrl *clusterControllerDefaultImp) putDefaultLockerOnExchange(exchangeId int,c *gin.Context) error {
	_,err := ctrl.DependencyLocator.PrismaClient.LockerExchange.CreateOne(
		db.LockerExchange.Exchange.Link(db.Exchange.ID.Equals(exchangeId)),
		db.LockerExchange.Enabled.Set(true),
		db.LockerExchange.Reason.Set("Default exchange"),
		db.LockerExchange.UserResponsibleEmail.Set("system@system"),
	).Exec(c)
	return err
}

/// importDefaultUser imports the default user from a cluster.
func (ctrl *clusterControllerDefaultImp) importDefaultUser(cluster *db.ClusterModel,username string,c *gin.Context) (user *db.UserModel,err error) {
	access := models.GetRabbitMqAccess(cluster)
	userFromCluster,err :=ctrl.DependencyLocator.UserManagement.GetUserByName(contracts.GetUserByNameRequest{ 
		RabbitAccess: access,
		Username: username,
	},c)
	if err != nil { 
		return nil,err
	}

	userCreated,err := ctrl.DependencyLocator.PrismaClient.User.CreateOne(
		db.User.Username.Set(userFromCluster.Name),
		db.User.PasswordHash.Set(userFromCluster.PasswordHash),
		db.User.Cluster.Link(db.Cluster.ID.Equals(cluster.ID)),
	).Exec(c)
	if err != nil { 
		return nil,err
	}
	return userCreated,nil

}
func (ctrl *clusterControllerDefaultImp) putDefaultLockerOnUser(userId int,c *gin.Context) error {
	_,err := ctrl.DependencyLocator.PrismaClient.LockerUser.CreateOne(
		db.LockerUser.User.Link(db.User.ID.Equals(userId)),
		db.LockerUser.Enabled.Set(true),
		db.LockerUser.Reason.Set("Default user"),
		db.LockerUser.UserResponsibleEmail.Set("system@system"),
	).Exec(c)
	return err
}
