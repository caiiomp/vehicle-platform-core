package vehiclerepository

import (
	"context"
	"time"

	interfaces "github.com/caiiomp/vehicle-platform-core/src/core/_interfaces"
	"github.com/caiiomp/vehicle-platform-core/src/core/domain/entity"
	"github.com/caiiomp/vehicle-platform-core/src/repositories/model"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type vehicleRepository struct {
	collection *mongo.Collection
}

func NewVehicleRepository(collection *mongo.Collection) interfaces.VehicleRepository {
	return &vehicleRepository{
		collection: collection,
	}
}

func (ref *vehicleRepository) Create(ctx context.Context, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	record := model.VehicleFromDomain(vehicle)

	record.EntityID = uuid.NewString()

	now := time.Now()
	record.CreatedAt = now
	record.UpdatedAt = now

	created, err := ref.collection.InsertOne(ctx, record)
	if err != nil {
		return nil, err
	}

	id := created.InsertedID.(primitive.ObjectID)

	result := ref.collection.FindOne(ctx, bson.M{"_id": id})
	if err = result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var recordToReturn model.Vehicle
	if err = result.Decode(&recordToReturn); err != nil {
		return nil, err
	}

	return recordToReturn.ToDomain(), nil
}

func (ref *vehicleRepository) Update(ctx context.Context, id string, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	record := model.VehicleFromDomain(vehicle)
	record.UpdatedAt = time.Now()

	update := bson.M{
		"$set": record,
	}

	_, err := ref.collection.UpdateOne(ctx, bson.M{"entity_id": id}, update)
	if err != nil {
		return nil, err
	}

	result := ref.collection.FindOne(ctx, bson.M{"entity_id": id})
	if err = result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var recordToReturn model.Vehicle
	if err = result.Decode(&recordToReturn); err != nil {
		return nil, err
	}

	return recordToReturn.ToDomain(), nil
}
