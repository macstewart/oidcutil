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
		state := ctx.Query("state")
		log.Printf("ctx query params %v", state)
		if state == "" {
			state = "default"
		}
		store.SaveUser(util.BuildCode(state), user)
		uri := util.BuildRedirect(store.LoadAuth())
		log.Printf("redir %s", uri)
		ctx.Redirect(http.StatusFound, uri)
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
