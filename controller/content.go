package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
	"path/filepath"
	"peanut/domain"
	"peanut/pkg/filemanager"
	"peanut/repository"
	"peanut/usecase"
)

type ContentController struct {
	Content usecase.ContentUsecase
}

func NewContentController(db *gorm.DB) *ContentController {
	return &ContentController{
		Content: usecase.NewContentUsecase(repository.NewContentRepo(db)),
	}
}

// ListContent godoc
//
//	@Summary		content
//	@Description	content
//	@Tags			Content
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	domain.Content
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		500	{object}	domain.ErrorResponse
//	@Router			/contents [get]
func (c *ContentController) ListContent(ctx *gin.Context) {
	contents, err := c.Content.GetContents(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": http.StatusText(http.StatusOK),
		"data":    contents,
	})
}

// CreateContent godoc
//
//	@Summary		content
//	@Description	content
//	@Tags			Content
//	@Accept			json
//	@Produce		json
//	@Param			json	body		domain.Content	true	"Body"
//	@Success		200		{object}	domain.Response
//	@Failure		400		{object}	domain.ErrorResponse
//	@Failure		500		{object}	domain.ErrorResponse
//	@Router			/contents [post]
func (c *ContentController) CreateContent(ctx *gin.Context) {
	var content domain.CreateContent
	if !bind(ctx, &content) {
		return
	}
	if !filemanager.CheckMaxSizeUpload(int(content.Thumbnail.Size)) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "file size is too big!",
		})
		return
	}

	extension := filepath.Ext(content.Thumbnail.Filename)
	extensions := []string{".jpeg", ".png", ".jpg"}
	if !filemanager.CheckExtensionAvailable(extension, extensions) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "file type not allow",
		})
		return
	}

	err, path := filemanager.SaveUploadedFileTo(ctx, content.Thumbnail, os.Getenv("TMP_PATH"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = c.Content.CreateContent(ctx, content, path)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": http.StatusText(http.StatusOK),
	})
}
