package handlers

import (
	"github.com/francoggm/save-rs-brazil/infra/database"
	"github.com/gofiber/fiber/v3"
)

func MissingHandlers(app *fiber.App, db *database.DB) {
	gp := app.Group("/missing")

	gp.Post("/", func(c fiber.Ctx) error {
		return createMissing(c, db)
	})
}

func createMissing(c fiber.Ctx, db *database.DB) error {
	return nil
}
