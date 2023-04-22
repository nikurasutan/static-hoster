package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/c4milo/unpackit"
	"github.com/gofiber/fiber/v2"
	"nikurasu.gay/static-hoster/envloader"
)

type retrunVal struct {
	Data string `json:"data"`
	Err  error  `json:"err"`
}

func PostUpdate(env *envloader.Environment) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		siteUpdate, err := ctx.FormFile("STATIC_PAGE")
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.SendString("File not recieved")
		}
		if err := ctx.SaveFile(siteUpdate, fmt.Sprintf("%s%s", env.RootDir, siteUpdate.Filename)); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.SendString("Unable to save file")
		}
		file, err := os.Open(fmt.Sprintf("%s%s", env.RootDir, siteUpdate.Filename))
		os.RemoveAll(env.StaticDir)
		os.Mkdir(env.StaticDir, os.ModePerm)
		unpackit.Unpack(file, env.StaticDir)
		os.RemoveAll(fmt.Sprintf("%s%s", env.RootDir, siteUpdate.Filename))
		retrunArray := new(retrunVal)
		retrunArray.Data = siteUpdate.Filename
		retrunArray.Err = err
		ctx.JSON(retrunArray)
		return ctx.SendStatus(http.StatusOK)
	}
}
