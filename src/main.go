package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/caiiomp/vehicle-platform-core/src/core/useCases/vehicle"
	_ "github.com/caiiomp/vehicle-platform-core/src/docs"
	"github.com/caiiomp/vehicle-platform-core/src/presentation"
	"github.com/caiiomp/vehicle-platform-core/src/presentation/vehicleApi"
	vehiclerepository "github.com/caiiomp/vehicle-platform-core/src/repositories/postgres/vehicleRepository"
)

func main() {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("error to connect database: %s", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("error to ping database: %s", err)
	}

	// Repositories
	vehicleRepository := vehiclerepository.NewVehicleRepository(db)

	// Services
	vehicleService := vehicle.NewVehicleService(vehicleRepository)

	app := presentation.SetupServer()

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	vehicleApi.RegisterVehicleRoutes(app, vehicleService)

	if err = app.Run(":4001"); err != nil {
		log.Fatalf("coult not initialize http server: %v", err)
	}
}
