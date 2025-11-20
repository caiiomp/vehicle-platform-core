package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type VehiclePlatformSalesHttpClient interface {
	CreateVehicle(ctx context.Context, vehicleID, brand, model, color string, year int, price float64) error
	UpdateVehicle(ctx context.Context, vehicleID, brand, model, color string, year int, price float64) error
}

type vehiclePlatformSalesHttpClient struct {
	client *http.Client
	host   string
}

func NewVehiclePlatformSalesHttpClient(client *http.Client, host string) VehiclePlatformSalesHttpClient {
	return &vehiclePlatformSalesHttpClient{
		client: client,
		host:   host,
	}
}

func (ref *vehiclePlatformSalesHttpClient) CreateVehicle(ctx context.Context, vehicleID, brand, model, color string, year int, price float64) error {
	url := ref.host + "/vehicles"

	vehicle := CreateVehicle{
		VehicleID: vehicleID,
		Brand:     brand,
		Model:     model,
		Year:      year,
		Color:     color,
		Price:     price,
	}

	data, err := json.Marshal(vehicle)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := ref.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	rawResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create vehicle at vehicle platform sales: %v", string(rawResponse))
	}

	return nil
}

func (ref *vehiclePlatformSalesHttpClient) UpdateVehicle(ctx context.Context, vehicleID string, brand string, model string, color string, year int, price float64) error {
	url := ref.host + "/vehicles/" + vehicleID

	vehicle := UpdateVehicle{
		Brand: brand,
		Model: model,
		Year:  year,
		Color: color,
		Price: price,
	}

	data, err := json.Marshal(vehicle)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := ref.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	rawResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update vehicle at vehicle platform sales: %v", string(rawResponse))
	}

	return nil
}
