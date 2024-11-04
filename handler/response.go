package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{"message": message, "errorCode": code})
}

func sendSuccessResponse(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("operation %s successful", operation), "data": data})
}
