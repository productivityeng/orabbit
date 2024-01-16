package dto

type GetVirtualHostDto struct {
	Id           int 
	Description  string
	Name         string
	IsInDatabase bool
	IsInCluster  bool
}
