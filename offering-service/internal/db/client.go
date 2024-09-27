package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"offering-service/internal/config"
)

type FlatsStorage struct {
	db *sql.DB
}

func NewClient(cfg *config.Config) (*FlatsStorage, error) {
	db, err := sql.Open("postgres", cfg.PostgresConn)
	if err != nil {
		return &FlatsStorage{}, err
	}

	return &FlatsStorage{db: db}, nil
}