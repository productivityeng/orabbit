package user

import (
	"context"
	"errors"
	"fmt"

	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	"github.com/productivityeng/orabbit/contracts"
	"github.com/sirupsen/logrus"
)



type UserManagementImpl struct {
}

func NewUserManagement() *UserManagementImpl {

	return &UserManagementImpl{}
}

func (management *UserManagementImpl) CreateNewUserWithHashPassword(request contracts.CreateNewUserWithHashPasswordRequest, ctx context.Context) (*contracts.CreateNewUserResult, error) {
	rmqc, err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.RabbitAccess.Username, request.RabbitAccess.Password)
	if err != nil {
		logrus.WithError(err).Error(fmt.Sprintf("Erro ao remover usuario %s do cluster %s", request.Username, request.Host))
		return nil, err
	}

	_, err = rmqc.PutUser(request.UsernameToCreate, rabbithole.UserSettings{
		Name:         request.UsernameToCreate,
		Tags:         []string{"management"},
		PasswordHash: request.PasswordHash,
	})

	if err != nil {
		logrus.WithError(err).Error("Erro ao criar usuario usando hash")
		return nil, err
	}
	_, err = rmqc.UpdatePermissionsIn("/", request.UsernameToCreate, rabbithole.Permissions{
		Configure: ".*",
		Write:     ".*",
		Read:      ".*",
	})

	if err != nil {
		logrus.WithError(err).Error("Erro ao adicionar permissoes no vhost")
		return nil, err
	}

	_, err = rmqc.UpdateTopicPermissionsIn("/", request.UsernameToCreate, rabbithole.TopicPermissions{
		Exchange: ".*",
		Write:    ".*",
		Read:     ".*",
	})
	if err != nil {
		logrus.WithError(err).Error("Erro ao adicionar topic permissions no vhost")
		return nil, err
	}

	return &contracts.CreateNewUserResult{PasswordHash: request.PasswordHash}, nil
}

func (management *UserManagementImpl) DeleteUser(request contracts.DeleteUserRequest, ctx context.Context) error {
	rmqc, err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.RabbitAccess.Username, request.RabbitAccess.Password)
	_, err = rmqc.DeleteUser(request.Username)
	if err != nil {
		logrus.WithError(err).Error(fmt.Sprintf("Erro ao remover usuario %s do cluster %s", request.Username, request.Host))
		return err
	}
	return nil
}

func (management *UserManagementImpl) ListAllUser(request contracts.ListAllUsersRequest) ([]contracts.ListUserResult, error) {
	rmqc, err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.Username, request.Password)
	if err != nil {
		logrus.WithError(err).Error("Error trying to connect to cluster")
		return nil, errors.New("[CLUSTER_CONNECT_FAIL]")
	}

	users, err := rmqc.ListUsers()
	var result []contracts.ListUserResult
	for _, user := range users {
		result = append(result, contracts.ListUserResult{
			PasswordHash: user.PasswordHash,
			Name:         user.Name,
		})
	}
	return result, nil
}

func (management *UserManagementImpl) GetUserHash(request contracts.GetUserHashRequest, ctx context.Context) (string, error) {
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

func (management *UserManagementImpl) GetUserByName(request contracts.GetUserByNameRequest, ctx context.Context) (*rabbithole.UserInfo, error) {
	rmqc, err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.Username, request.Password)
	if err != nil {
		logrus.WithError(err).Error("Error trying to connect to cluster")
		return nil, errors.New("[CLUSTER_CONNECT_FAIL]")
	}

	userInfo, err := rmqc.GetUser(request.Username)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("Fail to retrieve user hash password")
		return nil, errors.New("[CLUSTER_USERINFO_RETRIEVE_FAIL]")
	}

	return userInfo, nil
}

func (management *UserManagementImpl) CreateNewUser(request contracts.CreateNewUserRequest, ctx context.Context) (*contracts.CreateNewUserResult, error) {
	rmqc, err := rabbithole.NewClient(fmt.Sprintf("http://%s:%d", request.Host, request.Port), request.Username, request.Password)
	if err != nil {
		logrus.WithError(err).Error("Error trying to connect to cluster")
		return nil, errors.New("[CLUSTER_CONNECT_FAIL]")
	}
	logrus.WithField("request", request).Info("Creating new user in cluster")
	_, err = rmqc.PutUser(request.UserToCreate, rabbithole.UserSettings{
		Tags:     []string{"management"},
		Password: request.PasswordOfUsertToCreate,
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

	return &contracts.CreateNewUserResult{PasswordHash: userInfo.PasswordHash}, nil
}
