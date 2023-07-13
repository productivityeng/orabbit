package contracts

type PaginatedResult[T any] struct {
	Result     []*T  `json:"result"`
	PageNumber int   `json:"pageNumber"`
	PageSize   int   `json:"pageSize"`
	TotalItems int64 `json:"totalItems"`
}
