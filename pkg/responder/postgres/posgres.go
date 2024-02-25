package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func NewDatabase(dsn string) (db *sql.DB, err error) {
	db, err = sql.Open("postgres", dsn)
	return
}
