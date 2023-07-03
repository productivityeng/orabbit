package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	dsn := "root:password@tcp(127.0.0.1:3306)/orabbit_inventory?charset=utf8mb4&parseTime=True&loc=Local"

	Database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.WithError(err).Fatal("Error when trying open a first database connection")
	}
	Db = Database
}
