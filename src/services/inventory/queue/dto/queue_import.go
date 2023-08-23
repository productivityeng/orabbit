package dto

type QueueImportRequest struct {
	ClusterId int32  `json:"ClusterId"`
	QueueName string `json:"QueueName"`
}
