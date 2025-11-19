package interfaces

import (
	"context"
)

type VehiclePlatformSalesAdapter interface {
	CreateVehicle(ctx context.Context, vehicleID, brand, model, color string, year int, price float64) error
}
