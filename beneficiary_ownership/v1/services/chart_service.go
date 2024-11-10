package bo_v1_services

import (
	"context"
	"lexicon/bo-api/beneficiary_ownership"
	models "lexicon/bo-api/beneficiary_ownership/v1/models"
)

func GetChartData(ctx context.Context) (models.ChartsModel, error) {
	tx, err := beneficiary_ownership.Pool.Begin(ctx)

	if err != nil {
		return models.ChartsModel{}, err
	}

	list, err := models.ChartData(ctx, tx)

	if err != nil {
		return models.ChartsModel{}, err
	}

	tx.Commit(ctx)

	return list, nil
}
