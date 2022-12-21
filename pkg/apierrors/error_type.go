package apierrors

type ErrorType int

const (
	Unauthorized ErrorType = iota
	DisallowReregistration
	BadParams
	InvalidRequest
	NotFound
	UniqueViolation
	DatabaseError
	InternalError

	UsernameExisted
	LoginFailed
)

var errorHTTPCode = map[ErrorType]int{
	Unauthorized:           401,
	DisallowReregistration: 403,
	BadParams:              400,
	InvalidRequest:         400,
	NotFound:               404,
	UniqueViolation:        409,
	DatabaseError:          500,
	InternalError:          500,

	UsernameExisted: 400,
	LoginFailed:     400,
}

var errorCode = map[ErrorType]string{
	Unauthorized:           "UNAUTHORIZED",
	DisallowReregistration: "DISALLOW_REREGISTRATION",
	BadParams:              "BAD_PARAMS",
	InvalidRequest:         "INVALID_REQUEST",
	NotFound:               "NOT_FOUND",
	UniqueViolation:        "UNIQUE_VIOLATION",
	DatabaseError:          "DATABASE_ERROR",
	InternalError:          "INTERNAL_SERVER_ERROR",

	UsernameExisted: "USERNAME_EXISTED",
	LoginFailed:     "LOGIN_FAILED",
}

func (t ErrorType) HTTPCode() int {
	return errorHTTPCode[t]
}

func (t ErrorType) Code() string {
	return errorCode[t]
}
