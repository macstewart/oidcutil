package util

import (
	"fmt"
	"ssorry/internal/types"

	"github.com/zitadel/oidc/pkg/oidc"
)

func BuildDiscovery(hostname string) *oidc.DiscoveryConfiguration {
	return &oidc.DiscoveryConfiguration{
		TokenEndpoint:         hostname + "/token",
		AuthorizationEndpoint: hostname + "/authorize",
		JwksURI:               hostname + "/keys",
	}
}

func BuildRedirect(params types.AuthParams) string {
	return fmt.Sprintf("%s?code=%s&state=%s&session_state=%s", params.Redirect, params.Code(), params.State, params.State)
}
