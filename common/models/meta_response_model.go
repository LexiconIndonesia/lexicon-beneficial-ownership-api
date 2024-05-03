package models

type MetaResponse struct {
	CurrentPage int64 `json:"current_page"`
	LastPage    int64 `json:"last_page"`
	PerPage     int64 `json:"per_page"`
	Total       int64 `json:"total"`
}
