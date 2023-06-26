package store

import (
	"log"
	"ssorry/internal/types"
)

var user = types.User{
	Email:  "admin@macstew.art",
	Key:    "groups",
	Values: []string{"admin", "viewer"},
}

func LoadUser() types.User {
	//TODO load from temp storage
	log.Printf("%v", user.ValueJoin())
	return user
}

func SaveUser(new types.User) {
	user = new
}
