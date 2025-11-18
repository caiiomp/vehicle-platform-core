package vehicle

import (
	"context"

	interfaces "github.com/caiiomp/vehicle-platform-core/src/core/_interfaces"
	"github.com/caiiomp/vehicle-platform-core/src/core/domain/entity"
)

type vehicleService struct {
	vehicleRepository interfaces.VehicleRepository
}

func NewVehicleService(vehicleRepository interfaces.VehicleRepository) interfaces.VehicleService {
	return &vehicleService{
		vehicleRepository: vehicleRepository,
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

	return updated, nil
}
