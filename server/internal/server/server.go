package server

import (
	"github.com/francoggm/save-rs-brazil/config"
	"github.com/francoggm/save-rs-brazil/internal/middlewares"
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type server struct {
	cfg *config.Config
	app *fiber.App
	db  *sqlx.DB
}

func New(cfg *config.Config, db *sqlx.DB) *server {
	return &server{
		cfg: cfg,
		app: fiber.New(),
		db:  db,
	}
}

func (s *server) Run() error {
	middlewares.ConfigMiddlewares(s.cfg, s.app)
	s.ConfigHandlers()

	return s.app.Listen("0.0.0.0:" + s.cfg.Port)
}
