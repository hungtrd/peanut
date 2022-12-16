package response

import (
	"errors"
	"peanut/domain"
	"peanut/pkg/apierrors"
	"peanut/pkg/i18n"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Error(ctx *gin.Context, err error) {
	var debugMsg string
	if gin.Mode() != gin.ReleaseMode {
		debugMsg = err.Error()
	}
	errType := apierrors.ErrType(err)

	ctx.JSON(errType.HTTPCode(), domain.ErrorResponse{
		Code:         errType.Code(),
		DebugMessage: debugMsg,
		ErrorDetails: errorDetails(err, ctx.Query("locale")),
	})
}

func errorDetails(err error, locale string) (details []domain.ErrorDetail) {
	var vErrs validator.ValidationErrors
	if errors.As(err, &vErrs) {
		trans := i18n.GetTrans(locale)
		for _, err := range vErrs {
			details = append(details, domain.ErrorDetail{
				Field:        err.Field(),
				ErrorCode:    err.Tag(),
				ErrorMessage: err.Translate(trans),
			})
		}
	}
	return
}
