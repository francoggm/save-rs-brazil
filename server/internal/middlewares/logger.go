package middlewares

import (
	"log"
	"time"

	"github.com/francoggm/save-rs-brazil/config"
	"github.com/gofiber/fiber/v3"
)

func logMiddleware(cfg *config.Config, c fiber.Ctx) error {
	start := time.Now()
	err := c.Next()

	req := c.Request()
	res := c.Response()

	time := time.Since(start)

	log.Printf("Method: %v, URI: %v, Status: %v, Time: %v", c.Method(), req.URI(), res.StatusCode(), time)

	if cfg.Mode == "debug" {
		reqHeaders := string(req.Header.Header())
		reqBody := string(req.Body())

		resHeaders := string(res.Header.Header())
		resBody := string(res.Body())

		log.Printf("DEBUG | Request -> Headers: %v, Body: %v | Response -> Headers: %v, Body: %v", reqHeaders, reqBody, resHeaders, resBody)
	}

	return err
}
