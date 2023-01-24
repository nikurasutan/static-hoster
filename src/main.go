package main

import (
	"nikurasu.gay/static-hoster/router"
)

func main() {
	r := router.Create()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
