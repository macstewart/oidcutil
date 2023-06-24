package util

import "github.com/zitadel/oidc/pkg/oidc"

func BuildDiscovery(hostname string) *oidc.DiscoveryConfiguration {
	return &oidc.DiscoveryConfiguration{
		TokenEndpoint:         hostname + "/token",
		AuthorizationEndpoint: hostname + "/authorize",
		JwksURI:               hostname + "/keys",
	}
}
