package apierrors

import (
	"errors"
	"fmt"
)

type apiError struct {
	errType ErrorType
	err     error
}

// Unwrap implements the errors.Wrapper interface.
func (e *apiError) Unwrap() error {
	return e.err
}

// Error implements the error interface.
func (e *apiError) Error() string {
	msg := fmt.Sprintf("%v %s ", e.errType.HTTPCode(), e.errType.Code())
	if e.err != nil {
		msg += e.err.Error()
	}

	return msg
}

func New(errType ErrorType, err error) *apiError {
	return &apiError{
		errType: errType,
		err:     err,
	}
}

func NewErrorf(errType ErrorType, format string, a ...interface{}) *apiError {
	return New(errType, fmt.Errorf(format, a...))
}

func ErrType(err error) ErrorType {
	var apiError *apiError
	if errors.As(err, &apiError) {
		return apiError.errType
	}
	return InternalError
}

func IsErrType(err *apiError, errType ErrorType) bool {
	return err.errType == errType
}
