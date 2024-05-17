package server

import (
	animalsHTTP "github.com/francoggm/save-rs-brazil/internal/animal/http"
)

func (s *server) ConfigHandlers() {
	v1 := s.app.Group("/v1")

	animalsGroup := v1.Group("/animals")
	animalsHTTP.ConfigRoutes(animalsGroup, s.db)
}
