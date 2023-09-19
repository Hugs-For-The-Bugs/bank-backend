package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(context *gin.Context, data interface{}) {
	context.IndentedJSON(http.StatusOK, data)
}

func BadRequestResponse(context *gin.Context, msg string) {
	context.JSON(http.StatusBadRequest, gin.H{
		"msg": msg,
	})
}

func ServerErrorResponse(context *gin.Context, msg string) {
	context.JSON(http.StatusInternalServerError, gin.H{
		"msg": msg,
	})
}
