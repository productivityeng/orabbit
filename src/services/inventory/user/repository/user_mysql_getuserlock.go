package repository

import (
	"context"

	locker "github.com/productivityeng/orabbit/locker/entities"
	"github.com/productivityeng/orabbit/user/dto"
	log "github.com/sirupsen/logrus"
)

func (repo *UserRepositoryMySql) GetUserLock(userId int64, ctx context.Context) (*dto.LockUserDto, error) {

	var locker *locker.LockerEntity

	err := repo.Db.WithContext(ctx).First(&locker, "DisabledAt is null and UserId = ?", userId)

	if err.Error != nil {
		log.WithContext(ctx).Error(err.Error)
		return nil, err.Error
	}

	return &dto.LockUserDto{
		UserId:     userId,
		Reason:     locker.Reason,
		Enabledt:   locker.EnabledAt,
		DisabledAt: locker.DisabledAt,
	}, nil

}
