package contracts

type GetUserResponse struct {
	Id           uint   `json:"Id"`
	Username     string `json:"Username"`
	PasswordHash string `json:"PasswordHash"`
	IsInCluster  bool   `json:"IsInCluster"`
	IsInDatabase bool   `json:"IsInDatabase"`
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
