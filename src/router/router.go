package router

import (
	"fmt"
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

	router.Static(env.BaseRoute, env.StaticDir)
	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, env.BaseRoute)
	})

	router.LoadHTMLGlob(fmt.Sprintf("%s404.html", env.StaticDir))
	router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", gin.H{})
	})

	return router
}
