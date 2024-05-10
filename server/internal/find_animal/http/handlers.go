package http

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/francoggm/save-rs-brazil/internal/find_animal/repository"
	"github.com/francoggm/save-rs-brazil/internal/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

func createFindAnimal(c fiber.Ctx, db *sqlx.DB) error {
	findAnimal := new(models.FindAnimal)

	if err := c.Bind().Body(findAnimal); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]any{
			"error": "invalid fields!",
		})
	}

	validate := validator.New()
	if err := validate.Struct(findAnimal); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
			"error": "invalid required fields!",
		})
	}

	if err := repository.CreateFindAnimal(db, findAnimal); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]any{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(findAnimal)
}

func getFindAnimals(c fiber.Ctx, db *sqlx.DB) error {
	findAnimals, err := repository.GetFindAnimals(db)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]any{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(findAnimals)
}

func getFindAnimal(c fiber.Ctx, db *sqlx.DB) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]any{
			"error": "invalid id!",
		})
	}

	findAnimal, err := repository.GetFindAnimal(db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(http.StatusNotFound).JSON(map[string]any{
				"error": "find animal not found!",
			})
		}

		return c.Status(http.StatusBadRequest).JSON(map[string]any{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(findAnimal)
}
