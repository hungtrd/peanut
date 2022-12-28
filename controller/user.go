package controller

import (
	"peanut/domain"
	"peanut/pkg/jwt"
	"peanut/pkg/response"
	"peanut/repository"
	"peanut/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	Usecase usecase.UserUsecase
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		Usecase: usecase.NewUserUsecase(repository.NewUserRepo(db)),
	}
}

// Login godoc
//
//	@Summary		login
//	@Description	login
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			json	body		domain.RequestLogin	true	"Body"
//	@Success		200		{object}	domain.Response
//	@Failure		400		{object}	domain.ErrorResponse
//	@Failure		500		{object}	domain.ErrorResponse
//	@Router			/login [post]
func (c *UserController) Login(ctx *gin.Context) {
	req := domain.RequestLogin{}
	if !bindJSON(ctx, &req) {
		return
	}
	token, err := c.Usecase.Login(ctx, req)
	if checkError(ctx, err) {
		return
	}
	response.OK(ctx, gin.H{
		"token": "Bearer " + token,
	})
}

// GetUsers godoc
//
//	@Summary		get list users
//	@Description	get list users
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]domain.User
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		404	{object}	domain.ErrorResponse
//	@Failure		500	{object}	domain.ErrorResponse
//	@Security		Bearer
//	@Router			/users [get]
func (c *UserController) GetUsers(ctx *gin.Context) {
	users, err := c.Usecase.GetUsers(ctx)
	if checkError(ctx, err) {
		return
	}
	response.OK(ctx, users)
}

// CurrentUser godoc
//
//	@Summary		Get current user
//	@Description	Get current user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	domain.User
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		404	{object}	domain.ErrorResponse
//	@Failure		500	{object}	domain.ErrorResponse
//	@Security		BearerAuth
//	@Router			/users/current [get]
func (c *UserController) CurrentUser(ctx *gin.Context) {
	userID, err := jwt.ExtractTokenID(ctx)
	if checkError(ctx, err) {
		return
	}

	u, err := c.Usecase.GetUser(ctx, userID)
	if checkError(ctx, err) {
		return
	}

	response.OK(ctx, u)
}

// GetUser godoc
//
//	@Summary		Get a user
//	@Description	Get a user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	domain.User
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		404	{object}	domain.ErrorResponse
//	@Failure		500	{object}	domain.ErrorResponse
//	@Security		Bearer
//	@Router			/users/{id} [get]
func (c *UserController) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if checkError(ctx, err) {
		return
	}
	user, err := c.Usecase.GetUser(ctx, id)
	if checkError(ctx, err) {
		return
	}
	response.OK(ctx, user)
}

// CreateUser godoc
//
//	@Summary		create a user
//	@Description	Create a user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			json	formData	domain.User	true	"Body"
//	@Success		200		{object}	domain.User
//	@Failure		400		{object}	domain.ErrorResponse
//	@Failure		500		{object}	domain.ErrorResponse
//	@Security		Bearer
//	@Router			/users [post]
func (c *UserController) CreateUser(ctx *gin.Context) {
	user := domain.User{}
	if !bindJSON(ctx, &user) {
		return
	}

	err := c.Usecase.CreateUser(ctx, user)
	if checkError(ctx, err) {
		return
	}
	response.OK(ctx, nil)
}

// Register godoc
//
//	@Summary		register
//	@Description	register
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			json	body		object	true	"Body"
//	@Success		200		{object}	domain.Response
//	@Failure		400		{object}	domain.ErrorResponse
//	@Failure		500		{object}	domain.ErrorResponse
//	@Router			/register [post]
func (c *UserController) Register(ctx *gin.Context) {
	user := domain.User{}
	if !bindJSON(ctx, &user) {
		return
	}

	err := c.Usecase.CreateUser(ctx, user)
	if checkError(ctx, err) {
		return
	}
	response.OK(ctx, nil)
}
