package helper

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InvokeError(ctx *gin.Context, errorMessage string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": errors.New(errorMessage),
	})
}
