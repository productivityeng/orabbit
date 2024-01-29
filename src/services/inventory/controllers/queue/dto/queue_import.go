package dto

type QueueImportRequest struct {
	QueueName string `json:"QueueName"`
	Type      string `json:"Type"`
	VirtualHostName string `json:"VirtualHostName"`
}
