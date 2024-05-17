package http

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/francoggm/save-rs-brazil/internal/animal/repository"
	"github.com/francoggm/save-rs-brazil/internal/models"
	"github.com/francoggm/save-rs-brazil/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

func createAnimal(c fiber.Ctx, db *sqlx.DB) error {
	animal := new(models.Animal)

	if err := c.Bind().Body(animal); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.GetErrorMessage("invalid fields"))
	}

	validate := validator.New()
	if err := validate.Struct(animal); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.GetErrorMessage("invalid required fields"))
	}

	if err := repository.CreateAnimal(db, animal); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.GetErrorMessage(err.Error()))
	}

	return c.Status(http.StatusOK).JSON(animal)
}

func getAnimals(c fiber.Ctx, db *sqlx.DB) error {
	state := c.Queries()["state"]

	animals, err := repository.GetAnimals(state, db)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.GetErrorMessage(err.Error()))
	}

	return c.Status(http.StatusOK).JSON(animals)
}

func getAnimal(c fiber.Ctx, db *sqlx.DB) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]any{
			"error": "invalid id",
		})
	}

	state := c.Queries()["state"]

	animal, err := repository.GetAnimal(id, state, db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(http.StatusNotFound).JSON(utils.GetErrorMessage("animal not found"))
		}

		return c.Status(http.StatusBadRequest).JSON(utils.GetErrorMessage(err.Error()))
	}

	return c.Status(http.StatusOK).JSON(animal)
}
