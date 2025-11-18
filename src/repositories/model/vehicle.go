package model

import (
	"time"

	"github.com/caiiomp/vehicle-platform-core/src/core/domain/entity"
)

type Vehicle struct {
	ID        string    `db:"id"`
	Brand     string    `db:"brand"`
	Model     string    `db:"model"`
	Year      int32     `db:"year"`
	Color     string    `db:"color"`
	Price     float64   `db:"price"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func VehicleFromDomain(vehicle entity.Vehicle) Vehicle {
	return Vehicle{
		ID:        vehicle.ID,
		Brand:     vehicle.Brand,
		Model:     vehicle.Model,
		Year:      vehicle.Year,
		Color:     vehicle.Color,
		Price:     vehicle.Price,
		CreatedAt: vehicle.CreatedAt,
		UpdatedAt: vehicle.UpdatedAt,
	}
}

func (ref *Vehicle) ToDomain() *entity.Vehicle {
	return &entity.Vehicle{
		ID:        ref.ID,
		Brand:     ref.Brand,
		Model:     ref.Model,
		Year:      ref.Year,
		Color:     ref.Color,
		Price:     ref.Price,
		CreatedAt: ref.CreatedAt,
		UpdatedAt: ref.UpdatedAt,
	}
}
