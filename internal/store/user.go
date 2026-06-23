package store

import (
	"log"
	"ssorry/internal/types"
	"sync"
)

var maplock sync.Mutex

var users map[string]types.User

func init() {
	users = make(map[string]types.User)
}

func LoadUser(state string) types.User {
	maplock.Lock()
	defer maplock.Unlock()
	if user, ok := users[state]; ok {
		delete(users, state)
		return user
	} else {
		log.Printf("no user found for state %s", state)
		return DefaultUser()
	}
}

func SaveUser(state string, newUser types.User) {
	maplock.Lock()
	defer maplock.Unlock()
	log.Printf("save user %s : %s", state, newUser)
	users[state] = newUser
}

func DefaultUser() types.User {
	return types.User{
		Email:  "testuser@mymaas.net",
		Key:    "groups",
		Values: []string{"admin", "viewer"},
	}

}
