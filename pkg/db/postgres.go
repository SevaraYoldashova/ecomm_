package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectPostgres() (*sql.DB, error) {
	//connStr := "host=localhost port=5432 user=postgres password=postgres dbname=ecommerce sslmode=disable"
	connStr := "host=localhost port=5432 user=postgres password=sev21_ara dbname=ecommerce sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
