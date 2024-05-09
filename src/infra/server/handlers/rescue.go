package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/francoggm/save-rs-brazil/infra/database"
	"github.com/francoggm/save-rs-brazil/infra/database/repository"
	"github.com/francoggm/save-rs-brazil/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func RescueHandlers(app *fiber.App, db *database.DB) {
	gp := app.Group("/rescue")

	gp.Post("/", func(c fiber.Ctx) error {
		return createRescue(c, db)
	})

	gp.Get("/", func(c fiber.Ctx) error {
		return getRescues(c, db)
	})

	gp.Get("/:id", func(c fiber.Ctx) error {
		return getRescue(c, db)
	})
}

func createRescue(c fiber.Ctx, db *database.DB) error {
	rescue := new(models.Rescue)

	if err := c.Bind().Body(rescue); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]any{
			"error": "invalid fields!",
		})
	}

	validate := validator.New()
	if err := validate.Struct(rescue); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
			"error": "missing required fields!",
		})
	}

	if err := repository.CreateRescue(db, rescue); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]any{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(rescue)
}

func getRescues(c fiber.Ctx, db *database.DB) error {
	rescues, err := repository.GetRescues(db)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]any{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(rescues)
}

func getRescue(c fiber.Ctx, db *database.DB) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]any{
			"error": "invalid id!",
		})
	}

	rescue, err := repository.GetRescue(db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(http.StatusNotFound).JSON(map[string]any{
				"error": "rescue not found!",
			})
		}

		return c.Status(http.StatusBadRequest).JSON(map[string]any{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(rescue)
}
