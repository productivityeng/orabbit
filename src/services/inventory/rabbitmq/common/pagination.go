package common

type PageParam struct {
	PageSize   int `json:"PageSize" binding:"required,gt=0" default:"10"`
	PageNumber int `json:"PageNumber" binding:"required,gt=0" default:"1"`
}
