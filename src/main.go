package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	vehicleplatformsales "github.com/caiiomp/vehicle-platform-core/src/adapter/vehiclePlatformSales"
	vehiclePlatformSalesHttp "github.com/caiiomp/vehicle-platform-core/src/adapter/vehiclePlatformSales/http"
	"github.com/caiiomp/vehicle-platform-core/src/core/useCases/vehicle"
	_ "github.com/caiiomp/vehicle-platform-core/src/docs"
	"github.com/caiiomp/vehicle-platform-core/src/presentation"
	"github.com/caiiomp/vehicle-platform-core/src/presentation/vehicleApi"
	vehiclerepository "github.com/caiiomp/vehicle-platform-core/src/repositories/mongodb/vehicleRepository"
)

func main() {
	var (
		apiPort = os.Getenv("API_PORT")

		mongoURI      = os.Getenv("MONGO_URI")
		mongoDatabase = os.Getenv("MONGO_DATABASE")

		vehiclePlatformSalesHost = os.Getenv("VEHICLE_PLATFORM_SALES_HOST")
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)

	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("could not initialize mongodb client: %v", err)
	}

	if err = mongoClient.Ping(ctx, nil); err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// Collections
	vehiclesCollection := mongoClient.Database(mongoDatabase).Collection("vehicles")

	// Repositories
	vehicleRepository := vehiclerepository.NewVehicleRepository(vehiclesCollection)

	// Http Clients
	httpClient := &http.Client{
		Timeout: time.Second * 3,
	}
	vehiclePlatformSalesHttpClient := vehiclePlatformSalesHttp.NewVehiclePlatformSalesHttpClient(httpClient, vehiclePlatformSalesHost)

	// Adapters
	vehiclePlatformSalesAdapter := vehicleplatformsales.NewVehiclePlatformSalesAdapter(vehiclePlatformSalesHttpClient)

	// Services
	vehicleService := vehicle.NewVehicleService(vehicleRepository, vehiclePlatformSalesAdapter)

	app := presentation.SetupServer()

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	vehicleApi.RegisterVehicleRoutes(app, vehicleService)

	if apiPort == "" {
		apiPort = "8080"
	}

	if err = app.Run(":" + apiPort); err != nil {
		log.Fatalf("coult not initialize http server: %v", err)
	}
}
