package entity

import "time"

type Vehicle struct {
	ID        string
	Brand     string
	Model     string
	Year      int32
	Color     string
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
