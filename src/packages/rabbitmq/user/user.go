package user

import (
	"context"
	"errors"
	"fmt"
	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	"github.com/sirupsen/logrus"
)

type GetUserHashRequest struct {
	Host               string
	Port               int32
	Username           string
	Password           string
	UserToRetrieveHash string
}
type UserManagement interface {
	GetUserHash(request GetUserHashRequest, ctx context.Context) (string, error)
}
type UserManagementImpl struct {
}

func NewUserManagement() *UserManagementImpl {

	return &UserManagementImpl{}
}

func (management *UserManagementImpl) GetUserHash(request GetUserHashRequest, ctx context.Context) (string, error) {
	rmqc, err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.Username, request.Password)
	if err != nil {
		logrus.WithError(err).Error("Error trying to connect to cluster")
		return "", errors.New("[CLUSTER_CONNECT_FAIL]")
	}

	userInfo, err := rmqc.GetUser(request.UserToRetrieveHash)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("Fail to retrieve user hash password")
		return "", errors.New("[CLUSTER_USERINFO_RETRIEVE_FAIL]")
	}

	return userInfo.PasswordHash, nil
}
