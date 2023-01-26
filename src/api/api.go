package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/c4milo/unpackit"
	"github.com/gin-gonic/gin"
	"nikurasu.gay/static-hoster/envloader"
)

func PostUpdate(env *envloader.Environment) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println(env.ApiKey)
		siteUpdate, err := ctx.FormFile("STATIC_PAGE")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "File not recieved",
			})
			return
		}
		if err := ctx.SaveUploadedFile(siteUpdate, fmt.Sprintf("%s%s", env.RootDir, siteUpdate.Filename)); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save file",
			})
			return
		}
		file, err := os.Open(fmt.Sprintf("%s%s", env.RootDir, siteUpdate.Filename))
		os.RemoveAll(env.StaticDir)
		os.Mkdir(env.StaticDir, os.ModePerm)
		unpackit.Unpack(file, env.StaticDir)
		os.RemoveAll(fmt.Sprintf("%s%s", env.RootDir, siteUpdate.Filename))
		ctx.JSON(http.StatusOK, gin.H{"data": siteUpdate.Filename, "err": err})
	}
}
