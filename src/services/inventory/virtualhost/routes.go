package virtualhost

import (
	"github.com/gin-gonic/gin"
	repository2 "github.com/productivityeng/orabbit/cluster/repository"
	"github.com/productivityeng/orabbit/src/packages/rabbitmq/virtualhost"
	"github.com/productivityeng/orabbit/virtualhost/controllers"
	"github.com/productivityeng/orabbit/virtualhost/repository"
	"gorm.io/gorm"
)

var clusterRepository repository2.ClusterRepositoryInterface
var virtualHostController controllers.VirtualHostController
var virtualHostRepository repository.VirtualHostRepository

func Routes(routes *gin.Engine, db *gorm.DB) *gin.RouterGroup {
	clusterRepository = repository2.NewClusterMysqlRepositoryImpl(db)
	virtualHostRepository = repository.NewVirtualHostRepositoryMysql(db)

	virtualHostController = controllers.NewVirtualHostControllerImpl(virtualhost.NewirtualHostManagement(),
		clusterRepository, virtualHostRepository)

	userRouter := routes.Group("/:clusterId/virtualhost")
	userRouter.GET("/", virtualHostController.ListVirtualHost)
	userRouter.POST("/", virtualHostController.ImportOrCreateVirtualHost)

	return userRouter
}
