package http

type CreateVehicle struct {
	VehicleID string  `json:"vehicle_id"`
	Brand     string  `json:"brand"`
	Model     string  `json:"model"`
	Year      int     `json:"year"`
	Color     string  `json:"color"`
	Price     float64 `json:"price"`
}
