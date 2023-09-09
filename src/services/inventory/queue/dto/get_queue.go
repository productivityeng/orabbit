package dto

type GetQueueResponse struct {
	ID           uint   `json:"ID"`
	ClusterID    uint   `json:"ClusterId"`
	Name         string `json:"Name"`
	VHost        string `json:"VHost"`
	Type         string `json:"Type"`
	IsRegistered bool   `json:"IsRegistered"`
}
