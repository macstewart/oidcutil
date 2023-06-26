package idp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zitadel/oidc/pkg/oidc"
)

func ServeDiscovery(content *oidc.DiscoveryConfiguration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, content)
	}
}
