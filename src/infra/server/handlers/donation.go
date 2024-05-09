package handlers

import (
	"github.com/francoggm/save-rs-brazil/infra/database"
	"github.com/gofiber/fiber/v3"
)

func DonationHandlers(app *fiber.App, db *database.DB) {
	gp := app.Group("/donation")

	gp.Post("/", func(c fiber.Ctx) error {
		return createDonation(c, db)
	})
}

func createDonation(c fiber.Ctx, db *database.DB) error {
	return nil
}
