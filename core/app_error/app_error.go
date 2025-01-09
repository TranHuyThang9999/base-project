package apperror

import (
	"errors"
)

const (
	ErrorDataBaseCode = iota + 1
	ErrorAuthCode
	ErrorNotFoundCode
	ErrorUserExistsCode
	ErrorHashPasswordCode
	ErrorGenTokenCode
	VerifyTokenCode
)

var (
	ErrorDataBase       = errors.New("database error")
	ErrorAuthentication = errors.New("authentication error")
	ErrorNotFound       = errors.New("not found error")
	ErrorUserExists     = errors.New("user already exists")
	ErrorHashPassword   = errors.New("error hash password")
	ErrorGenToken       = errors.New("error generating token")
	VerifyToken         = errors.New("error verify token")
)
