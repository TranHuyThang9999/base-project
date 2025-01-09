package customerrors

import "fmt"

type CustomError struct {
	error
	Status  int    `json:"status,omitempty"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewError(err error, status, code int, message string) *CustomError {
	return &CustomError{
		error:   err,
		Status:  status,
		Code:    code,
		Message: message,
	}
}

func (u *CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", u.Code, u.Message)
}

func (u *CustomError) String() string {
	return u.Error()
}
