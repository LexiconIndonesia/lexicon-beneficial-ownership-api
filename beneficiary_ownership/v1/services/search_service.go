package beneficiary_ownership_v1

import (
	"context"
	models "lexicon/bo-api/beneficiary_ownership/v1/models"

	"github.com/jackc/pgx/v5"
)

var emptySearchResult []models.SearchResultModel

func searchByQuery(ctx context.Context, tx pgx.Tx, searchRequest models.SearchRequest) ([]models.SearchResultModel, error) {

	var itemCount int

	row := tx.QueryRow(ctx, "SELECT COUNT(id) as cnt FROM cases")
	err := row.Scan(&itemCount)

	if err != nil {
		return emptySearchResult, err
	}

	if itemCount == 0 {
		return emptySearchResult, nil
	}

	// rows, err := tx.Query(ctx, "SELECT id, subject, subject_type, person_in_charge, benificiary_ownership, nation, type, year FROM cases ORDER BY id DESC LIMIT $1 OFFSET $2", 10, 0)

	// if err != nil {
	// 	return emptySearchResult, err
	// }

	return emptySearchResult, nil // change to result of query
}
