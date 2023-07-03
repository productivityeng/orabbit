package broker

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/productivityeng/orabbit/broker/repository"
	"github.com/productivityeng/orabbit/core/validators"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBrokerControllerDefaultImp_ListBrokers_ShouldReturn_BadRequest_WhenAtLeastOneQueryParameterIsMissing(t *testing.T) {
	brokerValidatorObj := new(validators.BrokerValidatorMockedObject)
	brokerRepositoryObject := new(repository.BrokerRepositoryMockedObject)

	SUT := &brokerControllerDefaultImp{BrokerValidator: brokerValidatorObj, BrokerRepository: brokerRepositoryObject}

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/broker", SUT.ListBrokers)

	req, _ := http.NewRequest("GET", "/broker", nil)
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(t, res.Code, http.StatusBadRequest)
}
