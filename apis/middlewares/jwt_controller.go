package middlewares

import "github.com/gin-gonic/gin"
//use import "github.com/golang-jwt/jwt/v5"
type MiddlewareJwt struct {
}

func NewMiddlewareJwt() *MiddlewareJwt {
	return &MiddlewareJwt{}
}
//todo
func (u *MiddlewareJwt) Authorization() gin.HandlerFunc(){
	
}
