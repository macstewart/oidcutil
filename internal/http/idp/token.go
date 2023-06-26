package idp

import (
	"net/http"
	"ssorry/internal/store"
	"ssorry/internal/util"

	"github.com/gin-gonic/gin"
)

func ServeToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := store.LoadUser()
		response := util.Build(user)
		ctx.JSON(http.StatusOK, response)
	}
}
