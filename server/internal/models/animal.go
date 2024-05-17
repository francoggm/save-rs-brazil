package models

import "time"

type Animal struct {
	ID          int64  `json:"id" db:"id"`
	Specie      string `json:"specie" db:"specie" validate:"required"`
	Race        string `json:"race,omitempty" db:"race"`
	Gender      Gender `json:"gender" db:"gender" validate:"required,min=1,max=3"`
	State       State  `json:"state" db:"state" validate:"required,min=1,max=3"`
	Description string `json:"description,omitempty" db:"description"`
	Address     `json:"address"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
