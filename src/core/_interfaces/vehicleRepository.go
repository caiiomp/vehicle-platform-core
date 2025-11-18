package interfaces

import (
	"context"

	"github.com/caiiomp/vehicle-platform-core/src/core/domain/entity"
)

type VehicleRepository interface {
	Create(ctx context.Context, vehicle entity.Vehicle) (*entity.Vehicle, error)
	Update(ctx context.Context, id string, vehicle entity.Vehicle) (*entity.Vehicle, error)
}
