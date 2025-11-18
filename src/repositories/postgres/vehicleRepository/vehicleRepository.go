package vehiclerepository

import (
	"context"
	"database/sql"

	interfaces "github.com/caiiomp/vehicle-platform-core/src/core/_interfaces"
	"github.com/caiiomp/vehicle-platform-core/src/core/domain/entity"
	"github.com/caiiomp/vehicle-platform-core/src/repositories/model"
	"github.com/google/uuid"
)

type vehicleRepository struct {
	db *sql.DB
}

func NewVehicleRepository(db *sql.DB) interfaces.VehicleRepository {
	return &vehicleRepository{
		db: db,
	}
}

func (ref *vehicleRepository) Create(ctx context.Context, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	record := model.VehicleFromDomain(vehicle)
	record.ID = uuid.NewString()

	_, err := ref.db.ExecContext(ctx, insertVehicle, record)
	if err != nil {
		return nil, err
	}

	row := ref.db.QueryRowContext(ctx, getVehicleByID, record.ID)

	var created model.Vehicle
	if err = row.Scan(&created); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return created.ToDomain(), nil
}

func (ref *vehicleRepository) Update(ctx context.Context, id string, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	record := model.VehicleFromDomain(vehicle)
	record.ID = id

	_, err := ref.db.ExecContext(ctx, updateVehicle, record)
	if err != nil {
		return nil, err
	}

	row := ref.db.QueryRowContext(ctx, getVehicleByID, record.ID)

	var updated model.Vehicle
	if err = row.Scan(&updated); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return updated.ToDomain(), nil
}
