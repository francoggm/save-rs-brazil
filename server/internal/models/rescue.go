package models

import "time"

type Rescue struct {
	ID          int64       `json:"id" db:"id"`
	RescueCount int         `json:"rescue_count,omitempty" db:"rescue_count"`
	RescueType  RescueType  `json:"rescue_type" db:"rescue_type" validate:"required,min=1,max=2"`
	Urgency     UrgencyType `json:"urgency" db:"urgency" validate:"required,min=1,max=3"`
	Description string      `json:"description,omitempty" db:"description"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
	Address     `json:"address"`
}
