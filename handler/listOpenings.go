package handler

import (
	"net/http"

	"github.com/ViniSpirit/Gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendErrorResponse(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccessResponse(ctx, "list-openings", openings)
}
