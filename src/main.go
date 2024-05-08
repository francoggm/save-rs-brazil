package main

import (
	"log"

	"github.com/francoggm/save-rs-brazil/config"
	"github.com/francoggm/save-rs-brazil/infra/database"
	"github.com/francoggm/save-rs-brazil/infra/server"
)

func main() {
	cfg := config.New()

	db, err := database.New(cfg)
	if err != nil {
		log.Panic(err)
	}
	defer db.Conn.Close()

	err = db.Conn.Ping()
	if err != nil {
		log.Panic(err)
	}

	if err = server.Run(cfg, db); err != nil {
		log.Panic(err)
	}
}
