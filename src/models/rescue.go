package models

import "time"

type Rescue struct {
	ID          int64       `json:"id,omitempty" db:"id"`
	Photos      []string    `json:"photos,omitempty"`
	RescueCount int         `json:"rescue_count,omitempty" db:"rescue_count"`
	RescueType  RescueType  `json:"rescue_type" db:"rescue_type" validate:"required"`
	Urgency     UrgencyType `json:"urgency" db:"urgency" validate:"required"`
	Description string      `json:"description,omitempty" db:"description"`
	CreatedAt   time.Time   `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty" db:"updated_at"`
	Address     `json:"address"`
}
