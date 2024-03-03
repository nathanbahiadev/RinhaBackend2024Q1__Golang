package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func New(dsn string) (*sql.DB, error) {
	return sql.Open("postgres", dsn)
}
