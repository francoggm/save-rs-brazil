package server

import (
	"github.com/francoggm/save-rs-brazil/config"
	"github.com/francoggm/save-rs-brazil/infra/database"
	"github.com/francoggm/save-rs-brazil/infra/server/handlers"
	"github.com/gofiber/fiber/v3"
)

func Run(cfg *config.Config, db *database.DB) error {
	app := fiber.New()

	handlers.RescueHandlers(app, db)
	handlers.FindAnimalHandlers(app, db)
	handlers.DonationHandlers(app, db)
	handlers.CleaningHandlers(app, db)
	handlers.MissingHandlers(app, db)

	return app.Listen("0.0.0.0:" + cfg.Port)
}
