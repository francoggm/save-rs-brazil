package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/francoggm/save-rs-brazil/config"
	"github.com/francoggm/save-rs-brazil/internal/models"
	"github.com/francoggm/save-rs-brazil/pkg/database/postgres"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/utils"
)

func TestCreateFindAnimal(t *testing.T) {
	t.Parallel()

	cfg := config.New()

	db, err := postgres.New(cfg)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	app := fiber.New()
	app.Post("/v1/find_animal", func(c fiber.Ctx) error {
		return createFindAnimal(c, db)
	})

	body := strings.NewReader(`{"specie":"gato","race":"laranja","gender":1,"description":"Test-description","address":{"street":"Test-street","number":123,"neighborhood":"Test-neighborhood","zip_code":"Test-ZIPCode","complement":"Test-Complement"}}`)

	req := httptest.NewRequest("POST", "/v1/find_animal", body)
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req, -1)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode)

	var fa models.FindAnimal
	b, err := io.ReadAll(resp.Body)
	utils.AssertEqual(t, nil, err)

	err = json.Unmarshal(b, &fa)
	utils.AssertEqual(t, nil, err)
}
