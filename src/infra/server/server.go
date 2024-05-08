package server

import (
	"github.com/francoggm/save-rs-brazil/config"
	"github.com/francoggm/save-rs-brazil/infra/database"
	"github.com/gofiber/fiber"
)

func Run(cfg *config.Config, db *database.DB) error {
	app := fiber.New()

	return app.Listen("0.0.0.0:" + cfg.Port)
}