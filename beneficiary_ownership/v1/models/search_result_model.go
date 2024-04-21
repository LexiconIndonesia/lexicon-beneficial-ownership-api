package beneficiary_ownership_v1_models

type SearchResultModel struct {
	ID                   string `json:"id"`
	Subject              string `json:"subject"`
	SubjectType          string `json:"subject_type"`
	PersonInCharge       string `json:"person_in_charge"`
	BenificiaryOwnership string `json:"benificiary_ownership"`
	Nation               string `json:"nation"`
	Type                 string `json:"type"`
	Year                 string `json:"year"`
}
