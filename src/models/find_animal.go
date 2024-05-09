package models

type FindAnimal struct {
	Photos      []string   `json:"photos,omitempty"`
	Specie      string     `json:"specie"`
	Race        string     `json:"race,omitempty"`
	Gender      GenderType `json:"gender"`
	Description string     `json:"description,omitempty"`
	Addr        Address    `json:"address"`
}
