package db

func (queueType QueueType) String() string {
	switch queueType {
	case QueueTypeClassic:
		return "topic"
	case QueueTypeQuorum:
		return "direct"

	default:
		return ""
	}
}

func ParseQueueType (queueType String) QueueType {
	switch queueType {
	case "topic":
		return QueueTypeClassic
	case "direct":
		return QueueTypeQuorum

	default:
		return QueueTypeClassic
	}
}