package bo_v1_services

import (
	"context"
	"lexicon/bo-api/beneficiary_ownership"
	models "lexicon/bo-api/beneficiary_ownership/v1/models"
)

func GetDetail(ctx context.Context, id string) (models.DetailResultModel, error) {
	tx, err := beneficiary_ownership.Pool.Begin(ctx)

	if err != nil {
		return models.DetailResultModel{}, err
	}

	detail, err := models.GetDetailById(ctx, tx, id)

	if err != nil {
		return models.DetailResultModel{}, err
	}

	tx.Commit(ctx)

	return detail, nil
}
