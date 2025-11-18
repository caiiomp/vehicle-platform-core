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

	_, err := ref.db.ExecContext(ctx, insertVehicle, record.ID, record.Brand, record.Model, record.Year, record.Color, record.Price)
	if err != nil {
		return nil, err
	}

	row := ref.db.QueryRowContext(ctx, getVehicleByID, record.ID)

	var created model.Vehicle
	if err = row.Scan(&created.ID, &created.Brand, &created.Model, &created.Year, &created.Color, &created.Price, &created.CreatedAt, &created.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return created.ToDomain(), nil
}

func (ref *vehicleRepository) Update(ctx context.Context, id string, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	row := ref.db.QueryRowContext(ctx, getVehicleByID, id)

	var current model.Vehicle
	if err := row.Scan(&current.ID, &current.Brand, &current.Model, &current.Year, &current.Color, &current.Price, &current.CreatedAt, &current.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var hasUpdate bool

	if vehicle.Brand != "" {
		current.Brand = vehicle.Brand
		hasUpdate = true
	}

	if vehicle.Model != "" {
		current.Model = vehicle.Model
		hasUpdate = true
	}

	if vehicle.Color != "" {
		current.Color = vehicle.Color
		hasUpdate = true
	}

	if vehicle.Year != 0 {
		current.Year = vehicle.Year
		hasUpdate = true
	}

	if vehicle.Price != 0 {
		current.Price = vehicle.Price
		hasUpdate = true
	}

	if !hasUpdate {
		return nil, nil
	}

	_, err := ref.db.ExecContext(ctx, updateVehicle, id, current.Brand, current.Model, current.Year, current.Color, current.Price)
	if err != nil {
		return nil, err
	}

	row = ref.db.QueryRowContext(ctx, getVehicleByID, id)

	var updated model.Vehicle
	if err = row.Scan(&updated.ID, &updated.Brand, &updated.Model, &updated.Year, &updated.Color, &updated.Price, &updated.CreatedAt, &updated.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return updated.ToDomain(), nil
}
