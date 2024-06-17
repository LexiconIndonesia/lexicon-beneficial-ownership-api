package bo_v1_models

type SearchRequest struct {
	Query        string   `json:"query"`
	SubjectTypes []string `json:"subject_type"`
	Years        []string `json:"years"`
	Types        []string `json:"type"`
	Nations      []string `json:"nation"`
	Page         int64    `json:"page"`
}
