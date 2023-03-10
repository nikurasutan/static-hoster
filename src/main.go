package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"nikurasu.gay/static-hoster/envloader"
	"nikurasu.gay/static-hoster/helper"
	"nikurasu.gay/static-hoster/router"
)

//go:embed static/*
var assets embed.FS

func main() {
	var env = envloader.Load()
	if empty, _ := helper.IsEmpty(env.StaticDir); empty {
		copyEmbed(env.StaticDir)
		fmt.Println("Created default files")
	}
	copyEmbed(env.StaticDir)
	r := router.Create(env)
	r.Run(env.Port)
}

func copyEmbed(dest string) {
	files, err := assets.ReadDir("static")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		fileContent, err := assets.ReadFile(fmt.Sprintf("static/%s", file.Name()))
		if err != nil {
			fmt.Println(err)
		}
		fileName := fmt.Sprintf("%s%s", dest, file.Name())
		if err := os.WriteFile(fileName, fileContent, 0666); err != nil {
			log.Printf("Error writing default Files: %s", err)
		}
	}
}
