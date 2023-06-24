package client

import (
	"oidcutil/internal/types"

	"github.com/gin-gonic/gin"
)

func IdpCallback(cb chan types.Callback) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		println("callback", ctx.Query("state"), ctx.Query("code"))
		cb <- types.Callback{
			Code:  ctx.Query("code"),
			State: ctx.Query("state"),
		}
	}
}
