package dto

import userEntities "github.com/productivityeng/orabbit/user/entities"

func GetUserResponseFromUserEntity(user *userEntities.UserEntity) GetUserResponse {
	return GetUserResponse{
		Id:           user.ID,
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
		ClusterId:    user.ClusterId,
	}
}
