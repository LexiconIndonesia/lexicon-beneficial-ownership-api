package beneficiary_ownership_v1_models

type SearchRequest struct {
	Query       string `json:"query"`
	SubjectType string `json:"subject_type"`
	Year        string `json:"year"`
	Type        string `json:"type"`
	Nation      string `json:"nation"`
	Page        int    `json:"page"`
}
