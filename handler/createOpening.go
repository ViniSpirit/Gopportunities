package handler

import (
	"net/http"

	"github.com/ViniSpirit/Gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(ctx); err != nil {
		logger.Errorf("error validating request: %v", err.Error())
		sendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendErrorResponse(ctx, http.StatusInternalServerError, "error creating opening")
		return
	}

	sendSuccessResponse(ctx, "create opening", opening)
}
