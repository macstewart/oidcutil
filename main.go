package main

import (
	"encoding/gob"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Callback struct {
	state string
	code  string
}

var callbackChan = make(chan Callback, 1)

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
	router.GET("/sso/login", loginCallback())
	router.GET("/callback", fetchCallback())
	return router
}

func loginCallback() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		println("callback", ctx.Query("state"), ctx.Query("code"))
	}
}

func fetchCallback() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		callback := <-callbackChan
		ctx.JSON(http.StatusOK, callback)
	}
}
