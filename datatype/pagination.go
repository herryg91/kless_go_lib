package datatype

import "math"

type Pagination struct {
	CurrentPage  int `json:"current_page"`
	TotalPages   int `json:"total_pages"`
	PerPage      int `json:"per_page"`
	TotalEntries int `json:"total_entries"`
	From         int `json:"from"`
	To           int `json:"to"`
}

func (Pagination) New(total_datas, page, size int) Pagination {
	return Pagination{
		CurrentPage:  page,
		TotalPages:   int(math.Ceil(float64(total_datas) / float64(size))),
		PerPage:      size,
		TotalEntries: int(total_datas),
		From:         (page-1)*size + 1,
		To:           page * size,
	}
}
