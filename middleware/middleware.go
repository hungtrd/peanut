package middleware

import (
	"peanut/pkg/apierrors"
	"peanut/pkg/response"

	"github.com/gin-gonic/gin"
)

func HandleNoRoute(ctx *gin.Context) {
	err := apierrors.NewErrorf(apierrors.NotFound, "invalid url")
	response.Error(ctx, err)
}

func HandleNoMethod(ctx *gin.Context) {
	err := apierrors.NewErrorf(apierrors.NotFound, "invalid method")
	response.Error(ctx, err)
}

func HandleError(ctx *gin.Context) {
	ctx.Next()
	err := ctx.Errors.Last()
	if err == nil {
		return
	}
	response.Error(ctx, err)
}
