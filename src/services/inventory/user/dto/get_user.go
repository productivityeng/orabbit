package dto

import "github.com/productivityeng/orabbit/db"


type GetUserResponse struct {
	Id           int   `json:"Id"`
	ClusterId   int   `json:"ClusterId"`
	Username     string `json:"Username"`
	PasswordHash string `json:"PasswordHash"`
	IsInCluster  bool   `json:"IsInCluster"`
	IsInDatabase bool   `json:"IsInDatabase"`
	Lockers   	 []db.LockerUserModel 	`json:"Lockers"`
}

type GetUserResponseList []GetUserResponse

func (list GetUserResponseList) UserInListByName(username string) bool {
	for _, user := range list {
		if user.Username == username {
			return true
		}
	}
	return false
}
