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

func TestCreateRescue(t *testing.T) {
	t.Parallel()

	cfg := config.New()

	db, err := postgres.New(cfg)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	app := fiber.New()
	app.Post("/v1/rescue", func(c fiber.Ctx) error {
		return createRescue(c, db)
	})

	body := strings.NewReader(`{"rescue_count":1,"rescue_type":1,"urgency":1,"description":"Test-description","address":{"street":"Test-street","number":123,"neighborhood":"Test-neighborhood","zip_code":"Test-ZIPCode","complement":"Test-Complement"}}`)

	req := httptest.NewRequest("POST", "/v1/rescue", body)
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req, -1)

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode)

	var r models.Rescue
	b, err := io.ReadAll(resp.Body)
	utils.AssertEqual(t, nil, err)

	err = json.Unmarshal(b, &r)
	utils.AssertEqual(t, nil, err)
}
