package dto

type GetQueueResponse struct {
	ID           uint                   `json:"ID"`
	ClusterID    uint                   `json:"ClusterId"`
	Name         string                 `json:"Name"`
	VHost        string                 `json:"VHost"`
	Type         string                 `json:"Type"`
	IsInCluster  bool                   `json:"IsInCluster"`
	IsInDatabase bool                   `json:"IsInDatabase"`
	Arguments    map[string]interface{} `json:"Arguments"`
	Durable      bool                   `json:"Durable"`
}

type GetQueueResponseList []GetQueueResponse

func (Items GetQueueResponseList) GetByName(name string) *GetQueueResponse {
	for _, item := range Items {
		if item.Name == name {
			return &item
		}
	}
	return nil
}
