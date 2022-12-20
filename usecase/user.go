package usecase

import (
	"context"
	"fmt"
	"net/http"
	"peanut/domain"
	"peanut/pkg/crypto"
	jwtservices "peanut/pkg/jwt"
	"peanut/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserUsecase interface {
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (*domain.User, error)
	CreateUser(u domain.User) error
	Login(ctx *gin.Context, u domain.LoginForm) (string, error)
}

type userUsecase struct {
	UserRepo repository.UserRepo
}

func NewUserUsecase(db *gorm.DB) UserUsecase {
	return &userUsecase{
		UserRepo: repository.NewUserRepo(db),
	}
}

func (uc *userUsecase) GetUsers(ctx context.Context) (users []domain.User, err error) {
	return
}

func (uc *userUsecase) GetUser(ctx context.Context, id int) (user *domain.User, err error) {
	return
}

func (uc *userUsecase) CreateUser(u domain.User) (err error) {
	//hash password
	u.Password, err = crypto.HashString(u.Password)
	if err != nil {
		return err
	}
	//create user
	_, err = uc.UserRepo.CreateUser(u)
	if err != nil {
		return err
	}
	return
}

func (uc *userUsecase) Login(ctx *gin.Context, u domain.LoginForm) (string, error) {
	//lookup user
	user, err := uc.UserRepo.GetUserByUsername(u.Username)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	//compare password
	ok := crypto.DoMatch(user.Password, u.Password)
	if !ok {
		err = fmt.Errorf("Incorrect email or password")
		return "", err
	}
	//generate jwt token
	tokenString, err := jwtservices.GenerateToken(user)
	if err != nil {
		err = fmt.Errorf("Incorrect email or password")
		return "", err
	}
	//set cookie Authorization
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*60, "", "", false, true)
	//return
	return tokenString, err
}
