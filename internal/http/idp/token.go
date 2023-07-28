package idp

import (
	"net/http"
	"ssorry/internal/store"
	"ssorry/internal/util"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Code string `json:"code"`
}

func ServeToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// if ctx.ShouldBindJSON(&code) != nil {
		user := store.LoadUser(ctx.PostForm("code"))
		response := util.Build(user)
		ctx.JSON(http.StatusOK, response)
		// } else {
		// 	ctx.JSON(http.StatusBadRequest, "mfw no code :(")
		// }
	}
}
