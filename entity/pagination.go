package entity

type Pagination struct {
	Total       int    `json:"total"`
	CurrentPage int    `json:"currentPage"`
	Posts       []Post `json:"posts"`
}
