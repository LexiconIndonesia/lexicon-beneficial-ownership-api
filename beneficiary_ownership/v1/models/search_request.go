package beneficiary_ownership_v1_models

type SearchRequest struct {
	Query       string   `json:"query"`
	SubjectType string   `json:"subject_type"`
	Years       []string `json:"years"`
	Type        string   `json:"type"`
	Nation      string   `json:"nation"`
	Page        int64    `json:"page"`
	LastId      string   `json:"last_id"`
}
