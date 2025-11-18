package vehiclerepository

const (
	getVehicleByID = "SELECT * FROM vehicles WHERE id = $1;"

	insertVehicle = `
		INSERT INTO vehicles (id, brand, model, year, color, price) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *;
	`

	updateVehicle = `
		UPDATE vehicles
		SET
			brand = COALESCE($2, brand),
			model = COALESCE($3, model),
			year = COALESCE($4, year),
			color = COALESCE($5, color),
			price = COALESCE($6, price)
		WHERE id = $1
		RETURNING *;
	`
)
