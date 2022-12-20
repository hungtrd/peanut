package controller

import (
	"net/http"
	"peanut/domain"
	"peanut/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookController struct {
	Usecase usecase.BookUsecase
}

func NewBookController(db *gorm.DB) *BookController {
	return &BookController{
		Usecase: usecase.NewBookUsecase(db),
	}
}

func (c *BookController) GetBook(ctx *gin.Context) {
	bookId := ctx.Param("id")
	ID, _ := strconv.Atoi(bookId)

	book, err := c.Usecase.GetBook(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.Response{false, nil, err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, domain.Response{true, book, "Found book successful"})
}

func (c *BookController) GetBooks(ctx *gin.Context) {
	books, err := c.Usecase.GetBooks()
	if err != nil {
		ctx.JSON(http.StatusNoContent, gin.H{
			"message": "not found any record",
		})
		return
	}
	ctx.JSON(http.StatusOK, domain.Response{
		Success: true, Data: books, Message: "Get data successful",
	})

}

func (c *BookController) CreateBook(ctx *gin.Context) {
	newbook := domain.Book{}
	if !bindJSON(ctx, &newbook) {
		return
	}

	book, err := c.Usecase.CreateBook(newbook)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.Response{false, nil, err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, domain.Response{true, book, "create book success"})
}

func (c *BookController) UpdateBook(ctx *gin.Context) {
	bookId := ctx.Param("id")
	ID, _ := strconv.Atoi(bookId)
	updateForm := domain.UpdateBookForm{}
	bindJSON(ctx, &updateForm)
	newbook, err := c.Usecase.UpdateBook(updateForm, ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.Response{false, nil, err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, domain.Response{true, newbook, "updated"})
}

func (c *BookController) DeleteBook(ctx *gin.Context) {
	bookId := ctx.Param("id")
	ID, _ := strconv.Atoi(bookId)
	err := c.Usecase.DeleteBook(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.Response{false, nil, err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, domain.Response{true, nil, "deleted"})

}
