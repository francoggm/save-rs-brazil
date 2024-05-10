package middlewares

import (
	"github.com/francoggm/save-rs-brazil/config"
	"github.com/gofiber/fiber/v3"
)

// TODO: Create custom logger and use as param
func ConfigMiddlewares(cfg *config.Config, app *fiber.App) {
	app.Use(func(c fiber.Ctx) error {
		return logMiddleware(cfg, c)
	})
}
