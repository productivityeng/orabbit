package user

import (
	"context"
	"errors"
	"fmt"
	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	"github.com/sirupsen/logrus"
)

type RabbitAccess struct {
	Host     string
	Port     int32
	Username string
	Password string
}

type GetUserHashRequest struct {
	RabbitAccess
	UserToRetrieveHash string
}

type CreateNewUserRequest struct {
	RabbitAccess
	UserToCreate            string
	PasswordOfUsertToCreate string
}

type CreateNewUserResult struct {
	PasswordHash string
}

type ListUserResult struct {
	PasswordHash string
	Name         string
}

type ListAllUsersRequest struct {
	RabbitAccess
}

type UserManagement interface {
	GetUserHash(request GetUserHashRequest, ctx context.Context) (string, error)
	CreateNewUser(request CreateNewUserRequest, ctx context.Context) (*CreateNewUserResult, error)
	ListAllUser(request ListAllUsersRequest) ([]ListUserResult, error)
}
type UserManagementImpl struct {
}

func NewUserManagement() *UserManagementImpl {

	return &UserManagementImpl{}
}

func (management *UserManagementImpl) ListAllUser(request ListAllUsersRequest) ([]ListUserResult, error) {
	rmqc, err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.Username, request.Password)
	if err != nil {
		logrus.WithError(err).Error("Error trying to connect to cluster")
		return nil, errors.New("[CLUSTER_CONNECT_FAIL]")
	}

	users, err := rmqc.ListUsers()
	var result []ListUserResult
	for _, user := range users {
		result = append(result, ListUserResult{
			PasswordHash: user.PasswordHash,
			Name:         user.Name,
		})
	}
	return result, nil
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
func (management *UserManagementImpl) CreateNewUser(request CreateNewUserRequest, ctx context.Context) (*CreateNewUserResult, error) {
	rmqc, err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.Username, request.Password)
	if err != nil {
		logrus.WithError(err).Error("Error trying to connect to cluster")
		return nil, errors.New("[CLUSTER_CONNECT_FAIL]")
	}
	logrus.WithField("request", request).Info("Creating new user in cluster")
	_, err = rmqc.PutUser(request.UserToCreate, rabbithole.UserSettings{
		Tags:     []string{"management"},
		Password: request.Password,
	})

	if err != nil {
		logrus.WithError(err).WithContext(ctx).Error("Error trying to create a new user")
		return nil, errors.New("[CLUSTER_CREATE-USER_FAIL]")
	}

	userInfo, err := rmqc.GetUser(request.UserToCreate)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("Fail to retrieve user hash password")
		return nil, errors.New("[CLUSTER_USERINFO_RETRIEVE_FAIL]")
	}

	return &CreateNewUserResult{PasswordHash: userInfo.PasswordHash}, nil
}
