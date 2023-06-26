package idp

import (
	"log"
	"net/http"
	"strings"

	"ssorry/internal/store"
	"ssorry/internal/types"
	"ssorry/internal/util"

	"github.com/gin-gonic/gin"
)

func UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := parseUser(ctx)
		store.SaveUser(user)
		uri := util.BuildRedirect(store.LoadAuth())
		ctx.Redirect(http.StatusTemporaryRedirect, uri)
	}
}

func parseUser(ctx *gin.Context) types.User {
	var user types.User
	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, "")
		log.Printf("bind error: %v", err)
	}
	user.Values = []string{} // lol
	for _, val := range strings.Split(ctx.PostForm("values"), ",") {
		user.Values = append(user.Values, strings.TrimSpace(val))
	}
	return user
}
