package handlers

import (
	"github.com/francoggm/save-rs-brazil/infra/database"
	"github.com/gofiber/fiber/v3"
)

func CleaningHandlers(app *fiber.App, db *database.DB) {
	gp := app.Group("/cleaning")

	gp.Post("/", func(c fiber.Ctx) error {
		return createCleaning(c, db)
	})
}

func createCleaning(c fiber.Ctx, db *database.DB) error {
	return nil
}
