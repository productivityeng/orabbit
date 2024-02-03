package dto


type CreateExchangeDto struct { 
	Name string `json:"Name"`
	Type string `json:"Type"`
	ClusterId int `json:"ClusterId"`
	VirtualHostId int `json:"VirtualHostId"`
	Internal bool `json:"Internal"`
	Durable bool `json:"Durable"`
	Arguments map[string]interface{} `json:"Arguments"`
}