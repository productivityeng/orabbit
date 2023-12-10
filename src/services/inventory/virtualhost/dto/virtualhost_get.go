package dto

type GetVirtualHostDto struct {
	Id           uint
	Description  string
	Name         string
	IsInDatabase bool
	IsInCluster  bool
}
