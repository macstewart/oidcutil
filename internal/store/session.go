package store

import (
	"log"
	"ssorry/internal/types"
)

var (
	auth            types.AuthParams
	token           string
	tokenCallback   chan string
	session         string
	sessionCallback chan string
)

func LoadAuth() types.AuthParams {
	log.Println("load auth", auth)
	return auth
}

func SaveAuth(new types.AuthParams) {
	log.Println("save auth", new)
	auth = new
}

func SetTokenCallback(cb chan string) {
	tokenCallback = cb
}

func SetSessionCallback(cb chan string) {
	sessionCallback = cb
}

func SaveToken(input string) {
	token = input
	tokenCallback <- token
}

func SaveSession(input string) {
	log.Println("save session", input)
	session = input
	sessionCallback <- session
}
