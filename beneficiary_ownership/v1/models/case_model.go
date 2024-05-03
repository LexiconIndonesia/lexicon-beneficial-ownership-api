package beneficiary_ownership_v1_models

import (
	"context"

	"github.com/jackc/pgx/v5"
)

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

var emptyDetail DetailResultModel

func GetDetailById(ctx context.Context, tx pgx.Tx, id string) (DetailResultModel, error) {
	var result DetailResultModel

	query := `
	SELECT id, subject, subject_type, person_in_charge, benificiary_ownership, date, decision_number, source, link, nation, punishment_duration, type, year, summary
	FROM cases
	WHERE id = $1
	LIMIT 1
	`

	row := tx.QueryRow(ctx, query, id)
	err := row.Scan(&result.ID, &result.Subject, &result.SubjectType, &result.PersonInCharge, &result.BenificiaryOwnership, &result.Date, &result.DecisionNumber, &result.Source, &result.Link, &result.Nation, &result.PunishmentDuration, &result.Type, &result.Year, &result.Summary)

	if err != nil {
		return emptyDetail, err
	}

	return result, nil
}
