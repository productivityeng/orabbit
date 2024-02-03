package functions

import (
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core"
	"github.com/productivityeng/orabbit/db"
)


func GetVirtualHostById(DependencyLocator *core.DependencyLocator, c *gin.Context,virtualHostId int) (*db.VirtualHostModel, error) { 
	virtualHost,err := DependencyLocator.PrismaClient.VirtualHost.FindUnique(db.VirtualHost.ID.Equals(virtualHostId)).Exec(c)
	if err != nil { 
		return nil,err
	}
	return virtualHost,nil
}