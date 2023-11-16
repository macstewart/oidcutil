package client

import (
	"net/http"

	"ssorry/internal/types"

	"github.com/gin-gonic/gin"
)

func FetchCodeCallback(cb chan types.Callback) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		callback := <-cb
		ctx.JSON(http.StatusOK, callback)
	}
}

func FetchTokenCallback(cb chan string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := <-cb
		ctx.JSON(http.StatusOK, token)
	}
}

func FetchSessionCallback(cb chan string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := <-cb
		ctx.JSON(http.StatusOK, session)
	}
}
