package middlewares

//use 	cors "github.com/rs/cors/wrapper/gin"

type MiddlewareCors struct {
}

func NewMiddlewareCors() *MiddlewareCors {
	return &MiddlewareCors{}
}

// todo
func (u *MiddlewareCors) Cors()
