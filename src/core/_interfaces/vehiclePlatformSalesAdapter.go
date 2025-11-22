package interfaces

import (
	"context"
)

type VehiclePlatformSalesAdapter interface {
	CreateVehicle(ctx context.Context, entityID, brand, model, color string, year int, price float64) error
	UpdateVehicle(ctx context.Context, entityID, brand, model, color string, year int, price float64) error
}
