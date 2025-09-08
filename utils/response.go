package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendJSON(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}

func SendStatusMessage(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": message,
	})
}

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"status": "error",
		"error":  msg,
	})
}
