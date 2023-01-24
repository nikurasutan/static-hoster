package api

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func PostUpdate(ctx *gin.Context) {
	siteUpdate, err := ctx.FormFile("Hello")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "File not recieved",
		})
		return
	}
	if err := ctx.SaveUploadedFile(siteUpdate, "./"+siteUpdate.Filename); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save file",
		})
		return
	}
	unTar("/", "./"+siteUpdate.Filename, ctx)

	ctx.JSON(http.StatusOK, gin.H{"data": siteUpdate.Filename, "err": err})
}

func unTar(dst string, src string, ctx *gin.Context) error {
	file, err := os.Open(src)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to open saved file",
		})
	}
	gzr, err := gzip.NewReader(file)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to unzip file",
		})
	}
	defer gzr.Close()
	tar := tar.NewReader(gzr)
	for {
		header, err := tar.Next()
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case header == nil:
			continue
		}
	}
}
