package repository

import (
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
	"testing"
)

type BrokerRepositorySuite struct {
	suite.Suite
	conn   *sql.DB
	DB     *gorm.DB
	mock   sqlmock.Sqlmock
	repo   *BrokerRepositoryMysqlImpl
	broker *entities.BrokerEntity
}

func (brs *BrokerRepositorySuite) SetupSuite() {
	log.SetLevel(log.FatalLevel)
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

	brs.repo = NewBrokerMysqlImpl(brs.DB)
	assert.IsType(brs.T(), &BrokerRepositoryMysqlImpl{}, brs.repo)

	brs.broker = &entities.BrokerEntity{
		Name:        "TestBroker",
		Description: "Broker of test",
		Host:        "rabbitmq.com",
		Port:        999,
		User:        "rabbit",
		Password:    "rabbit",
	}
}

func (brs *BrokerRepositorySuite) TestCreateBroker() {

	brs.mock.ExpectBegin()
	brs.mock.ExpectExec("INSERT INTO `broker` (.+) VALUES (.+)").
		WillReturnResult(sqlmock.NewResult(1, 1))

	brs.mock.ExpectCommit()

	broker, err := brs.repo.CreateBroker(brs.broker)

	assert.NoError(brs.T(), err)
	brokerGTZero := broker.Id >= 1
	assert.True(brs.T(), brokerGTZero)
}

func (brs *BrokerRepositorySuite) TestCreateBrokerError() {
	brs.mock.ExpectBegin()
	brs.mock.ExpectExec("INSERT INTO `broker` (.+) VALUES (.+)").
		WillReturnResult(sqlmock.NewErrorResult(errors.New("error executing query proposital")))
	brs.mock.ExpectCommit()

	broker, err := brs.repo.CreateBroker(brs.broker)
	assert.Error(brs.T(), err)
	assert.Nil(brs.T(), broker)

}

func TestBrokerRepositorySuit(t *testing.T) {
	suite.Run(t, new(BrokerRepositorySuite))
}
