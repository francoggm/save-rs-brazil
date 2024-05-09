package models

type (
	UrgencyType int8
	RescueType  int8
	GenderType  int8
)

const (
	Low UrgencyType = iota
	Medium
	Critical
)

const (
	Person RescueType = iota
	Animal
)

const (
	Unknown GenderType = iota
	Male
	Female
)

type Address struct {
	Street       string `json:"street,omitempty" db:"street"`
	Number       int    `json:"number,omitempty" db:"number"`
	Neighborhood string `json:"neighborhood" db:"neighborhood" validate:"required"`
	ZIPCode      string `json:"zip_code,omitempty" db:"zip_code"`
	Complement   string `json:"complement,omitempty" db:"complement"`
}
