package usecase

import (
	"context"
	"fmt"
	"peanut/domain"
	"peanut/pkg/apierrors"
	"peanut/pkg/crypto"
	"peanut/pkg/jwt"
	"peanut/repository"
)

type UserUsecase interface {
	SignUp(ctx context.Context, sr domain.SignupReq) (domain.SignupResp, error)
	Login(ctx context.Context, lr domain.LoginReq) (domain.LoginResp, error)
}

type userUsecase struct {
	UserRepo repository.UserRepo
}

func NewUserUsecase(r repository.UserRepo) UserUsecase {
	return &userUsecase{
		UserRepo: r,
	}
}

func (uc *userUsecase) SignUp(ctx context.Context, req domain.SignupReq) (resp domain.SignupResp, err error) {
	// Check username existed
	existed, err := uc.UserRepo.GetUserByUsername(req.Username)
	if err != nil {
		err = fmt.Errorf("[usecase.userUsecase.CreateUser] failed: %w", err)
		return
	}

	if existed != nil {
		err = apierrors.NewErrorf(apierrors.UsernameExisted, "Username existed")
		return
	}

	// Create user
	u := domain.User{
		Username:    req.Username,
		DisplayName: req.DisplayName,
		Email:       req.Email,
	}
	u.Password = crypto.HashString(req.Password)
	user, err := uc.UserRepo.CreateUser(u)
	if err != nil {
		err = fmt.Errorf("[usecase.userUsecase.CreateUser] failed: %w", err)
		return
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		err = fmt.Errorf("[usecase.userUsecase.CreateUser] failed: %w", err)
		return
	}
	resp.Token = token

	return
}

func (uc *userUsecase) Login(ctx context.Context, lr domain.LoginReq) (lrs domain.LoginResp, err error) {
	user, err := uc.UserRepo.GetUserByUsername(lr.Username)
	if err != nil {
		err = fmt.Errorf("[usercase.userUsecase.Login] failed: %w", err)
		return
	}

	if user == nil {
		err = apierrors.NewErrorf(apierrors.LoginFailed, "username or password invalid")
		return
	}

	if !crypto.DoMatch(user.Password, lr.Password) {
		err = apierrors.NewErrorf(apierrors.LoginFailed, "username or password invalid")
		return
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		err = fmt.Errorf("[usercase.userUsecase.Login] failed: %w", err)
		return
	}

	lrs.Token = token

	return
}
