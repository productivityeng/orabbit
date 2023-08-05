package broker

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/broker/entities"
	"github.com/productivityeng/orabbit/broker/repository"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/core/validators"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type BrokerControllerTestSuit struct {
	suite.Suite
	brokerRepository      repository.BrokerRepositoryInterface
	brokerValidator       validators.BrokerValidator
	SUT                   *brokerControllerDefaultImp
	EndpointPath          string
	ParameterEndpointPath string
	CreateBrokerRequest   contracts.CreateBrokerRequest
	TestBrokers           []*entities.BrokerEntity
}

func (bct *BrokerControllerTestSuit) SetupSuite() {
	brokerValidatorObj := new(validators.BrokerValidatorMockedObject)
	brokerRepositoryObject := new(repository.BrokerRepositoryMockedObject)

	bct.brokerRepository = brokerRepositoryObject
	bct.brokerValidator = brokerValidatorObj
	bct.SUT = NewBrokerController(bct.brokerRepository, bct.brokerValidator)
	bct.EndpointPath = "/broker"
	bct.ParameterEndpointPath = fmt.Sprintf("%s/:brokerId", bct.EndpointPath)
	bct.CreateBrokerRequest = contracts.CreateBrokerRequest{
		Name:        "test",
		Description: "test",
		Host:        "host", //Missing Host parameter
		Port:        1111,
		User:        "testuser",
		Password:    "testpassword",
	}
	bct.TestBrokers = []*entities.BrokerEntity{{
		Id:          1,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   gorm.DeletedAt{},
		Name:        bct.CreateBrokerRequest.Name,
		Description: bct.CreateBrokerRequest.Description,
		Host:        bct.CreateBrokerRequest.Host,
		Port:        bct.CreateBrokerRequest.Port,
		User:        bct.CreateBrokerRequest.User,
		Password:    bct.CreateBrokerRequest.Password,
	}, {
		Id:          2,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   gorm.DeletedAt{},
		Name:        bct.CreateBrokerRequest.Name,
		Description: bct.CreateBrokerRequest.Description,
		Host:        bct.CreateBrokerRequest.Host,
		Port:        bct.CreateBrokerRequest.Port,
		User:        bct.CreateBrokerRequest.User,
		Password:    bct.CreateBrokerRequest.Password,
	}}

}

func (s *BrokerControllerTestSuit) TearDownTest() {
	s.SetupSuite()
}

func (bct *BrokerControllerTestSuit) TestBrokerControllerDefaultImpListBrokersShouldReturnBadRequestWhenAtLeastOneQueryParameterIsMissing() {

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET(bct.EndpointPath, bct.SUT.ListBrokers)

	req, _ := http.NewRequest("GET", bct.EndpointPath, nil)
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusBadRequest, res.Code)

}

func (bct *BrokerControllerTestSuit) TestBrokerControllerListBrokerShouldListBrokerWithSuccess() {
	pageSizeParam := 1
	pageNumberParam := 2
	listBrokerRepoResult := contracts.PaginatedResult[entities.BrokerEntity]{
		PageSize:   pageSizeParam,
		PageNumber: pageNumberParam,
		Result:     bct.TestBrokers,
	}
	bct.SUT.BrokerRepository.(*repository.BrokerRepositoryMockedObject).On("ListBroker", pageSizeParam, pageNumberParam).Return(&listBrokerRepoResult, nil)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET(bct.EndpointPath, bct.SUT.ListBrokers)

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s?PageSize=%d&PageNumber=%d", bct.EndpointPath, pageSizeParam, pageNumberParam), nil)
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusOK, res.Code)
	assert.Equal(bct.T(), listBrokerRepoResult.Result, bct.TestBrokers)
	assert.Equal(bct.T(), listBrokerRepoResult.PageSize, pageSizeParam)
	assert.Equal(bct.T(), listBrokerRepoResult.PageNumber, pageNumberParam)
}
func (bct *BrokerControllerTestSuit) TestBrokerControllerListBrokerShouldReturnInternalErrorWhenFailtToListBrokers() {
	pageSizeParam := 1
	pageNumberParam := 2
	bct.SUT.BrokerRepository.(*repository.BrokerRepositoryMockedObject).On("ListBroker", pageSizeParam, pageNumberParam).Return(nil, errors.New("generic error"))
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET(bct.EndpointPath, bct.SUT.ListBrokers)

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s?PageSize=%d&PageNumber=%d", bct.EndpointPath, pageSizeParam, pageNumberParam), nil)
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusInternalServerError, res.Code)

}

func (bct *BrokerControllerTestSuit) TestBrokerControllerDeleteBrokerShouldReturnBadRequestWhenBrokerIdIsMalformated() {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE(fmt.Sprintf(bct.ParameterEndpointPath), bct.SUT.DeleteBroker)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/banana", bct.EndpointPath), nil)
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusBadRequest, res.Code)

}

func (bct *BrokerControllerTestSuit) TestBrokerControllerDeleteBrokerShouldReturnInternalServerErorWhenFailToDeleteBroker() {
	brokerIdTobeDeleted := 2

	bct.SUT.BrokerRepository.(*repository.BrokerRepositoryMockedObject).On("DeleteBroker", int32(brokerIdTobeDeleted), mock.Anything).Return(errors.New("generic error to delete broker"))
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE(bct.ParameterEndpointPath, bct.SUT.DeleteBroker)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/%d", bct.EndpointPath, brokerIdTobeDeleted), nil)
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusInternalServerError, res.Code)

}

func (bct *BrokerControllerTestSuit) TestBrokerControllerDeleteBrokerShouldBeSuccess() {
	brokerIdTobeDeleted := 1

	bct.SUT.BrokerRepository.(*repository.BrokerRepositoryMockedObject).On("DeleteBroker", int32(brokerIdTobeDeleted), mock.Anything).Return(nil)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE(bct.ParameterEndpointPath, bct.SUT.DeleteBroker)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/%d", bct.EndpointPath, brokerIdTobeDeleted), nil)
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusOK, res.Code)

}

func (bct *BrokerControllerTestSuit) TestBrokerControllerCreateBrokerShouldReturnbadRequestWhenRequestIsMalformed() {
	brokerRequest := contracts.CreateBrokerRequest{
		Name:        "test",
		Description: "test",
		Host:        "host", //Missing Host parameter
		Port:        1111,
		User:        "testuser",
		Password:    "testpassword",
	}
	brokerRequestPayload, _ := json.Marshal(brokerRequest)

	//Dependencies
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST(bct.EndpointPath, bct.SUT.CreateBroker)

	//Create malformed payload
	brokerRequestPayload[4] = byte(4)
	req, _ := http.NewRequest("POST", bct.EndpointPath, bytes.NewBuffer(brokerRequestPayload))
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusBadRequest, res.Code)
}

func (bct *BrokerControllerTestSuit) TestBrokerControllerCreateBrokerShouldReturnInternalServerErrorWhenFailtToCreateBroker() {

	brokerRequest := contracts.CreateBrokerRequest{
		Name:        "test",
		Description: "test",
		Host:        "host", //Missing Host parameter
		Port:        1111,
		User:        "testuser",
		Password:    "testpassword",
	}
	brokerRequestPayload, _ := json.Marshal(brokerRequest)

	//Dependencies
	bct.SUT.BrokerValidator.(*validators.BrokerValidatorMockedObject).On("ValidateCreateRequest").Return(nil)
	bct.SUT.BrokerRepository.(*repository.BrokerRepositoryMockedObject).On("CreateBroker", mock.Anything).Return(nil, errors.New("fail to create broker"))

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST(bct.EndpointPath, bct.SUT.CreateBroker)

	req, _ := http.NewRequest("POST", bct.EndpointPath, bytes.NewBuffer(brokerRequestPayload))
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusInternalServerError, res.Code)
}

func (bct *BrokerControllerTestSuit) TestBrokerControllerCreateBrokerShouldReturnBadRequestWhenValidateBrokerFail() {
	brokerRequest := contracts.CreateBrokerRequest{
		Name:        "test",
		Description: "test",
		Host:        "host", //Missing Host parameter
		Port:        1111,
		User:        "testuser",
		Password:    "testpassword",
	}
	brokerRequestPayload, _ := json.Marshal(brokerRequest)

	//Dependencies
	bct.SUT.BrokerValidator.(*validators.BrokerValidatorMockedObject).On("ValidateCreateRequest", mock.Anything, mock.Anything).Return(errors.New("invalid broker"))

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST(bct.EndpointPath, bct.SUT.CreateBroker)

	req, _ := http.NewRequest("POST", bct.EndpointPath, bytes.NewBuffer(brokerRequestPayload))
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusBadRequest, res.Code)
}

func (bct *BrokerControllerTestSuit) TestBrokerControllerDefaultImpCreateBrokerInternalServerErrorWhenFailToSave() {

	brokerRequest := contracts.CreateBrokerRequest{
		Name:        "test",
		Description: "test",
		Host:        "host", //Missing Host parameter
		Port:        1111,
		User:        "testuser",
		Password:    "testpassword",
	}
	brokerRequestPayload, _ := json.Marshal(brokerRequest)
	expectedBrokerCreated := bct.TestBrokers[0]

	//Dependencies
	bct.SUT.BrokerValidator.(*validators.BrokerValidatorMockedObject).On("ValidateCreateRequest", mock.Anything, mock.Anything).Return(nil)
	bct.SUT.BrokerRepository.(*repository.BrokerRepositoryMockedObject).On("CreateBroker", mock.Anything).Return(expectedBrokerCreated, nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST(bct.EndpointPath, bct.SUT.CreateBroker)

	req, _ := http.NewRequest("POST", bct.EndpointPath, bytes.NewBuffer(brokerRequestPayload))
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusCreated, res.Code)
}

func (bct *BrokerControllerTestSuit) TestBrokerControllerGetBrokerShouldReturnBadRequestWhenBrokerIdIsMalformated() {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET(fmt.Sprintf(bct.ParameterEndpointPath), bct.SUT.GetBroker)

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/banana", bct.EndpointPath), nil)
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusBadRequest, res.Code)

}

func TestBrokerControllerSuit(t *testing.T) {
	suite.Run(t, new(BrokerControllerTestSuit))
}
