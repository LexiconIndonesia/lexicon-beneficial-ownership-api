package bo_v1_models

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
	"gopkg.in/guregu/null.v4"
)

type CaseType int

func (c CaseType) String() string {
	return [...]string{"verdict", "blacklist", "sanction"}[c-1]
}

func newCaseType(s string) CaseType {
	switch s {
	case "verdict":
		return verdict
	case "blacklist":
		return blacklist
	case "sanction":
		return sanction
	default:
		return 0
	}
}

const (
	verdict CaseType = iota + 1
	blacklist
	sanction
)

type CaseStatus int

const (
	deleted CaseStatus = iota
	validated
	draft
)

func (c CaseStatus) String() string {
	return [...]string{"deleted", "validated", "draft"}[c]
}

type SubjectTypeInt int

const (
	individual SubjectTypeInt = iota + 1
	company
	organization
)

func newSubjectType(s string) SubjectTypeInt {

	switch s {
	case "individual":
		return individual
	case "company":
		return company
	case "organization":
		return organization
	default:
		return 0
	}
}
func (s SubjectTypeInt) String() string {
	return [...]string{"individual", "company", "organization"}[s-1]
}

type DetailResultModel struct {
	ID                   ulid.ULID   `json:"id"`
	Subject              string      `json:"subject"`
	SubjectType          string      `json:"subject_type"`
	PersonInCharge       null.String `json:"person_in_charge"`
	BenificiaryOwnership null.String `json:"benificiary_ownership"`
	CaseDate             null.Time   `json:"date"`
	DecisionNumber       null.String `json:"decision_number"`
	Source               string      `json:"source"`
	Link                 string      `json:"link"`
	Nation               string      `json:"nation"`
	PunishmentDuration   null.String `json:"punishment_duration"`
	Type                 string      `json:"type"`
	Year                 string      `json:"year"`
	Summary              string      `json:"summary"`
	Status               string      `json:"status"`
	CreatedAt            null.Time   `json:"created_at"`
	UpdatedAt            null.Time   `json:"updated_at"`
}

var emptyDetail DetailResultModel

func GetDetailById(ctx context.Context, tx pgx.Tx, id string) (DetailResultModel, error) {

	log.Info().Msg("Start getting detail by id: " + id)
	query := `
	SELECT id, subject, subject_type, person_in_charge, benificiary_ownership, case_date, decision_number, source, link, nation, punishment_start, punishment_end, case_type, year, summary, status, created_at, updated_at
	FROM cases
	WHERE id = $1
	AND status = $2
	LIMIT 1
	`

	log.Info().Msg("Executing query: " + query)
	row := tx.QueryRow(ctx, query, id, validated)

	temp := struct {
		ID                   ulid.ULID
		Subject              string
		SubjectType          SubjectTypeInt
		PersonInCharge       null.String
		BenificiaryOwnership null.String
		CaseDate             null.Time
		DecisionNumber       null.String
		Source               string
		Link                 string
		Nation               string
		PunishmentStartDate  null.Time
		PunishmentEndDate    null.Time
		Type                 CaseType
		Year                 string
		Summary              string
		Status               CaseStatus
		CreatedAt            null.Time
		UpdatedAt            null.Time
	}{}
	err := row.Scan(&temp.ID, &temp.Subject, &temp.SubjectType, &temp.PersonInCharge, &temp.BenificiaryOwnership, &temp.CaseDate, &temp.DecisionNumber, &temp.Source, &temp.Link, &temp.Nation, &temp.PunishmentStartDate, &temp.PunishmentEndDate, &temp.Type, &temp.Year, &temp.Summary, &temp.Status, &temp.CreatedAt, &temp.UpdatedAt)

	if err != nil {
		log.Info().Msg("Data Not Found")

		return emptyDetail, err
	}
	// mapping temp to result
	result := DetailResultModel{
		ID:                   temp.ID,
		Subject:              temp.Subject,
		SubjectType:          temp.SubjectType.String(),
		PersonInCharge:       temp.PersonInCharge,
		BenificiaryOwnership: temp.BenificiaryOwnership,
		CaseDate:             temp.CaseDate,
		DecisionNumber:       temp.DecisionNumber,
		Source:               temp.Source,
		Link:                 temp.Link,
		Nation:               temp.Nation,
		PunishmentDuration:   null.NewString(temp.PunishmentStartDate.Time.Format("02 Jan 2006")+" - "+temp.PunishmentEndDate.Time.Format("02 Jan 2006"), temp.PunishmentStartDate.Valid && temp.PunishmentEndDate.Valid),
		Type:                 temp.Type.String(),
		Year:                 temp.Year,
		Summary:              temp.Summary,
		Status:               temp.Status.String(),
		CreatedAt:            temp.CreatedAt,
		UpdatedAt:            temp.UpdatedAt,
	}

	log.Info().Msg("Finish getting detail by id: " + id)
	log.Info().Msg("Data Found")
	return result, nil
}
