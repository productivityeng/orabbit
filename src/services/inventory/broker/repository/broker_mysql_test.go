package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/productivityeng/orabbit/broker/entities"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"regexp"
	"testing"
	"time"
)

type BrokerRepositorySuite struct {
	suite.Suite
	conn   *sql.DB
	DB     *gorm.DB
	mock   sqlmock.Sqlmock
	SUT    *BrokerRepositoryMysqlImpl
	broker *entities.BrokerEntity
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

	brs.SUT = NewBrokerMysqlImpl(brs.DB)
	assert.IsType(brs.T(), &BrokerRepositoryMysqlImpl{}, brs.SUT)

	brs.broker = &entities.BrokerEntity{
		Name:        "TestBroker",
		Description: "Broker of test",
		Host:        "rabbitmq.com",
		Port:        999,
		User:        "rabbit",
		Password:    "rabbit",
	}
}

// TestCreateBroker check if can execute query correctly
func (brs *BrokerRepositorySuite) TestCreateBroker() {

	brs.mock.ExpectBegin()
	brs.mock.ExpectExec("").
		WillReturnResult(sqlmock.NewResult(1, 1))

	brs.mock.ExpectCommit()

	broker, err := brs.SUT.CreateBroker(brs.broker)

	assert.NoError(brs.T(), err)
	brokerGTZero := broker.Id >= 1
	assert.True(brs.T(), brokerGTZero)
}

// TestCreateBrokerError check if can deal with error in create a broker
func (brs *BrokerRepositorySuite) TestCreateBrokerError() {
	brs.mock.ExpectBegin()
	brs.mock.ExpectExec("").
		WillReturnResult(sqlmock.NewErrorResult(errors.New("error executing query proposital")))

	broker, err := brs.SUT.CreateBroker(brs.broker)
	assert.Error(brs.T(), err)
	assert.Nil(brs.T(), broker)

}

func (brs *BrokerRepositorySuite) TestListBroker() {

	expectedResult := &entities.BrokerEntity{
		Id:          1,
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
		expectedResult.Id,
		expectedResult.CreatedAt,
		expectedResult.UpdatedAt,
		expectedResult.Name,
		expectedResult.Description,
		expectedResult.Host,
		expectedResult.Port,
		expectedResult.User,
		expectedResult.Password,
	).AddRow(
		expectedResult.Id,
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
	result, err := brs.SUT.ListBroker(pageSize, pageNumber)
	assert.Nil(brs.T(), err)
	assert.Equal(brs.T(), expectedResult, result.Result[0])
	assert.Equal(brs.T(), pageSize, result.PageSize)
	assert.Equal(brs.T(), pageNumber, result.PageNumber)
}

func (brs *BrokerRepositorySuite) TestListBrokerErrorTryingToRetrieveResult() {

	pageSize := 1
	pageNumber := 2

	brs.mock.ExpectQuery("").WillReturnError(errors.New("genericerro"))
	result, err := brs.SUT.ListBroker(pageSize, pageNumber)
	assert.NotNil(brs.T(), err)
	assert.Nil(brs.T(), result)
}

func (brs *BrokerRepositorySuite) TestBrokerDeleteShouldReturnErrorWhenBrokerNotExists() {
	brokerid := int32(10)
	brs.mock.ExpectQuery("SELECT").WillReturnError(errors.New("genericerro"))

	err := brs.SUT.DeleteBroker(brokerid, context.TODO())
	assert.NotNil(brs.T(), err)
	assert.Equal(brs.T(), err.Error(), "broker id cound't not be found")

}

func (brs *BrokerRepositorySuite) TestBrokerDeleteShouldReturnErrorWhenFailToDeleteBroker() {
	expectedResult := &entities.BrokerEntity{
		Id: 1,

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
		expectedResult.Id,
		expectedResult.CreatedAt,
		expectedResult.UpdatedAt,
		expectedResult.Name,
		expectedResult.Description,
		expectedResult.Host,
		expectedResult.Port,
		expectedResult.User,
		expectedResult.Password,
	)

	brokerid := int32(10)

	brs.mock.ExpectQuery("").WillReturnRows(rows)

	brs.mock.ExpectBegin()
	brs.mock.ExpectRollback()
	//brs.mock.ExpectQuery("DELETE").WillReturnError(errors.New("generic delete error"))
	err := brs.SUT.DeleteBroker(brokerid, context.TODO())
	assert.NotNil(brs.T(), err)
	assert.Equal(brs.T(), "fail to delete broker", err.Error())
}

func (brs *BrokerRepositorySuite) TestBrokerDeleteShouldSuccess() {
	expectedResult := &entities.BrokerEntity{
		Id:          1,
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
		expectedResult.Id,
		expectedResult.CreatedAt,
		expectedResult.UpdatedAt,
		expectedResult.Name,
		expectedResult.Description,
		expectedResult.Host,
		expectedResult.Port,
		expectedResult.User,
		expectedResult.Password,
	)

	brokerid := int32(10)

	brs.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `broker` WHERE `broker`.`deleted_at` IS NULL AND `broker`.`id` = ? ORDER BY `broker`.`id` LIMIT 1")).WillReturnRows(rows)
	brs.mock.ExpectBegin()
	brs.mock.ExpectExec(regexp.QuoteMeta("UPDATE `broker` SET `deleted_at`=? WHERE `broker`.`id` = ? AND `broker`.`deleted_at` IS NULL")).WillReturnResult(sqlmock.NewResult(1, 1))
	brs.mock.ExpectCommit()
	err := brs.SUT.DeleteBroker(brokerid, context.TODO())
	assert.Nil(brs.T(), err)
}

func (brs *BrokerRepositorySuite) TestGetBrokerShouldReturnErrorWhenBrokerNotExists() {
	brokerid := int32(10)
	brs.mock.ExpectQuery("SELECT").WillReturnError(errors.New("genericerro"))

	broker, err := brs.SUT.GetBroker(brokerid, context.TODO())
	assert.NotNil(brs.T(), err)
	assert.Equal(brs.T(), err.Error(), "broker id cound't not be found")
	assert.Nil(brs.T(), broker)
}

func (brs *BrokerRepositorySuite) TestGetBrokerShouldReturnSuccess() {
	expectedResult := &entities.BrokerEntity{
		Id: 1,

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
		expectedResult.Id,
		expectedResult.CreatedAt,
		expectedResult.UpdatedAt,
		expectedResult.Name,
		expectedResult.Description,
		expectedResult.Host,
		expectedResult.Port,
		expectedResult.User,
		expectedResult.Password,
	)

	brokerid := int32(10)

	brs.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `broker` WHERE `broker`.`deleted_at` IS NULL AND `broker`.`id` = ? ORDER BY `broker`.`id` LIMIT 1")).WillReturnRows(rows)
	broker, err := brs.SUT.GetBroker(brokerid, context.TODO())
	assert.Nil(brs.T(), err)
	assert.Equal(brs.T(), expectedResult, broker)
}
func TestBrokerRepositorySuit(t *testing.T) {
	suite.Run(t, new(BrokerRepositorySuite))
}
