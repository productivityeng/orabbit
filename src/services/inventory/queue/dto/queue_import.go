package dto

type QueueImportRequest struct {
	ClusterId uint   `json:"ClusterId"`
	QueueName string `json:"QueueName"`
}
