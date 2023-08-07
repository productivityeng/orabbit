package dto

import userEntities "github.com/productivityeng/orabbit/user/entities"

func GetUserResponseFromUserEntity(user *userEntities.UserEntity) GetUserResponse {
	return GetUserResponse{
		Id:           user.Id,
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
		BrokerId:     user.BrokerId,
	}
}
