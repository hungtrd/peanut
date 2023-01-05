package controller

import (
	"peanut/domain"
	"peanut/pkg/response"
	"peanut/repository"
	"peanut/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookController struct {
	Usecase usecase.BookUsecase
}

func NewBookController(db *gorm.DB) *BookController {
	return &BookController{
		Usecase: usecase.NewBookUsecase(repository.NewBookRepo(db)),
	}
}

// CreateBook godoc
//
//	@Summary		Create book
//	@Description	Create book
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Param			param	body	domain.CreateBookReq	true	"Create book request"
//	@Success		200	{object}	domain.Response{data=domain.CreateBookResp}
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		500	{object}	domain.ErrorResponse
//	@Security		Bearer
//	@Router			/v1/books [post]
func (c *BookController) CreateBook(ctx *gin.Context) {
	req := domain.CreateBookReq{}
	if !bindJSON(ctx, &req) {
		return
	}

	resp, err := c.Usecase.CreateBook(ctx, req)
	if checkError(ctx, err) {
		return
	}

	response.OK(ctx, resp)
}
