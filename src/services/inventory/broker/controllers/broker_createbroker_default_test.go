package broker

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/productivityeng/orabbit/broker/entities"
	"github.com/productivityeng/orabbit/broker/repository"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/core/validators"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func executeRequest(path string, handler gin.HandlerFunc, request []byte) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST(path, handler)

	req, _ := http.NewRequest("POST", path, bytes.NewBuffer(request))
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)
	return res
}
func TestValidateShouldReturnBadRequestWhenContractValidationFail(t *testing.T) {

	//Dependencies
	brokerValidatorObj := new(validators.BrokerValidatorMockedObject)
	brokerValidatorObj.On("ValidateCreateRequest").Return(errors.New("invalid contract"))

	brokerRepositoryObject := new(repository.BrokerRepositoryMockedObject)

	//SUT
	SUT := &brokerControllerDefaultImp{BrokerValidator: brokerValidatorObj, BrokerRepository: brokerRepositoryObject}

	brokerRequest := contracts.CreateBrokerRequest{
		Name:        "test",
		Description: "test",
		Host:        "test",
		Port:        777,
		User:        "testuser",
		Password:    "testpassword",
	}
	brokerRequestPayload, _ := json.Marshal(brokerRequest)

	res := executeRequest("/broker", SUT.CreateBroker, brokerRequestPayload)
	assert.Equal(t, http.StatusBadRequest, res.Code)

}

func TestBrokerControllerDefaultImpCreateBrokerErrorWithMissingRequiredParameter(t *testing.T) {
	//Dependencies
	brokerValidatorObj := new(validators.BrokerValidatorMockedObject)
	brokerRepositoryObject := new(repository.BrokerRepositoryMockedObject)

	//SUT
	SUT := &brokerControllerDefaultImp{BrokerValidator: brokerValidatorObj, BrokerRepository: brokerRepositoryObject}
	brokerRequest := contracts.CreateBrokerRequest{
		Name:        "test",
		Description: "test",
		Host:        "", //Missing Host parameter
		Port:        777,
		User:        "testuser",
		Password:    "testpassword",
	}
	brokerRequestPayload, _ := json.Marshal(brokerRequest)

	res := executeRequest("/broker", SUT.CreateBroker, brokerRequestPayload)
	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func TestBrokerControllerDefaultImpCreateBrokerInternalServerErrorWhenFailToSave(t *testing.T) {

	brokerRequest := contracts.CreateBrokerRequest{
		Name:        "test",
		Description: "test",
		Host:        "host", //Missing Host parameter
		Port:        1111,
		User:        "testuser",
		Password:    "testpassword",
	}
	brokerRequestPayload, _ := json.Marshal(brokerRequest)
	expectedBrokerCreated := &entities.BrokerEntity{
		Id:          1,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   gorm.DeletedAt{},
		Name:        brokerRequest.Name,
		Description: brokerRequest.Description,
		Host:        brokerRequest.Host,
		Port:        brokerRequest.Port,
		User:        brokerRequest.User,
		Password:    brokerRequest.Password,
	}

	//Dependencies
	brokerValidatorObj := new(validators.BrokerValidatorMockedObject)
	brokerValidatorObj.On("ValidateCreateRequest").Return(nil)

	brokerRepositoryObject := new(repository.BrokerRepositoryMockedObject)
	brokerRepositoryObject.On("CreateBroker", mock.Anything).Return(expectedBrokerCreated, nil)
	//SUT
	SUT := &brokerControllerDefaultImp{BrokerValidator: brokerValidatorObj, BrokerRepository: brokerRepositoryObject}

	res := executeRequest("/broker", SUT.CreateBroker, brokerRequestPayload)
	assert.Equal(t, http.StatusCreated, res.Code)
}
