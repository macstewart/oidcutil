package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"oidcutil/internal/endpoint/client"
	"oidcutil/internal/types"

	"github.com/gin-gonic/gin"
)

var callbackChan = make(chan types.Callback, 1)

func main() {
	rtr := router()
	log.Print("Server listening on http://localhost:3333")
	if err := http.ListenAndServe("0.0.0.0:3333", rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}

func router() *gin.Engine {
	router := gin.Default()
	gob.Register(map[string]interface{}{})
	router.GET("/callback", client.FetchCallback(callbackChan))
	router.GET("/sso/login", client.IdpCallback(callbackChan))
	return router
}

func loginCallback() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		println("callback", ctx.Query("state"), ctx.Query("code"))
	}
}
