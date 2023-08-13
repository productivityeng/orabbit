package repository

import (
	"context"
	"errors"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/productivityeng/orabbit/user/dto"
	userEntities "github.com/productivityeng/orabbit/user/entities"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepositoryMySql struct {
	Db *gorm.DB
}

func NewUserRepositoryMySql(Db *gorm.DB) *UserRepositoryMySql {
	err := Db.AutoMigrate(&userEntities.UserEntity{})
	if err != nil {
		log.WithError(err).Panic("erro migrating entity broker")
	}
	return &UserRepositoryMySql{
		Db: Db,
	}
}

// /CreateUser store a new broker in storage with provided parameter
func (repo *UserRepositoryMySql) CreateUser(userToCreate *userEntities.UserEntity) (*userEntities.UserEntity, error) {
	tx := repo.Db.Save(userToCreate)
	if tx.Error != nil {
		log.WithError(tx.Error).WithField("request", userToCreate).Error("Erro when trying save")
		tx.Rollback()
		return nil, tx.Error
	}
	return userToCreate, nil
}

// ListUsers retrieve a lista of broker with paginated options
func (repo *UserRepositoryMySql) ListUsers(brokerId int32, pageSize int, pageNumber int, ctx context.Context) (*contracts.PaginatedResult[dto.GetUserResponse], error) {
	entryFields := log.Fields{"pageSize": pageSize, "pageNumber": pageNumber}
	var result contracts.PaginatedResult[dto.GetUserResponse]
	result.PageSize = pageSize
	result.PageNumber = pageNumber

	offset := (pageNumber - 1) * pageSize

	var resultsFromDb []*userEntities.UserEntity

	err := repo.Db.WithContext(ctx).Where(userEntities.UserEntity{ClusterId: brokerId}).Offset(offset).Limit(pageSize).Find(&resultsFromDb).Error

	if err != nil {
		log.WithError(err).WithFields(entryFields).Error("error trying to query items for users")
		return nil, err
	}

	tx := repo.Db.WithContext(ctx).Model(&userEntities.UserEntity{}).Count(&result.TotalItems)
	if tx.Error != nil {
		log.WithError(tx.Error).WithFields(entryFields).Error("error trying to get count items for users")
	}

	for _, user := range resultsFromDb {
		userDto := dto.GetUserResponseFromUserEntity(user)
		result.Result = append(result.Result, &userDto)
	}

	return &result, nil
}

// DeleteUser soft delete a broker with a provided brokerId
func (repo *UserRepositoryMySql) DeleteUser(userId int32, ctx context.Context) error {
	fields := log.Fields{"userId": userId}
	var broker = userEntities.UserEntity{Id: userId}
	err := repo.Db.WithContext(ctx).First(&broker)
	if err.Error != nil {
		errorMsg := "broker id cound't not be found"
		log.WithFields(fields).WithError(err.Error).Error(errorMsg)
		return errors.New(errorMsg)
	}
	log.WithFields(fields).Infof("user founded, trying delete")
	err = repo.Db.Delete(&broker)
	if err.Error != nil {
		errorMsg := "fail to delete user"
		log.WithFields(fields).WithError(err.Error).Error(errorMsg)

		return errors.New(errorMsg)
	}
	log.WithFields(fields).Info("user deleted successfully")
	return nil
}

// GetUser retrieve a broker with a provided brokerId
func (repo *UserRepositoryMySql) GetUser(clusterId int32, userId int32, ctx context.Context) (*userEntities.UserEntity, error) {
	fields := log.Fields{"brokerId": userId}
	var user = userEntities.UserEntity{}
	err := repo.Db.WithContext(ctx).Where(&userEntities.UserEntity{Id: userId, ClusterId: clusterId}).First(&user)

	if err.Error != nil {
		errorMsg := "user id cound't not be found"
		log.WithFields(fields).WithError(err.Error).Error(errorMsg)
		return nil, errors.New(errorMsg)
	}

	return &user, nil
}

func (repo *UserRepositoryMySql) CheckIfUserExistsForCluster(cluster int32, username string, ctx context.Context) (bool, error) {
	count := int64(0)
	err := repo.Db.WithContext(ctx).Model(&userEntities.UserEntity{}).Where("cluster_id = ? and username = ?", cluster, username).Count(&count).Limit(1)
	if err.Error != nil {
		return false, err.Error
	}
	log.WithContext(ctx).Warn(count)
	return count > 0, nil
}
