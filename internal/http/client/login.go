package client

import (
	"encoding/json"
	"fmt"
	"log"
	"ssorry/internal/store"
	"ssorry/internal/types"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const SESSION_COOKIE_NAME = "SESSIONID"

var (
	tokenPath   = "/api/v0/auth/sso/tokens"
	sessionPath = "/api/v2/platform/auth/sso/login"
	host        = "http://localhost:8010"
)

type TokenResponse struct {
	Token string `json:"token"`
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := parseSession(ctx)
		//TODO make this configurable? Can I do it automatically?
		resp, err := resty.New().R().SetBody(session).Post(fmt.Sprintf("%s%s", host, tokenPath))
		if err != nil {
			log.Println("Error fetching token:", err)
		}
		if resp.IsError() {
			log.Println("Error fetching token:", resp.Error())
		}
		var token TokenResponse
		if err := json.Unmarshal(resp.Body(), &token); err != nil {
			log.Println("Error parsing body:", err)
		} else {
			store.SaveToken(token.Token)
		}
	}
}

func LoginSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := parseSession(ctx)
		//TODO make this configurable? Can I do it automatically?
		resp, err := resty.New().R().SetBody(session).Post(fmt.Sprintf("%s%s", host, sessionPath))
		if err != nil {
			log.Println("Error fetching session:", err)
			return
		}
		if resp.IsError() {
			log.Println("Error fetching session:", resp.Error(), resp.StatusCode())
			return
		}
		for _, cookie := range resp.Cookies() {
			if cookie.Name == SESSION_COOKIE_NAME {
				store.SaveSession(cookie.Value)
				return
			}
		}
		log.Println("Error fetching session: no session cookie")
	}
}

func parseSession(ctx *gin.Context) types.Session {
	return types.Session{
		Code:         ctx.Query("code"),
		State:        ctx.Query("state"),
		SessionState: ctx.Query("session_state"),
	}
}
