package dto

import "github.com/productivityeng/orabbit/db"

// GetExchangeDto is a dto that represents a exchange that is created in the cluster and in the database
type GetExchangeDto struct {
	Id             int                    `json:"Id"`
	Name           string                 `json:"Name"`
	ClusterId      int                    `json:"ClusterId"`
	Internal       bool                   `json:"Internal"`
	Durable        bool                   `json:"Durable"`
	Arguments      map[string]interface{} `json:"Arguments"`
	Lockers        []db.LockerExchangeModel `json:"Lockers"`
	VHost          string                 `json:"VHost"`
	Type           string                 `json:"Type"`
	IsInCluster   bool                   `json:"IsInCluster"`
	IsInDatabase  bool                   `json:"IsInDatabase"`
}

