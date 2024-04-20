package module

import (
	"database/sql"
	"errors"
)

var (
	mysqlClient *sql.DB
)

func SetDatabase(database *sql.DB) error {

	if database == nil {
		return errors.New("cannot assign nil database")
	}
	mysqlClient = database
	return nil
}
