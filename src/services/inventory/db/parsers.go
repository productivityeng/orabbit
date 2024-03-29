package db

func (queueType QueueType) String() string {
	switch queueType {
	case QueueTypeClassic:
		return "classic"
	case QueueTypeQuorum:
		return "direct"
		
	default:
		return ""
	}
}

func ParseQueueType (queueType String) QueueType {
	switch queueType {
	case "classic":
		return QueueTypeClassic
	case "direct":
		return QueueTypeQuorum

	default:
		return QueueTypeClassic
	}
}