package router

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"nikurasu.gay/static-hoster/api"
	"nikurasu.gay/static-hoster/envloader"
)

func Create(env *envloader.Environment) *fiber.App {
	router := fiber.New()

	router.Static("", env.StaticDir)

	apiRoutes := router.Group("/api", basicauth.New(basicauth.Config{
		Users: map[string]string{
			env.User: env.ApiKey,
		},
	}))
	apiRoutes.Post("/update", api.PostUpdate(env))
	// Ping test
	router.Get("/ping", func(c *fiber.Ctx) error {
		c.Status(http.StatusOK)
		return c.SendString("Pong!")
	})

	// Use the "old" method becuase I don't know how to pass the error code to the user generated html
	router.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusNotFound).SendFile(fmt.Sprintf("%s404.html", env.StaticDir))
	})

	return router
}
