package bo_v1_services

import (
	"context"
	"lexicon/bo-api/beneficiary_ownership"
	models "lexicon/bo-api/beneficiary_ownership/v1/models"
	baseModel "lexicon/bo-api/common/models"
)

func Search(ctx context.Context, searchRequest models.SearchRequest) (baseModel.BasePaginationResponse, error) {
	tx, err := beneficiary_ownership.Pool.Begin(ctx)

	if err != nil {
		return baseModel.BasePaginationResponse{}, err
	}

	list, err := models.SearchByRequest(ctx, tx, searchRequest)

	if err != nil {
		return baseModel.BasePaginationResponse{}, err
	}

	tx.Commit(ctx)

	return list, nil // change to result of query
}
