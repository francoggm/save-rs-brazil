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

func TestCreateAnimal(t *testing.T) {
	t.Parallel()

	cfg := config.New()

	db, err := postgres.New(cfg)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	app := fiber.New()
	app.Post("/v1/animals", func(c fiber.Ctx) error {
		return createAnimal(c, db)
	})

	body := strings.NewReader(`{"specie":"gato","race":"laranja","gender":1,"state":1,"description":"Test-description","address":{"street":"Test-street","number":123,"neighborhood":"Test-neighborhood","zip_code":"Test-ZIPCode","complement":"Test-Complement","phone": "980808000"}}`)

	req := httptest.NewRequest("POST", "/v1/animals", body)
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		log.Println(err)
	}

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode)

	var an models.Animal
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	utils.AssertEqual(t, nil, err)

	err = json.Unmarshal(b, &an)
	if err != nil {
		log.Println(err)
	}

	utils.AssertEqual(t, nil, err)
}
