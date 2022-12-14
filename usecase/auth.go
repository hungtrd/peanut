package usecase

import (
	"context"
	"golang_first_pj/domain/request/auth"
	"golang_first_pj/helper"
	"golang_first_pj/repository"
	"gorm.io/gorm"
	"os"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

type AuthUsecase interface {
	Login(ctx context.Context, u auth.Auth) (*string, error)
}
type authUsecase struct {
	UserRepo repository.UserRepo
}

func NewAuthUsecase(db *gorm.DB) AuthUsecase {
	return &authUsecase{
		UserRepo: repository.NewUserRepo(db),
	}
}
func (uc *authUsecase) Login(ctx context.Context, u auth.Auth) (*string, error) {
	user, err := uc.UserRepo.GetUserByGmail(ctx, u.Email)

	if err != nil {
		return nil, err
	}
	match := helper.CheckPasswordHash(u.Password, user.Password)
	if match {
		jwt, err := helper.GenerateJWT(user)
		if err != nil {
			return nil, err
		}
		return &jwt, nil
	}
	return nil, nil
}
