package utils

type Pagination struct {
	Limit     int         `json:"limit"`
	Page      int         `json:"page"`
	LastPage  int         `json:"lastPage"`
	TotalRows int64       `json:"totalRows"`
	Sort      string      `json:"sort"`
	SortBy    string      `json:"sortBy"`
	Keyword   string      `json:"keyword"`
	FromRow   int         `json:"fromRow"`
	ToRow     int         `json:"toRow"`
	Items     interface{} `json:"items"`
}
