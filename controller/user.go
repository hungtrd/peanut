package controller

import (
	"log"
	"net/http"
	"peanut/domain"
	"peanut/pkg/jwt"
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
		Usecase: usecase.NewUserUsecase(db),
	}
}

func (c *UserController) Login(ctx *gin.Context) {
	req := domain.RequestLogin{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := c.Usecase.Login(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "login success!",
		"token":   "Bearer " + token,
	})
}

func (c *UserController) GetUsers(ctx *gin.Context) {
	users, err := c.Usecase.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": users})
}

func (c *UserController) CurrentUser(ctx *gin.Context) {
	userId, err := jwt.ExtractTokenID(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := c.Usecase.GetUser(ctx, userId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusText(http.StatusBadRequest),
			"message": "id invalid",
		})
	}
	user, err := c.Usecase.GetUser(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusText(http.StatusBadRequest),
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	user := domain.User{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		res := ctx.Error(err).SetType(gin.ErrorTypeBind)
		log.Printf("Error: %v", res)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": http.StatusText(http.StatusBadRequest),
		})
		return
	}

	err = c.Usecase.CreateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": http.StatusText(http.StatusBadRequest),
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": http.StatusText(http.StatusCreated),
	})
}
