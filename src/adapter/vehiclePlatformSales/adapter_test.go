package vehicleplatformsales

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	mocks "github.com/caiiomp/vehicle-platform-core/src/core/_mocks"
)

func TestCreateVehicle(t *testing.T) {
	ctx := context.TODO()
	vehicleID := uuid.NewString()
	brand := "Fiat"
	model := "Argo"
	year := 2023
	color := "Cinza"
	price := float64(80000)

	httpClientMocked := mocks.NewVehiclePlatformSalesHttpClient(t)

	httpClientMocked.On("CreateVehicle", ctx, vehicleID, brand, model, color, year, price).Return(nil)

	adapter := NewVehiclePlatformSalesAdapter(httpClientMocked)

	err := adapter.CreateVehicle(ctx, vehicleID, brand, model, color, year, price)

	assert.Nil(t, err)
}

func TestUpdateVehicle(t *testing.T) {
	ctx := context.TODO()
	vehicleID := uuid.NewString()
	brand := "Fiat"
	model := "Argo"
	year := 2023
	color := "Cinza"
	price := float64(80000)

	httpClientMocked := mocks.NewVehiclePlatformSalesHttpClient(t)

	httpClientMocked.On("UpdateVehicle", ctx, vehicleID, brand, model, color, year, price).Return(nil)

	adapter := NewVehiclePlatformSalesAdapter(httpClientMocked)

	err := adapter.UpdateVehicle(ctx, vehicleID, brand, model, color, year, price)

	assert.Nil(t, err)
}
