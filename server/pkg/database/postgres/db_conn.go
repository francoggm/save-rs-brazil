package postgres

import (
	"fmt"

	"github.com/francoggm/save-rs-brazil/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(cfg *config.Config) (*sqlx.DB, error) {
	uri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DB)
	db, err := sqlx.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	return db, nil
}
