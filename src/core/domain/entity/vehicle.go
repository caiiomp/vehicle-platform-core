package entity

import (
	"time"

	"github.com/google/uuid"
)

type Vehicle struct {
	ID        string
	EntityID  uuid.UUID
	Brand     string
	Model     string
	Year      int
	Color     string
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
