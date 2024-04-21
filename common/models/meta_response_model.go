package models

type MetaResponse struct {
	CurrentPage int64 `json:"current_page"`
	From        int64 `json:"from"`
	LastPage    int64 `json:"last_page"`
	PerPage     int64 `json:"per_page"`
	To          int64 `json:"to"`
	Total       int64 `json:"total"`
}
