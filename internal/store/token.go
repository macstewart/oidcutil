package store

import (
	"ssorry/internal/types"
)

func LoadUser() types.User {
	//TODO load from temp storage
	return types.User{
		Email:  "admin@macstew.art",
		Key:    "groups",
		Values: []string{"admin", "viewer"},
	}
}

func SaveUser(user types.User) {
	panic("not implemented")
}
