package client

import (
	"encoding/json"
	"log"
	"ssorry/internal/store"
	"ssorry/internal/types"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type TokenResponse struct {
	Token string `json:"token"`
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := parseSession(ctx)
		//TODO make this configurable? Can I do it automatically?
		resp, err := resty.New().R().SetBody(session).Post("localhost:8010/api/v0/auth/sso/tokens")
		if err != nil {
			log.Println("Error fetching token:", err)
		}
		if resp.IsError() {
			log.Println("Error resp fetching token:", resp)
		}
		var token TokenResponse
		if err := json.Unmarshal(resp.Body(), &token); err != nil {
			log.Println("Error parsing body:", err)
		} else {
			store.SaveToken(token.Token)
		}
	}
}

func parseSession(ctx *gin.Context) types.Session {
	return types.Session{
		Code:         ctx.Query("code"),
		State:        ctx.Query("state"),
		SessionState: ctx.Query("session_state"),
	}
}
