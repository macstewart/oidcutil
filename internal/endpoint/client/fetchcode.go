package client

import (
	"net/http"

	"oidcutil/internal/types"

	"github.com/gin-gonic/gin"
)

func FetchCallback(cb chan types.Callback) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		callback := <-cb
		ctx.JSON(http.StatusOK, callback)
	}
}
