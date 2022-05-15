package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func NewPostgres(connectionURL string) *sql.DB {
	db, err := sql.Open("postgres", connectionURL)
	if err != nil {
		panic(err)
	}

	return db
}
