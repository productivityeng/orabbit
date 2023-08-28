package dto

type GetQueueResponse struct {
	ID          uint   `json:"ID"`
	ClusterID   uint   `json:"ClusterId"`
	Name        string `json:"Name"`
	VirtualHost string `json:"VirtualHost"`
	Type        string `json:"Type"`
}
