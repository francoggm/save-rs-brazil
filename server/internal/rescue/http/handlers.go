package http

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/francoggm/save-rs-brazil/internal/models"
	"github.com/francoggm/save-rs-brazil/internal/rescue/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

func createRescue(c fiber.Ctx, db *sqlx.DB) error {
	rescue := new(models.Rescue)

	if err := c.Bind().Body(rescue); err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(map[string]any{
			"error": "invalid fields!",
		})
	}

	validate := validator.New()
	if err := validate.Struct(rescue); err != nil {
		fmt.Println(err)
		return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
			"error": "invalid required fields!",
		})
	}

	if err := repository.CreateRescue(db, rescue); err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(map[string]any{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(rescue)
}

func getRescues(c fiber.Ctx, db *sqlx.DB) error {
	rescues, err := repository.GetRescues(db)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]any{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(rescues)
}

func getRescue(c fiber.Ctx, db *sqlx.DB) error {
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
