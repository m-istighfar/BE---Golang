package pagedto

type PageSortDto struct {
	Page   int    `form:"page" binding:"omitempty,min=1"`
	Limit  int    `form:"limit" binding:"omitempty,min=1,max=100"`
	Search string `form:"s"`
	SortBy string `form:"sort_by"`
}

type PageInfoDto struct {
	Page     int   `json:"page"`
	Limit    int   `json:"limit"`
	TotalRow int64 `json:"total_row"`
	HasNext  bool  `json:"has_next"`
	HasPrev  bool  `json:"has_prev"`
}
