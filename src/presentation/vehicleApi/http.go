package vehicleApi

import (
	"net/http"

	interfaces "github.com/caiiomp/vehicle-platform-core/src/core/_interfaces"
	"github.com/caiiomp/vehicle-platform-core/src/core/responses"
	"github.com/gin-gonic/gin"
)

type vehicleApi struct {
	vehicleService interfaces.VehicleService
}

func RegisterVehicleRoutes(app *gin.Engine, vehicleService interfaces.VehicleService) {
	service := vehicleApi{
		vehicleService: vehicleService,
	}

	app.POST("/vehicles", service.create)
	app.PATCH("/vehicles/:vehicle_id", service.update)
	app.POST("/vehicles/webhook", service.webhook)
}

// Create godoc
// @Summary Create Vehicle
// @Description Create a vehicle
// @Tags Vehicle
// @Accept json
// @Produce json
// @Param user body vehicleApi.createVehicleRequest true "Body"
// @Success 201 {object} responses.Vehicle
// @Failure 204 {object} responses.ErrorResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /vehicles [post]
func (ref *vehicleApi) create(ctx *gin.Context) {
	var request createVehicleRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	vehicle, err := ref.vehicleService.Create(ctx, *request.ToDomain())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if vehicle == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	response := responses.VehicleFromDomain(*vehicle)
	ctx.JSON(http.StatusCreated, response)
}

// Create godoc
// @Summary Update Vehicle
// @Description Update a vehicle
// @Tags Vehicle
// @Accept json
// @Produce json
// @Param user body vehicleApi.updateVehicleRequest false "Body"
// @Success 200 {object} responses.Vehicle
// @Failure 204 {object} responses.ErrorResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /vehicles/{vehicle_id} [patch]
func (ref *vehicleApi) update(ctx *gin.Context) {
	var uri vehicleURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	var request updateVehicleRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	vehicle, err := ref.vehicleService.Update(ctx, uri.VehicleID, *request.ToDomain())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if vehicle == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	response := responses.VehicleFromDomain(*vehicle)
	ctx.JSON(http.StatusOK, response)
}

// Create godoc
// @Summary Vehicle Webhook
// @Description Vehicle Webhook
// @Tags Vehicle
// @Accept json
// @Produce json
// @Param user body vehicleApi.vehicleWebhookRequest true "Body"
// @Success 200 {object} responses.Vehicle
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /vehicles/webhook [post]
func (ref *vehicleApi) webhook(ctx *gin.Context) {
	var request vehicleWebhookRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	response := request

	ctx.JSON(http.StatusOK, response)
}
