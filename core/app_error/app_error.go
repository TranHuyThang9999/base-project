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
)

var (
	ErrorDataBase       = errors.New("database error")
	ErrorAuthentication = errors.New("authentication error")
	ErrorNotFound       = errors.New("not found error")
	ErrorUserExists     = errors.New("user already exists")
	ErrorHashPassword   = errors.New("error hash password")
)
