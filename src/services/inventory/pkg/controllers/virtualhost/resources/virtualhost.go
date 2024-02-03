package resources

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/core"
	"github.com/productivityeng/orabbit/pkg/controllers/virtualhost/dto"
	"github.com/sirupsen/logrus"
)

type VirtualHostController interface {
	ListVirtualHost(c *gin.Context)
	Import(c *gin.Context)
	Syncronize(c *gin.Context)
	DeleteVirtualHost(c *gin.Context)
}

type VirtualHostControllerImpl struct {
	DependencyLocator     *core.DependencyLocator	
}

func NewVirtualHostControllerImpl(
	DependencyLocator *core.DependencyLocator) VirtualHostControllerImpl {
	return VirtualHostControllerImpl{
		DependencyLocator: DependencyLocator,
	}
}



// parseImportVirtualHostBody parses the request body into a ImportVirtualHostRequest struct.
// It takes a gin.Context as a parameter and returns a ImportVirtualHostRequest and an error.
// If there is an error parsing the body, it logs the error and returns a bad request response.
func (ctrl *VirtualHostControllerImpl) parseImportVirtualHostBody(c *gin.Context) (request dto.ImportVirtualHostRequest,err error){
	err = c.ShouldBindJSON(&request)
	if err != nil {
		logrus.WithContext(c).WithError(err).Error("Fail to parse body")
		c.JSON(http.StatusBadRequest, "Error parsing body")
		return
	}
	return
}

