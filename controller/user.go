package controller

import (
	"peanut/domain"
	"peanut/pkg/response"
	"peanut/repository"
	"peanut/usecase"

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

// SignUp godoc
//
//	@Summary		SignUp
//	@Description	Sign Up
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			param	body	domain.SignupReq	true	"Sign up request"
//	@Success		200	{object}	domain.Response{data=domain.SignupResp}
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		500	{object}	domain.ErrorResponse
//	@Router			/v1/users/signup [post]
func (c *UserController) SignUp(ctx *gin.Context) {
	req := domain.SignupReq{}
	if !bindJSON(ctx, &req) {
		return
	}

	resp, err := c.Usecase.SignUp(ctx, req)
	if checkError(ctx, err) {
		return
	}

	response.OK(ctx, resp)
}

// Login godoc
//
//	@Summary		User Login
//	@Description	API Login
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			login_param	body		domain.LoginReq	true	"Login request"
//	@Success		200			{object}	domain.Response{data=domain.LoginReq}
//	@Failure		400			{object}	domain.ErrorResponse
//	@Failure		500			{object}	domain.ErrorResponse
//	@Router			/v1/users/login [post]
func (c *UserController) Login(ctx *gin.Context) {
	req := domain.LoginReq{}
	if !bindJSON(ctx, &req) {
		return
	}

	resp, err := c.Usecase.Login(ctx, req)
	if checkError(ctx, err) {
		return
	}

	response.OK(ctx, resp)
}

// GetUser godoc
//
//	@Summary		Create an user
//	@Description	Create an user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	domain.User
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		404	{object}	domain.ErrorResponse
//	@Failure		500	{object}	domain.ErrorResponse
//	@Security		Bearer
//	@Router			/v1/users/{id} [get]
func (c *UserController) GetUser(ctx *gin.Context) {

}
