package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"nikurasu.gay/static-hoster/api"
	"nikurasu.gay/static-hoster/envloader"
	"nikurasu.gay/static-hoster/middleware/auth"
)

func Create(env *envloader.Environment) *gin.Engine {
	router := gin.Default()

	apiRoutes := router.Group("/api", auth.AuthMiddleware(env))
	{
		apiRoutes.POST("/update", api.PostUpdate(env))
	}
	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.Static("/home", env.StaticDir)

	// TODO: Load 404 error Page

	return router
}
