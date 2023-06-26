package store

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

type Jwks struct {
	Keys []map[string]interface{} `json:"keys"`
}

func init() {
	var err error
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey = &privateKey.PublicKey
}

func LoadPrivateKey() *rsa.PrivateKey {
	return privateKey
}

func LoadJwks() Jwks {
	n := base64.StdEncoding.EncodeToString(publicKey.N.Bytes())
	key := map[string]interface{}{
		"e":   "AQAB",
		"n":   n,
		"x5c": []string{},
		"x5t": "7jxY8Ml6u5tZJQzUZzK6gk7Qq3A",
		"kid": "2011-04-29",
		"kty": "RSA",
		"alg": "RS256",
		"use": "sig",
	}
	return Jwks{Keys: []map[string]interface{}{key}}
}
