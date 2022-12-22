package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"peanut/domain"
	"peanut/pkg/jwt"
	"peanut/usecase"
	"strconv"
)

type TodoController struct {
	Todo usecase.TodoUsecase
}

func NewTodoController(db *gorm.DB) *TodoController {
	return &TodoController{
		Todo: usecase.NewTodoUsecase(db),
	}
}

func (c *TodoController) ListTodo(ctx *gin.Context) {
	userID, _ := jwt.ExtractTokenID(ctx)
	todos, err := c.Todo.ListTodo(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": http.StatusText(http.StatusOK),
		"data":    todos,
	})
}

func (c *TodoController) CreateTodo(ctx *gin.Context) {
	todo := domain.Todo{}
	err := ctx.ShouldBindJSON(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userID, _ := jwt.ExtractTokenID(ctx)
	err = c.Todo.CreateTodo(userID, todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": http.StatusText(http.StatusCreated),
	})
	return
}

func (c *TodoController) UpdateTodo(ctx *gin.Context) {
	todoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	todo := domain.Todo{}
	err = ctx.ShouldBindJSON(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = c.Todo.UpdateTodo(todoID, todo)
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

func (c *TodoController) DeleteTodo(ctx *gin.Context) {
	userID, _ := jwt.ExtractTokenID(ctx)
	todoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = c.Todo.DeleteTodo(userID, todoID)
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
