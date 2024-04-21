package beneficiary_ownership_v1_models

type DetailResultModel struct {
	ID                   string `json:"id"`
	Subject              string `json:"subject"`
	SubjectType          string `json:"subject_type"`
	PersonInCharge       string `json:"person_in_charge"`
	BenificiaryOwnership string `json:"benificiary_ownership"`
	Date                 string `json:"date"`
	DecisionNumber       string `json:"decision_number"`
	Source               string `json:"source"`
	Link                 string `json:"link"`
	Nation               string `json:"nation"`
	PunishmentDuration   string `json:"punishment_duration"`
	Type                 string `json:"type"`
	Year                 string `json:"year"`
	Summary              string `json:"summary"`
}
