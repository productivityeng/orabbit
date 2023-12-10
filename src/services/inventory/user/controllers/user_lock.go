package controllers

import (
	"github.com/gin-gonic/gin"
	dto "github.com/productivityeng/orabbit/user/dto"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// LockUser godoc
// @Summary Lock a user
// @Schemes
// @Description Lock a user with a reason
// @Tags User
// @Accept json
// @Produce json
// @Param lockUserRequest body dto.LockUserDto true "Request"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /{clusterId}/user/{userId}/lock [post]
func (ctrl *UserControllerImpl) LockUser(c *gin.Context) {
	var lockUserRequest dto.LockUserDto

	err := c.BindJSON(&lockUserRequest)
	if err != nil {
		log.WithContext(c).WithError(err).Error("Fail to parse lockuser request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	return

}
