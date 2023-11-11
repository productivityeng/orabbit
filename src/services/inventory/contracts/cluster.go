package contracts

type CreateClusterRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required" `
	Host        string `json:"host" binding:"required"`
	Port        int32  `json:"port" binding:"required"`
	User        string `json:"user" binding:"required"`
	Password    string `json:"password" binding:"required"`
}
