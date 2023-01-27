package main

import (
	"errors"
	"fmt"
	"os"

	"nikurasu.gay/static-hoster/envloader"
	"nikurasu.gay/static-hoster/router"
)

func main() {
	var env = envloader.Load()
	if _, err := os.Stat(fmt.Sprintf("%s404.html", env.StaticDir)); errors.Is(err, os.ErrNotExist) {
		os.Create(fmt.Sprintf("%s404.html", env.StaticDir))
	}
	r := router.Create(env)
	r.Run(env.Port)
}
