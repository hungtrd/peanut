package controller

import (
	"net/http"
	"peanut/domain"
	"peanut/pkg/response"
	"peanut/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	Usecase usecase.UserUsecase
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		Usecase: usecase.NewUserUsecase(db),
	}
}

func (c *UserController) GetUsers(ctx *gin.Context) {

}

func (c *UserController) GetUser(ctx *gin.Context) {

}

func (c *UserController) CreateUser(ctx *gin.Context) {
	user := domain.User{}
	if !bindJSON(ctx, &user) {
		return
	}

	err := c.Usecase.CreateUser(user)
	if checkError(ctx, err) {
		return
	}

	response.OK(ctx, nil)
}

func (c *UserController) Login(ctx *gin.Context) {
	var loginForm domain.LoginForm
	bindJSON(ctx, &loginForm)
	tokenString, err := c.Usecase.Login(ctx, loginForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			domain.Response{false, nil, err.Error()},
		)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString, "message": "login success"})
}
