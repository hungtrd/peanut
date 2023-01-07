package filemanager

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"path/filepath"
	"peanut/pkg/ary"
	"strconv"
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

func CheckMaxSizeUpload(size int) bool {
	maxSize, err := strconv.Atoi(os.Getenv("MAX_SIZE_UPLOAD"))
	if err != nil {
		return false
	}
	return !(size > maxSize)
}

func CheckExtensionAvailable(ext string, listExt []string) bool {
	ext = strings.ToLower(ext)
	return ary.InArray(ext, listExt)
}
