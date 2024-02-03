package functions

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/db"
)

//VerifyIfExchangesIsAlreadyInDatabase
// verify if exchange is already in database
func VerifyIfExchangesIsAlreadyInDatabase(prismaClient *db.PrismaClient, exchangeName string, c *gin.Context) (error) {
	_,err :=prismaClient.Exchange.FindUnique(db.Exchange.Name.Equals(exchangeName)).Exec(c)
	if errors.Is(err, db.ErrNotFound) { 
		return nil
	}
	c.JSON(http.StatusBadRequest, "Exchange already in database")
	return errors.New("exchange already in database")
}
