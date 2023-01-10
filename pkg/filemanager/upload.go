package filemanager

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path/filepath"
	"peanut/pkg/ary"
	"strings"

	"github.com/google/uuid"
)

func SaveUploadedFileTo(ctx *gin.Context, file *multipart.FileHeader, path string) (error, string) {
	filename := uuid.NewString() + "_" + filepath.Base(file.Filename)
	filePathServer := path + filename
	err := ctx.SaveUploadedFile(file, filePathServer)
	if err != nil {
		return err, ""
	}
	return nil, filePathServer
}

func CheckExtensionAvailable(ext string, listExt []string) bool {
	ext = strings.ToLower(ext)
	return ary.InArray(ext, listExt)
}
