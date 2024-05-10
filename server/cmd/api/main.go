package main

import (
	"log"

	"github.com/francoggm/save-rs-brazil/config"
	"github.com/francoggm/save-rs-brazil/internal/server"
	"github.com/francoggm/save-rs-brazil/pkg/database/postgres"
)

func main() {
	cfg := config.New()

	db, err := postgres.New(cfg)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}

	sv := server.New(cfg, db)
	if err = sv.Run(); err != nil {
		log.Panic(err)
	}
}
