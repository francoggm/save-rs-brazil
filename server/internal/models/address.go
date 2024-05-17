package models

type Address struct {
	Street       string `json:"street,omitempty" db:"street"`
	Number       int    `json:"number,omitempty" db:"number"`
	Neighborhood string `json:"neighborhood" db:"neighborhood" validate:"required"`
	ZIPCode      string `json:"zip_code,omitempty" db:"zip_code"`
	Complement   string `json:"complement,omitempty" db:"complement"`
	Phone        string `json:"phone,omitempty" db:"phone"`
}
