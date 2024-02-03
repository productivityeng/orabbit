package dto

import (
	"github.com/productivityeng/orabbit/db"
)

func GetUserResponseFromUserEntity(user *db.UserModel) GetUserResponse {
	return GetUserResponse{
		Id:           user.ID,
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
		ClusterId:    user.ClusterID,
	}
}
