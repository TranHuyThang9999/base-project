package customerrors

type CustomError struct {
	error
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewError() *CustomError {
	return &CustomError{}
}
func (u *CustomError) String() {

}
