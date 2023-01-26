package main

import (
	"nikurasu.gay/static-hoster/envloader"
	"nikurasu.gay/static-hoster/router"
)

func main() {
	var env = envloader.Load()
	r := router.Create(env)
	r.Run(env.Port)
}
