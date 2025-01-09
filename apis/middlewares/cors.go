package middlewares

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type MiddlewareCors struct {
	cors.Options
}

func NewMiddlewareCors() *MiddlewareCors {
	return &MiddlewareCors{}
}

func (u *MiddlewareCors) Cors() gin.HandlerFunc {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	})
}
