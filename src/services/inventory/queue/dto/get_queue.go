package dto

type GetQueueResponse struct {
	ID          int32  `json:"ID"`
	ClusterID   int32  `json:"ClusterID"`
	Name        string `json:"Name"`
	VirtualHost string `json:"VirtualHost"`
	Type        string `json:"Type"`
}
