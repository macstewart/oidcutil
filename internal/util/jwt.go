package util

import (
	"log"
	"ssorry/internal/store"
	"ssorry/internal/types"

	jwt "github.com/dgrijalva/jwt-go"
)

type TokenResponse struct {
	IdToken     string `json:"id_token"`
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

func Build(user types.User) TokenResponse {
	return TokenResponse{
		IdToken:     buildIdToken(user),
		TokenType:   "Bearer",
		AccessToken: "idk",
	}
}

func buildIdToken(user types.User) string {
	idToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{})
	idClaims := idToken.Claims.(jwt.MapClaims)
	idClaims["sub"] = "hi"
	idClaims["email"] = user.Email
	idClaims["email_verified"] = true
	idClaims["aud"] = "local"
	if user.Key != "" {
		idClaims[user.Key] = user.Values
	}
	tokenString, err := idToken.SignedString(store.LoadPrivateKey())
	if err != nil {
		log.Println("Error signing token string:", err)
	}
	return tokenString
}
