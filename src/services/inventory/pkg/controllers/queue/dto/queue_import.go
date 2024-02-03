package dto

type QueueImportRequest struct {
	QueueName string `json:"QueueName"`
	Type      string `json:"Type"`
	VirtualHostId int `json:"VirtualHostId"`
}
