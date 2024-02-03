package dto

import "github.com/productivityeng/orabbit/db"

type GetVirtualHostDto struct {
	Id           int  `json:"Id"`
	ClusterId   int  `json:"ClusterId"`
	Description  string `json:"Description" binding:"required"`
	Name         string `json:"Name" binding:"required"`
	IsInDatabase bool `json:"IsInDatabase"`
	IsInCluster  bool `json:"IsInCluster"`
	DefaultQueueType string `json:"DefaultQueueType"`
	Tags []string `json:"Tags"`
	Lockers []db.LockerVirtualHostModel `json:"Lockers"`
}

type ImportVirtualHostRequest struct { 
	Name string `json:"Name" binding:"required"`
}

type SaveVirtualHostDto struct {
	ClusterId   int  `json:"ClusterId"`
	Description  string `json:"Description" binding:"required"`
	Name         string `json:"Name" binding:"required"`
	DefaultQueueType string `json:"DefaultQueueType"`
	Tags []string `json:"Tags"`
}
