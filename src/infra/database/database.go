package database

import (
	"database/sql"
	"fmt"

	"github.com/francoggm/save-rs-brazil/config"
	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sql.DB
}

func New(cfg *config.Config) (*DB, error) {
	uri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DB)
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	return &DB{
		Conn: db,
	}, nil
}
