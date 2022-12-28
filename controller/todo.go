package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"peanut/domain"
	"peanut/pkg/jwt"
	"peanut/repository"
	"peanut/usecase"
	"strconv"
)

type TodoController struct {
	Todo usecase.TodoUsecase
}

func NewTodoController(db *gorm.DB) *TodoController {
	return &TodoController{
		Todo: usecase.NewTodoUsecase(repository.NewTodoRepo(db)),
	}
}

// ListTodo godoc
//
//	@Summary		list todo
//	@Description	list todo
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	domain.ListTodo
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		500	{object}	domain.ErrorResponse
//	@Security		Bearer
//	@Router			/users/todo [post]
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

// CreateTodo godoc
//
//	@Summary		create todo
//	@Description	create todo
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	domain.Response
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		500	{object}	domain.ErrorResponse
//	@Security		Bearer
//	@Router			/users/todo [get]
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

// UpdateTodo godoc
//
//	@Summary		update todo
//	@Description	update todo
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Todo ID"
//	@Success		200	{object}	domain.Response
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		500	{object}	domain.ErrorResponse
//	@Security		Bearer
//	@Router			/users/todo [patch]
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

// DeleteTodo godoc
//
//	@Summary		delete todo
//	@Description	delete todo
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Todo ID"
//	@Success		200	{object}	domain.Response
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		500	{object}	domain.ErrorResponse
//	@Security		Bearer
//	@Router			/users/todo [delete]
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
