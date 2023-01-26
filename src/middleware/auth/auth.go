package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"nikurasu.gay/static-hoster/envloader"
)

func AuthMiddleware(env *envloader.Environment) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Header["Static-Hoster-Key"][0] != env.ApiKey {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		ctx.Next()
	}
}
