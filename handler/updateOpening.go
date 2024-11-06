package handler

import (
	"net/http"

	"github.com/ViniSpirit/Gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)
	if err := request.Validate(ctx); err != nil {
		logger.Errorf("Validation request error %v", err.Error())
		sendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendErrorResponse(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendErrorResponse(ctx, http.StatusNotFound, "opening not found")
		return
	}

	// update the opening
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
	// save the updated opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		sendErrorResponse(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}

	sendSuccessResponse(ctx, "opening updated successfully", opening)
}
