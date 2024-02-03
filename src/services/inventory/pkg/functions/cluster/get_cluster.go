package functions

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core"
	"github.com/productivityeng/orabbit/db"
)


func GetCluster(DependencyLocator *core.DependencyLocator, clusterId int ,c *gin.Context) (*db.ClusterModel, error) { 
	cluster,err := DependencyLocator.PrismaClient.Cluster.FindUnique(db.Cluster.ID.Equals(clusterId)).Exec(c)
	if err != nil { 
		return nil,err
	}
	return cluster,nil
}