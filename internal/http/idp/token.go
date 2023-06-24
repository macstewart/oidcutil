package idp

import (
	"net/http"

	"ssorry/internal/store"

	"github.com/gin-gonic/gin"
)

func ServeToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, store.LoadUser())
	}
}
