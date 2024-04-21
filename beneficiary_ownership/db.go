package beneficiary_ownership

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pool *pgxpool.Pool
)

func SetDatabase(newPool *pgxpool.Pool) error {

	if newPool == nil {
		return errors.New("cannot assign nil database")
	}
	pool = newPool
	return nil
}
