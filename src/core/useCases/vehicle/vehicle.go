package vehicle

import (
	"context"

	interfaces "github.com/caiiomp/vehicle-platform-core/src/core/_interfaces"
	"github.com/caiiomp/vehicle-platform-core/src/core/domain/entity"
)

type vehicleService struct {
	vehicleRepository           interfaces.VehicleRepository
	vehiclePlatformSalesAdapter interfaces.VehiclePlatformSalesAdapter
}

func NewVehicleService(
	vehicleRepository interfaces.VehicleRepository,
	vehiclePlatformSalesAdapter interfaces.VehiclePlatformSalesAdapter,
) interfaces.VehicleService {
	return &vehicleService{
		vehicleRepository:           vehicleRepository,
		vehiclePlatformSalesAdapter: vehiclePlatformSalesAdapter,
	}
}

func (ref *vehicleService) Create(ctx context.Context, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	created, err := ref.vehicleRepository.Create(ctx, vehicle)
	if err != nil {
		return nil, err
	}

	if created == nil {
		return nil, nil
	}

	err = ref.vehiclePlatformSalesAdapter.CreateVehicle(ctx, created.EntityID.String(), created.Brand, created.Model, created.Color, created.Year, created.Price)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (ref *vehicleService) Update(ctx context.Context, id string, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	updated, err := ref.vehicleRepository.Update(ctx, id, vehicle)
	if err != nil {
		return nil, err
	}

	if updated == nil {
		return nil, nil
	}

	err = ref.vehiclePlatformSalesAdapter.UpdateVehicle(ctx, updated.EntityID.String(), updated.Brand, updated.Model, updated.Color, updated.Year, updated.Price)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
