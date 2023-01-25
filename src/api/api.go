package api

import (
	"net/http"
	"os"

	"github.com/c4milo/unpackit"
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
	file, err := os.Open("./" + siteUpdate.Filename)
	unpackit.Unpack(file, "./")

	ctx.JSON(http.StatusOK, gin.H{"data": siteUpdate.Filename, "err": err})
}
