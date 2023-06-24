package types

type User struct {
	Email  string   `json:"email,omitempty"`
	Key    string   `json:"key,omitempty"`
	Values []string `json:"values,omitempty"`
}
