package beneficiary_ownership_v1_models

import (
	"context"
	commonModels "lexicon/bo-api/common/models"
	"math"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
	"gopkg.in/guregu/null.v4"
)

type SearchResultModel struct {
	ID                   ulid.ULID   `json:"id"`
	Subject              string      `json:"subject"`
	SubjectType          string      `json:"subject_type"`
	PersonInCharge       null.String `json:"person_in_charge"`
	BenificiaryOwnership null.String `json:"benificiary_ownership"`
	Nation               string      `json:"nation"`
	Type                 string      `json:"type"`
	Year                 string      `json:"year"`
}

var emptyBaseModel commonModels.BasePaginationResponse

func SearchByRequest(ctx context.Context, tx pgx.Tx, searchRequest SearchRequest) (commonModels.BasePaginationResponse, error) {
	var itemCount int

	limit := 20
	offset := (int(searchRequest.Page) - 1) * limit
	log.Info().Msg("Start counting query")
	countQuery := `
	SELECT COUNT(id) as cnt
	FROM cases
	`
	if searchRequest.Query != "" {
		countQuery += "WHERE fulltext_search_index @@ phraseto_tsquery('english',$1)"
	} else {
		countQuery += "WHERE $1 = $1"
	}

	countQuery += `
	AND subject_type ~* $2
	AND year ~* $3
	AND type ~* $4
	AND nation ~* $5
	`
	log.Info().Msg("Executing query: " + countQuery)

	row := tx.QueryRow(ctx, countQuery, searchRequest.Query, searchRequest.SubjectType, normalizeYears(searchRequest.Years), searchRequest.Type, searchRequest.Nation)
	err := row.Scan(&itemCount)
	log.Info().Msg("Finish counting query")

	if err != nil {
		return emptyBaseModel, err
	}

	if itemCount == 0 {
		return emptyBaseModel, nil
	}

	log.Info().Msg("Start searching query")
	searchQuery := `
	SELECT id, subject, subject_type, person_in_charge, benificiary_ownership, nation, type, year`

	if searchRequest.Query != "" {
		searchQuery += ", ts_rank_cd(fulltext_search_index, phraseto_tsquery('english', $1), 32 /* rank/(rank+1) */ ) AS rank "
	} else {
		searchQuery += ", 0 AS rank "
	}

	searchQuery += " FROM cases "

	if searchRequest.Query != "" {
		searchQuery += "WHERE fulltext_search_index @@ phraseto_tsquery('english', $1)"
	} else {
		searchQuery += "WHERE $1 = $1 "
	}

	searchQuery += `
	AND subject_type ~* $2
	AND year ~* $3
	AND type ~* $4
	AND nation ~* $5
	`

	if searchRequest.Query != "" {
		searchQuery += "ORDER BY rank DESC"
	}

	searchQuery += " LIMIT $6 OFFSET $7 "

	log.Info().Msg("Executing query: " + searchQuery)

	rows, err := tx.Query(ctx, searchQuery, searchRequest.Query, searchRequest.SubjectType, normalizeYears(searchRequest.Years), searchRequest.Type, searchRequest.Nation, limit, offset)

	log.Info().Msg("Finish searching query")
	if err != nil {
		log.Error().Err(err).Msg("Error querying database")
		return emptyBaseModel, err
	}

	defer rows.Close()

	var searchResults []SearchResultModel

	for rows.Next() {
		var rank float64
		var searchResult SearchResultModel
		err = rows.Scan(&searchResult.ID, &searchResult.Subject, &searchResult.SubjectType, &searchResult.PersonInCharge, &searchResult.BenificiaryOwnership, &searchResult.Nation, &searchResult.Type, &searchResult.Year, &rank)

		if err != nil {
			return emptyBaseModel, err
		}

		searchResults = append(searchResults, searchResult)
	}

	if len(searchResults) == 0 {
		return emptyBaseModel, nil
	}

	metaResponse := commonModels.MetaResponse{
		CurrentPage: searchRequest.Page,
		LastPage:    int64(math.Ceil(float64(itemCount) / float64(limit))),
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
