package controller

import (
	"github.com/google/uuid"
	"mime/multipart"
	"path/filepath"
	"peanut/pkg/apierrors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func bindJSON(ctx *gin.Context, obj interface{}) bool {
	err := ctx.ShouldBindJSON(obj)
	if err == nil {
		return true
	}
	_, ok := err.(validator.ValidationErrors)
	if ok {
		err = apierrors.New(apierrors.InvalidRequest, err)
	} else {
		err = apierrors.New(apierrors.BadParams, err)
	}
	ctx.Error(err).SetType(gin.ErrorTypeBind)

	return false
}

func bind(ctx *gin.Context, obj interface{}) bool {
	err := ctx.ShouldBind(obj)
	if err == nil {
		return true
	}
	_, ok := err.(validator.ValidationErrors)
	if ok {
		err = apierrors.New(apierrors.InvalidRequest, err)
	} else {
		err = apierrors.New(apierrors.BadParams, err)
	}
	ctx.Error(err).SetType(gin.ErrorTypeBind)

	return false
}

func bindQueryParams(ctx *gin.Context, obj interface{}) bool {
	err := ctx.ShouldBindQuery(obj)

	if err == nil {
		return true
	}
	_, ok := err.(validator.ValidationErrors)
	if ok {
		err = apierrors.New(apierrors.InvalidRequest, err)
	} else {
		err = apierrors.New(apierrors.BadParams, err)
	}
	ctx.Error(err).SetType(gin.ErrorTypeBind)

	return false
}

func checkError(ctx *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	_ = ctx.Error(err).SetType(gin.ErrorTypePublic)
	return true
}

func SaveUploadedFileTo(ctx *gin.Context, file *multipart.FileHeader, path string) (error, string) {
	filename := uuid.NewString() + "_" + filepath.Base(file.Filename)
	filePathServer := path + filename
	err := ctx.SaveUploadedFile(file, filePathServer)
	if err != nil {
		return err, ""
	}
	return nil, filePathServer
}
