package main

import (
	"nikurasu.gay/static-hoster/envloader"
	"nikurasu.gay/static-hoster/router"
)

func main() {
	env := envloader.Load()
	r := router.Create(env)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
