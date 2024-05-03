package beneficiary_ownership_v1_models

import (
	"context"
	commonModels "lexicon/bo-api/common/models"
	"strings"

	"github.com/jackc/pgx/v5"
)

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

var emptyBaseModel commonModels.BasePaginationResponse

func SearchByRequest(ctx context.Context, tx pgx.Tx, searchRequest SearchRequest) (commonModels.BasePaginationResponse, error) {
	var itemCount int

	limit := 20

	countQuery := `
	SELECT COUNT(id) as cnt
	FROM cases
	WHERE search_index @@ to_tsquery($1)
	AND subject_type ~* $2
	AND year ~* $3
	AND type ~* $4
	AND nation ~* $5
	`

	row := tx.QueryRow(ctx, countQuery, searchRequest.Query, searchRequest.SubjectType, normalizeYears(searchRequest.Years), searchRequest.Type, searchRequest.Nation)
	err := row.Scan(&itemCount)

	if err != nil {
		return emptyBaseModel, err
	}

	if itemCount == 0 {
		return emptyBaseModel, nil
	}

	searchQuery := `
	SELECT id, subject, subject_type, person_in_charge, benificiary_ownership, nation, type, year FROM cases
	WHERE search_index @@ to_tsquery($1)
	AND subject_type ~* $2
	AND year ~* $3
	AND type ~* $4
	AND nation ~* $5
	AND id > $6
	ORDER BY id DESC LIMIT $2

	`
	rows, err := tx.Query(ctx, searchQuery, searchRequest.Query, searchRequest.SubjectType, normalizeYears(searchRequest.Years), searchRequest.Type, searchRequest.Nation, searchRequest.LastId, limit)

	if err != nil {
		return emptyBaseModel, err
	}

	defer rows.Close()

	var searchResults []SearchResultModel

	for rows.Next() {
		var searchResult SearchResultModel
		err = rows.Scan(&searchResult.ID, &searchResult.Subject, &searchResult.SubjectType, &searchResult.PersonInCharge, &searchResult.BenificiaryOwnership, &searchResult.Nation, &searchResult.Type, &searchResult.Year)

		if err != nil {
			return emptyBaseModel, err
		}

		searchResults = append(searchResults, searchResult)
	}

	metaResponse := commonModels.MetaResponse{
		LastId:      searchResults[len(searchResults)-1].ID,
		CurrentPage: searchRequest.Page,
		LastPage:    int64(itemCount / limit),
		PerPage:     int64(limit),
		Total:       int64(itemCount),
	}

	baseResponse := commonModels.BasePaginationResponse{
		Data: searchResults,
		Meta: metaResponse,
	}

	return baseResponse, nil
}

func normalizeYears(years []string) string {
	return strings.Join(years, "|")
}
