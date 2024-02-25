package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://root:root@localhost:5432/library?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}
