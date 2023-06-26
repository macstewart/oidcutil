package types

import "strings"

type User struct {
	Email  string   `form:"email,omitempty"`
	Key    string   `form:"key,omitempty"`
	Values []string `form:"values,omitempty"`
}

func (u User) ValueJoin() string {
	return strings.Join(u.Values, ", ")
}
