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

func (c *UserController) GetUsers(ctx *gin.Context) {
	users, err := c.Usecase.GetUsers(ctx)
	if checkError(ctx, err) {
		return
	}
	response.OK(ctx, users)
}

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
