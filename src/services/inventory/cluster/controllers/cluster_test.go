package cluster

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/productivityeng/orabbit/cluster/entities"
	"github.com/productivityeng/orabbit/cluster/repository"
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
	brokerRepository      repository.ClusterRepositoryInterface
	brokerValidator       validators.ClusterValidator
	SUT                   *clusterControllerDefaultImp
	EndpointPath          string
	ParameterEndpointPath string
	CreateClusterRequest  contracts.CreateClusterRequest
	TestBrokers           []*entities.ClusterEntity
}

func (bct *BrokerControllerTestSuit) SetupSuite() {
	brokerValidatorObj := new(validators.ClusterValidatorMockedObject)
	brokerRepositoryObject := new(repository.ClusterRepositoryMockedObject)

	bct.brokerRepository = brokerRepositoryObject
	bct.brokerValidator = brokerValidatorObj
	bct.SUT = NewClusterController(bct.brokerRepository, bct.brokerValidator)
	bct.EndpointPath = "/broker"
	bct.ParameterEndpointPath = fmt.Sprintf("%s/:brokerId", bct.EndpointPath)
	bct.CreateClusterRequest = contracts.CreateClusterRequest{
		Name:        "test",
		Description: "test",
		Host:        "host", //Missing Host parameter
		Port:        1111,
		User:        "testuser",
		Password:    "testpassword",
	}
	bct.TestBrokers = []*entities.ClusterEntity{{
		Id:          1,
		Name:        bct.CreateClusterRequest.Name,
		Description: bct.CreateClusterRequest.Description,
		Host:        bct.CreateClusterRequest.Host,
		Port:        bct.CreateClusterRequest.Port,
		User:        bct.CreateClusterRequest.User,
		Password:    bct.CreateClusterRequest.Password, Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		},
	}, {
		Id:          2,
		Name:        bct.CreateClusterRequest.Name,
		Description: bct.CreateClusterRequest.Description,
		Host:        bct.CreateClusterRequest.Host,
		Port:        bct.CreateClusterRequest.Port,
		User:        bct.CreateClusterRequest.User,
		Password:    bct.CreateClusterRequest.Password, Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		},
	}}

}

func (s *BrokerControllerTestSuit) TearDownTest() {
	s.SetupSuite()
}

func (bct *BrokerControllerTestSuit) TestBrokerControllerDefaultImpListClustersShouldReturnBadRequestWhenAtLeastOneQueryParameterIsMissing() {

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET(bct.EndpointPath, bct.SUT.ListClusters)

	req, _ := http.NewRequest("GET", bct.EndpointPath, nil)
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusBadRequest, res.Code)

}

func (bct *BrokerControllerTestSuit) TestBrokerControllerListBrokerShouldListBrokerWithSuccess() {
	pageSizeParam := 1
	pageNumberParam := 2
	listBrokerRepoResult := contracts.PaginatedResult[entities.ClusterEntity]{
		PageSize:   pageSizeParam,
		PageNumber: pageNumberParam,
		Result:     bct.TestBrokers,
	}
	bct.SUT.ClusterRepository.(*repository.ClusterRepositoryMockedObject).On("ListCluster", pageSizeParam, pageNumberParam).Return(&listBrokerRepoResult, nil)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET(bct.EndpointPath, bct.SUT.ListClusters)

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s?PageSize=%d&PageNumber=%d", bct.EndpointPath, pageSizeParam, pageNumberParam), nil)
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusOK, res.Code)
	assert.Equal(bct.T(), listBrokerRepoResult.Result, bct.TestBrokers)
	assert.Equal(bct.T(), listBrokerRepoResult.PageSize, pageSizeParam)
	assert.Equal(bct.T(), listBrokerRepoResult.PageNumber, pageNumberParam)
}
func (bct *BrokerControllerTestSuit) TestBrokerControllerListBrokerShouldReturnInternalErrorWhenFailtToListClusters() {
	pageSizeParam := 1
	pageNumberParam := 2
	bct.SUT.ClusterRepository.(*repository.ClusterRepositoryMockedObject).On("ListCluster", pageSizeParam, pageNumberParam).Return(nil, errors.New("generic error"))
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET(bct.EndpointPath, bct.SUT.ListClusters)

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s?PageSize=%d&PageNumber=%d", bct.EndpointPath, pageSizeParam, pageNumberParam), nil)
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusInternalServerError, res.Code)

}

func (bct *BrokerControllerTestSuit) TestBrokerControllerDeleteBrokerShouldReturnBadRequestWhenBrokerIdIsMalformated() {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE(fmt.Sprintf(bct.ParameterEndpointPath), bct.SUT.DeleteCluster)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/banana", bct.EndpointPath), nil)
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusBadRequest, res.Code)

}

func (bct *BrokerControllerTestSuit) TestBrokerControllerDeleteBrokerShouldReturnInternalServerErorWhenFailToDeleteBroker() {
	brokerIdTobeDeleted := 2

	bct.SUT.ClusterRepository.(*repository.ClusterRepositoryMockedObject).On("DeleteCluster", int32(brokerIdTobeDeleted), mock.Anything).Return(errors.New("generic error to delete broker"))
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE(bct.ParameterEndpointPath, bct.SUT.DeleteCluster)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/%d", bct.EndpointPath, brokerIdTobeDeleted), nil)
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusInternalServerError, res.Code)

}

func (bct *BrokerControllerTestSuit) TestBrokerControllerDeleteBrokerShouldBeSuccess() {
	brokerIdTobeDeleted := 1

	bct.SUT.ClusterRepository.(*repository.ClusterRepositoryMockedObject).On("DeleteCluster", int32(brokerIdTobeDeleted), mock.Anything).Return(nil)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE(bct.ParameterEndpointPath, bct.SUT.DeleteCluster)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/%d", bct.EndpointPath, brokerIdTobeDeleted), nil)
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusOK, res.Code)

}

func (bct *BrokerControllerTestSuit) TestBrokerControllerCreateClusterShouldReturnbadRequestWhenRequestIsMalformed() {
	brokerRequest := contracts.CreateClusterRequest{
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
	router.POST(bct.EndpointPath, bct.SUT.CreateCluster)

	//Create malformed payload
	brokerRequestPayload[4] = byte(4)
	req, _ := http.NewRequest("POST", bct.EndpointPath, bytes.NewBuffer(brokerRequestPayload))
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusBadRequest, res.Code)
}

func (bct *BrokerControllerTestSuit) TestBrokerControllerCreateClusterShouldReturnInternalServerErrorWhenFailtToCreateCluster() {

	brokerRequest := contracts.CreateClusterRequest{
		Name:        "test",
		Description: "test",
		Host:        "host", //Missing Host parameter
		Port:        1111,
		User:        "testuser",
		Password:    "testpassword",
	}
	brokerRequestPayload, _ := json.Marshal(brokerRequest)

	//Dependencies
	bct.SUT.ClusterValidator.(*validators.ClusterValidatorMockedObject).On("ValidateCreateRequest").Return(nil)
	bct.SUT.ClusterRepository.(*repository.ClusterRepositoryMockedObject).On("CreateCluster", mock.Anything).Return(nil, errors.New("fail to create broker"))

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST(bct.EndpointPath, bct.SUT.CreateCluster)

	req, _ := http.NewRequest("POST", bct.EndpointPath, bytes.NewBuffer(brokerRequestPayload))
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusInternalServerError, res.Code)
}

func (bct *BrokerControllerTestSuit) TestBrokerControllerCreateClusterShouldReturnBadRequestWhenValidateBrokerFail() {
	brokerRequest := contracts.CreateClusterRequest{
		Name:        "test",
		Description: "test",
		Host:        "host", //Missing Host parameter
		Port:        1111,
		User:        "testuser",
		Password:    "testpassword",
	}
	brokerRequestPayload, _ := json.Marshal(brokerRequest)

	//Dependencies
	bct.SUT.ClusterValidator.(*validators.ClusterValidatorMockedObject).On("ValidateCreateRequest", mock.Anything, mock.Anything).Return(errors.New("invalid broker"))

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST(bct.EndpointPath, bct.SUT.CreateCluster)

	req, _ := http.NewRequest("POST", bct.EndpointPath, bytes.NewBuffer(brokerRequestPayload))
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusBadRequest, res.Code)
}

func (bct *BrokerControllerTestSuit) TestBrokerControllerDefaultImpCreateClusterInternalServerErrorWhenFailToSave() {

	brokerRequest := contracts.CreateClusterRequest{
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
	bct.SUT.ClusterValidator.(*validators.ClusterValidatorMockedObject).On("ValidateCreateRequest", mock.Anything, mock.Anything).Return(nil)
	bct.SUT.ClusterRepository.(*repository.ClusterRepositoryMockedObject).On("CreateCluster", mock.Anything).Return(expectedBrokerCreated, nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST(bct.EndpointPath, bct.SUT.CreateCluster)

	req, _ := http.NewRequest("POST", bct.EndpointPath, bytes.NewBuffer(brokerRequestPayload))
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusCreated, res.Code)
}

func (bct *BrokerControllerTestSuit) TestBrokerControllerGetBrokerShouldReturnBadRequestWhenBrokerIdIsMalformated() {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET(fmt.Sprintf(bct.ParameterEndpointPath), bct.SUT.GetCluster)

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/banana", bct.EndpointPath), nil)
	res := httptest.NewRecorder()
	//Execute request
	router.ServeHTTP(res, req)

	assert.Equal(bct.T(), http.StatusBadRequest, res.Code)

}

func TestBrokerControllerSuit(t *testing.T) {
	suite.Run(t, new(BrokerControllerTestSuit))
}
