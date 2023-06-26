package types

type Session struct {
	Code         string `json:"sessionCode"`
	State        string `json:"sessionId"`
	SessionState string `json:"sessionState"`
}
