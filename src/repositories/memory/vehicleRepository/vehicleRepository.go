package vehiclerepository

import (
	"context"
	"time"

	"github.com/google/uuid"

	interfaces "github.com/caiiomp/vehicle-platform-core/src/core/_interfaces"
	"github.com/caiiomp/vehicle-platform-core/src/core/domain/entity"
	"github.com/caiiomp/vehicle-platform-core/src/repositories/model"
)

type vehicleRepository struct {
	vehicles []model.Vehicle
}

func NewVehicleRepository() interfaces.VehicleRepository {
	return &vehicleRepository{
		vehicles: []model.Vehicle{},
	}
}

func (ref *vehicleRepository) Create(ctx context.Context, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	record := model.VehicleFromDomain(vehicle)

	record.ID = uuid.NewString()

	now := time.Now()
	record.CreatedAt = now
	record.UpdatedAt = now

	ref.vehicles = append(ref.vehicles, record)

	for _, vehicle := range ref.vehicles {
		if vehicle.ID == record.ID {
			return vehicle.ToDomain(), nil
		}
	}

	return nil, nil
}

func (ref *vehicleRepository) Update(ctx context.Context, id string, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	vehicleIndex := -1

	for i, vehicle := range ref.vehicles {
		if vehicle.ID == id {
			vehicleIndex = i
			break
		}
	}

	if vehicleIndex == -1 {
		return nil, nil
	}

	var hasUpdate bool

	if vehicle.Brand != "" && vehicle.Brand != ref.vehicles[vehicleIndex].Brand {
		ref.vehicles[vehicleIndex].Brand = vehicle.Brand
		hasUpdate = true
	}

	if vehicle.Model != "" && vehicle.Model != ref.vehicles[vehicleIndex].Model {
		ref.vehicles[vehicleIndex].Model = vehicle.Model
		hasUpdate = true
	}

	if vehicle.Year != 0 && vehicle.Year != ref.vehicles[vehicleIndex].Year {
		ref.vehicles[vehicleIndex].Year = vehicle.Year
		hasUpdate = true
	}

	if vehicle.Color != "" && vehicle.Color != ref.vehicles[vehicleIndex].Color {
		ref.vehicles[vehicleIndex].Color = vehicle.Color
		hasUpdate = true
	}

	if vehicle.Price != 0 && vehicle.Price != ref.vehicles[vehicleIndex].Price {
		ref.vehicles[vehicleIndex].Price = vehicle.Price
		hasUpdate = true
	}

	if hasUpdate {
		ref.vehicles[vehicleIndex].UpdatedAt = time.Now()
	}

	return ref.vehicles[vehicleIndex].ToDomain(), nil
}
