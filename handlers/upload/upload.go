package upload

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

const (
	saveDir = "./data"
)

func init() {
	os.MkdirAll(saveDir, 0664)
}

// HandleSingalFile handle POST /data for update data
// https://gin-gonic.com/docs/examples/upload-file/single-file/
func HandleSingalFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Println("get formfile from ctx had error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "get file error",
		})
		return
	}

	// Upload the file to specific dst.
	log.Println("receive file, name is", file.Filename)
	err = ctx.SaveUploadedFile(file, path.Join(saveDir, file.Filename))
	if err != nil {
		log.Println("SaveUploadedFile had error", err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("%s uploaded", file.Filename),
	})
}

// HandleMultipleFile support upload multiple files
// https://gin-gonic.com/docs/examples/upload-file/multiple-file/
func HandleMultipleFile(ctx *gin.Context) {
	// Multipart form
	form, _ := ctx.MultipartForm()
	files := form.File["files[]"]

	for _, file := range files {
		log.Println("receive file, name is", file.Filename)

		err := ctx.SaveUploadedFile(file, path.Join(saveDir, file.Filename))
		if err != nil {
			log.Println("SaveUploadedFile had error", err)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("%d files uploaded", len(files)),
	})
}
