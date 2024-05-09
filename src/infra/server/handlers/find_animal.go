package handlers

import (
	"github.com/francoggm/save-rs-brazil/infra/database"
	"github.com/gofiber/fiber/v3"
)

func FindAnimalHandlers(app *fiber.App, db *database.DB) {
	gp := app.Group("/animal")

	gp.Post("/", func(c fiber.Ctx) error {
		return createFindAnimal(c, db)
	})
}

func createFindAnimal(c fiber.Ctx, db *database.DB) error {
	return nil
}
