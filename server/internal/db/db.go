package db

import (
	"database/sql"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func Connect(connectionString string) error {
	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	return DB.Ping()
}
