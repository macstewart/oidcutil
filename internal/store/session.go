package store

import (
	"log"
	"ssorry/internal/types"
)

var (
	auth          types.AuthParams
	token         string
	tokenCallback chan string
)

func LoadAuth() types.AuthParams {
	log.Println("load auth", auth)
	return auth
}

func SaveAuth(new types.AuthParams) {
	log.Println("save auth", new)
	auth = new
}

func SetCallback(cb chan string) {
	tokenCallback = cb
}

func SaveToken(input string) {
	token = input
	tokenCallback <- token
}
