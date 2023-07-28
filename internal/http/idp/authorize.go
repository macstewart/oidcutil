package idp

import (
	"log"
	"net/http"
	"ssorry/internal/store"
	"ssorry/internal/types"

	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("origin %s", ctx.Request.Header.Get("Origin"))
		params := parseParams(ctx)
		store.SaveAuth(params)
		user := store.DefaultUser()
		ctx.HTML(http.StatusOK, "auth.tmpl", gin.H{"state": params.State, "email": user.Email, "key": user.Key, "values": user.ValueJoin()})
	}
}

func parseParams(ctx *gin.Context) types.AuthParams {
	return types.AuthParams{
		ResponseType: ctx.Query("response_type"),
		Redirect:     ctx.Query("redirect_uri"),
		Scope:        ctx.Query("scope"),
		State:        ctx.Query("state"),
	}
}
