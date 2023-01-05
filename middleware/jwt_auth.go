package middleware

import (
	"peanut/pkg/apierrors"
	"peanut/pkg/jwt"
	"peanut/pkg/response"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := jwt.TokenValid(ctx)
		if err != nil {
			err = apierrors.NewErrorf(apierrors.Unauthorized, "Unauthorized")
			response.Error(ctx, err)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
