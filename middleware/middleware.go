package middleware

import (
	"fmt"
	"net/http"
	"os"
	"peanut/pkg/apierrors"
	"peanut/pkg/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

func Authentication(c *gin.Context) {
	// get tokenstring from cookie
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//decode - validate tokenstring
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("HMAC_SECRET_SIGNING_KEY")), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//find user with the token
		userId := claims["userId"]

		//attach to req
		c.Set("userId", userId)
		// continue next controller func
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Next()
}
