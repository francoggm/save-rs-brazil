package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

func ConfigRoutes(router fiber.Router, db *sqlx.DB) {
	router.Post("/", func(c fiber.Ctx) error {
		return createFindAnimal(c, db)
	})

	router.Get("/", func(c fiber.Ctx) error {
		return getFindAnimals(c, db)
	})

	router.Get("/:id", func(c fiber.Ctx) error {
		return getFindAnimal(c, db)
	})
}
