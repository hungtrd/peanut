package controller

import (
	"github.com/gin-gonic/gin"
	"golang_first_pj/domain/request/auth"
	"golang_first_pj/usecase"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type AuthController struct {
	Usecase usecase.AuthUsecase
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		Usecase: usecase.NewAuthUsecase(db),
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	user := auth.Auth{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		res := ctx.Error(err).SetType(gin.ErrorTypeBind)
		log.Printf("Error: %v", res)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": http.StatusText(http.StatusBadRequest),
		})
		return
	}
	jwt, _ := c.Usecase.Login(ctx, user)

	if jwt != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"jwt": jwt,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": http.StatusText(http.StatusUnauthorized),
		})
	}
}
