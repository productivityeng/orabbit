package db

type QueueList []QueueModel
func (queueList QueueList) GetQueueFromListByName(queueName string) *QueueModel {
	for _, queueFromList := range queueList {
		if queueFromList.Name == queueName {
			return &queueFromList
		}
	}
	return nil
}

