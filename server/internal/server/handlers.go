package server

import (
	findAnimalHTTP "github.com/francoggm/save-rs-brazil/internal/find_animal/http"
	rescueHTTP "github.com/francoggm/save-rs-brazil/internal/rescue/http"
)

func (s *server) ConfigHandlers() {
	v1 := s.app.Group("/v1")

	rescue := v1.Group("/rescue")
	findAnimal := v1.Group("/find_animal")

	rescueHTTP.ConfigRoutes(rescue, s.db)
	findAnimalHTTP.ConfigRoutes(findAnimal, s.db)
}
