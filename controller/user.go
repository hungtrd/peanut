package controller

import (
	user2 "golang_first_pj/domain/request/user"
	"golang_first_pj/usecase"
	"log"
	"net/http"

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
	user := user2.CreateUserReq{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		res := ctx.Error(err).SetType(gin.ErrorTypeBind)
		log.Printf("Error: %v", res)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": http.StatusText(http.StatusBadRequest),
		})
		return
	}
	c.Usecase.CreateUser(ctx, user)
	ctx.JSON(http.StatusCreated, gin.H{
		"message": http.StatusText(http.StatusCreated),
	})
}
