package model

import (
	"time"

	"github.com/google/uuid"

	"github.com/caiiomp/vehicle-platform-core/src/core/domain/entity"
)

type Vehicle struct {
	ID        string    `bson:"_id,omitempty"`
	EntityID  string    `bson:"entity_id,omitempty"`
	Brand     string    `bson:"brand,omitempty"`
	Model     string    `bson:"model,omitempty"`
	Year      int       `bson:"year,omitempty"`
	Color     string    `bson:"color,omitempty"`
	Price     float64   `bson:"price,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
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
		EntityID:  uuid.MustParse(ref.EntityID),
		Brand:     ref.Brand,
		Model:     ref.Model,
		Year:      ref.Year,
		Color:     ref.Color,
		Price:     ref.Price,
		CreatedAt: ref.CreatedAt,
		UpdatedAt: ref.UpdatedAt,
	}
}
