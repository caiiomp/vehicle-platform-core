package vehicleplatformsales

import (
	"context"

	"github.com/caiiomp/vehicle-platform-core/src/adapter/vehiclePlatformSales/http"
	interfaces "github.com/caiiomp/vehicle-platform-core/src/core/_interfaces"
)

type vehiclePlatformSalesAdapter struct {
	httpClient http.VehiclePlatformSalesHttpClient
}

func NewVehiclePlatformSalesAdapter(httpClient http.VehiclePlatformSalesHttpClient) interfaces.VehiclePlatformSalesAdapter {
	return &vehiclePlatformSalesAdapter{
		httpClient: httpClient,
	}
}

func (ref *vehiclePlatformSalesAdapter) CreateVehicle(ctx context.Context, vehicleID, brand, model, color string, year int, price float64) error {
	return ref.httpClient.CreateVehicle(ctx, vehicleID, brand, model, color, year, price)
}

func (ref *vehiclePlatformSalesAdapter) UpdateVehicle(ctx context.Context, vehicleID string, brand string, model string, color string, year int, price float64) error {
	return ref.httpClient.UpdateVehicle(ctx, vehicleID, brand, model, color, year, price)
}
