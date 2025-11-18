package vehicle

import (
	"context"
	"errors"
	"testing"
	"time"

	mocks "github.com/caiiomp/vehicle-platform-core/src/core/_mocks"
	"github.com/caiiomp/vehicle-platform-core/src/core/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ctx := context.TODO()
	expectedError := errors.New("some error")

	t.Run("should not create vehicle when failed to create", func(t *testing.T) {
		vehicle := entity.Vehicle{
			Brand: "Chevrolet",
			Model: "Tracker",
			Year:  2023,
			Color: "Black",
			Price: 100000,
		}

		vehicleRepositoryMocked := mocks.NewVehicleRepository(t)

		vehicleRepositoryMocked.On("Create", ctx, vehicle).
			Return(nil, expectedError)

		service := NewVehicleService(vehicleRepositoryMocked)

		actual, err := service.Create(ctx, vehicle)

		assert.Nil(t, actual)
		assert.Equal(t, expectedError, err)
	})

	t.Run("should not create vehicle when vehicle not found after created", func(t *testing.T) {
		vehicle := entity.Vehicle{
			Brand: "Chevrolet",
			Model: "Tracker",
			Year:  2023,
			Color: "Black",
			Price: 100000,
		}

		vehicleRepositoryMocked := mocks.NewVehicleRepository(t)

		vehicleRepositoryMocked.On("Create", ctx, vehicle).
			Return(nil, nil)

		service := NewVehicleService(vehicleRepositoryMocked)

		actual, err := service.Create(ctx, vehicle)

		assert.Nil(t, actual)
		assert.Nil(t, err)
	})

	t.Run("should create vehicle successfully", func(t *testing.T) {
		vehicle := entity.Vehicle{
			Brand: "Chevrolet",
			Model: "Tracker",
			Year:  2023,
			Color: "Black",
			Price: 100000,
		}

		vehicleID := uuid.NewString()
		now := time.Now()

		created := vehicle
		created.ID = vehicleID
		created.CreatedAt = now
		created.UpdatedAt = now

		vehicleRepositoryMocked := mocks.NewVehicleRepository(t)

		vehicleRepositoryMocked.On("Create", ctx, vehicle).
			Return(&created, nil)

		service := NewVehicleService(vehicleRepositoryMocked)

		expected := created

		actual, err := service.Create(ctx, vehicle)

		assert.Equal(t, &expected, actual)
		assert.Nil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	ctx := context.TODO()
	expectedError := errors.New("some error")

	t.Run("should not update vehicle when failed to update", func(t *testing.T) {
		vehicleID := uuid.NewString()

		vehicle := entity.Vehicle{
			Brand: "Chevrolet",
			Model: "Tracker",
			Year:  2023,
			Color: "Black",
			Price: 100000,
		}

		vehicleRepositoryMocked := mocks.NewVehicleRepository(t)

		vehicleRepositoryMocked.On("Update", ctx, vehicleID, vehicle).
			Return(nil, expectedError)

		service := NewVehicleService(vehicleRepositoryMocked)

		actual, err := service.Update(ctx, vehicleID, vehicle)

		assert.Nil(t, actual)
		assert.Equal(t, expectedError, err)
	})

	t.Run("should not update vehicle when vehicle not found after update", func(t *testing.T) {
		vehicleID := uuid.NewString()

		vehicle := entity.Vehicle{
			Brand: "Chevrolet",
			Model: "Tracker",
			Year:  2023,
			Color: "Black",
			Price: 100000,
		}

		vehicleRepositoryMocked := mocks.NewVehicleRepository(t)

		vehicleRepositoryMocked.On("Update", ctx, vehicleID, vehicle).
			Return(nil, nil)

		service := NewVehicleService(vehicleRepositoryMocked)

		actual, err := service.Update(ctx, vehicleID, vehicle)

		assert.Nil(t, actual)
		assert.Nil(t, err)
	})

	t.Run("should update vehicle successfully", func(t *testing.T) {
		vehicleID := uuid.NewString()

		vehicle := entity.Vehicle{
			Brand: "Chevrolet",
			Model: "Tracker",
			Year:  2023,
			Color: "Black",
			Price: 100000,
		}

		now := time.Now()

		updated := vehicle
		updated.ID = vehicleID
		updated.CreatedAt = now
		updated.UpdatedAt = now

		vehicleRepositoryMocked := mocks.NewVehicleRepository(t)

		vehicleRepositoryMocked.On("Update", ctx, vehicleID, vehicle).
			Return(&updated, nil)

		service := NewVehicleService(vehicleRepositoryMocked)

		expected := updated

		actual, err := service.Update(ctx, vehicleID, vehicle)

		assert.Equal(t, &expected, actual)
		assert.Nil(t, err)
	})
}
