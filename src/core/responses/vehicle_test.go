package responses

import (
	"testing"
	"time"

	"github.com/caiiomp/vehicle-platform-core/src/core/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestVehicleFromDomain(t *testing.T) {
	vehicleID := uuid.NewString()

	now := time.Now()

	vehicle := entity.Vehicle{
		ID:        vehicleID,
		Brand:     "Some Brand",
		Model:     "Some Model",
		Year:      2025,
		Color:     "Gray",
		Price:     80000,
		CreatedAt: now,
		UpdatedAt: now,
	}

	expected := Vehicle{
		ID:        vehicleID,
		Brand:     "Some Brand",
		Model:     "Some Model",
		Year:      2025,
		Color:     "Gray",
		Price:     80000,
		CreatedAt: now,
		UpdatedAt: now,
	}

	actual := VehicleFromDomain(vehicle)

	assert.Equal(t, expected, actual)
}
