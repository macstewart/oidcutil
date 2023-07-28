package util

import (
	"fmt"
	"ssorry/internal/types"

	"github.com/zitadel/oidc/pkg/oidc"
)

func BuildDiscovery(local bool, hostname string) *oidc.DiscoveryConfiguration {
	authEndpoint := hostname + "/authorize"
	if local {
		authEndpoint = "http://localhost:3333/authorize"
	}
	return &oidc.DiscoveryConfiguration{
		TokenEndpoint:         hostname + "/token",
		AuthorizationEndpoint: authEndpoint,
		JwksURI:               hostname + "/keys",
	}
}

func BuildRedirect(params types.AuthParams) string {
	return fmt.Sprintf("%s?code=%s&state=%s&session_state=%s", params.Redirect, params.Code(), params.State, params.State)
}

func BuildCode(state string) string {
	return state + "code"
}
