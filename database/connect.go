package database

import (
	"backend_golang/config"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
