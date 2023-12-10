package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/productivityeng/orabbit/cluster/entities"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

type BrokerRepositorySuite struct {
	suite.Suite
	conn   *sql.DB
	DB     *gorm.DB
	mock   sqlmock.Sqlmock
	SUT    *ClusterRepositoryMysqlImpl
	broker *entities.ClusterEntity
}

func (suite *BrokerRepositorySuite) TearDownTest() {
	assert.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (brs *BrokerRepositorySuite) SetupSuite() {
	log.SetLevel(log.InfoLevel)
	var err error
	brs.conn, brs.mock, err = sqlmock.New()
	assert.NoError(brs.T(), err)

	brs.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      brs.conn,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})

	assert.NoError(brs.T(), err)

	brs.SUT = NewClusterMysqlRepositoryImpl(brs.DB)
	assert.IsType(brs.T(), &ClusterRepositoryMysqlImpl{}, brs.SUT)

	brs.broker = &entities.ClusterEntity{
		Name:        "TestBroker",
		Description: "Broker of test",
		Host:        "rabbitmq.com",
		Port:        999,
		User:        "rabbit",
		Password:    "rabbit",
	}
}

func (brs *BrokerRepositorySuite) TestListBroker() {

	expectedResult := &entities.ClusterEntity{
		Name:        "Test Broker",
		Description: "Test Description",
		Host:        "localhost",
		Port:        1234,
		User:        "test_user",
		Password:    "test_password",
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{
		"id",
		"created_at",
		"updated_at",
		"name",
		"description",
		"host",
		"port",
		"user",
		"password",
	}).AddRow(
		expectedResult.ID,
		expectedResult.CreatedAt,
		expectedResult.UpdatedAt,
		expectedResult.Name,
		expectedResult.Description,
		expectedResult.Host,
		expectedResult.Port,
		expectedResult.User,
		expectedResult.Password,
	).AddRow(
		expectedResult.ID,
		expectedResult.CreatedAt,
		expectedResult.UpdatedAt,
		expectedResult.Name,
		expectedResult.Description,
		expectedResult.Host,
		expectedResult.Port,
		expectedResult.User,
		expectedResult.Password,
	)

	pageSize := 1
	pageNumber := 2

	brs.mock.ExpectQuery("").WillReturnRows(rows)
	result, err := brs.SUT.ListCluster(pageSize, pageNumber)
	assert.Nil(brs.T(), err)
	assert.Equal(brs.T(), expectedResult, result.Result[0])
	assert.Equal(brs.T(), pageSize, result.PageSize)
	assert.Equal(brs.T(), pageNumber, result.PageNumber)
}

// TestListBrokerErrorTryingToRetrieveResult check if can deal with error in list brokers
func (brs *BrokerRepositorySuite) TestListBrokerErrorTryingToRetrieveResult() {

	pageSize := 1
	pageNumber := 2

	brs.mock.ExpectQuery("").WillReturnError(errors.New("genericerro"))
	result, err := brs.SUT.ListCluster(pageSize, pageNumber)
	assert.NotNil(brs.T(), err)
	assert.Nil(brs.T(), result)
}

// TestBrokerDeleteShouldReturnErrorWhenBrokerNotExists check if can deal with error in delete a broker
func (brs *BrokerRepositorySuite) TestBrokerDeleteShouldReturnErrorWhenBrokerNotExists() {
	brokerid := uint(10)
	brs.mock.ExpectQuery("SELECT").WillReturnError(errors.New("genericerro"))

	err := brs.SUT.DeleteCluster(brokerid, context.TODO())
	assert.NotNil(brs.T(), err)
	assert.Equal(brs.T(), err.Error(), "broker id cound't not be found")

}

// TestBrokerDeleteShouldReturnErrorWhenFailToDeleteBroker check if can deal with error in delete a broker
func (brs *BrokerRepositorySuite) TestBrokerDeleteShouldReturnErrorWhenFailToDeleteBroker() {
	expectedResult := &entities.ClusterEntity{

		Name:        "Test Broker",
		Description: "Test Description",
		Host:        "localhost",
		Port:        1234,
		User:        "test_user",
		Password:    "test_password", Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{
		"id",
		"created_at",
		"updated_at",
		"name",
		"description",
		"host",
		"port",
		"user",
		"password",
	}).AddRow(
		expectedResult.ID,
		expectedResult.CreatedAt,
		expectedResult.UpdatedAt,
		expectedResult.Name,
		expectedResult.Description,
		expectedResult.Host,
		expectedResult.Port,
		expectedResult.User,
		expectedResult.Password,
	)

	brokerid := uint(10)

	brs.mock.ExpectQuery("").WillReturnRows(rows)

	brs.mock.ExpectBegin()
	brs.mock.ExpectRollback()
	//brs.mock.ExpectQuery("DELETE").WillReturnError(errors.New("generic delete error"))
	err := brs.SUT.DeleteCluster(brokerid, context.TODO())
	assert.NotNil(brs.T(), err)
	assert.Equal(brs.T(), "fail to delete broker", err.Error())
}
