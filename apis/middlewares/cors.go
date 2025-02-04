package middlewares

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type MiddlewareCors struct {
	cors.Options
}

func NewMiddlewareCors() *MiddlewareCors {
	return &MiddlewareCors{}
}

func (u *MiddlewareCors) CorsAPI() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		timeoutCtx, cancel := context.WithTimeout(ctx.Request.Context(), 15*time.Second)
		defer cancel()

		ctx.Request = ctx.Request.WithContext(timeoutCtx)

		cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Origin", "Content-Type", "Accept"},
			ExposedHeaders:   []string{"Content-Length"},
			AllowCredentials: true,
		})(ctx)

		select {
		case <-timeoutCtx.Done():
			if timeoutCtx.Err() == context.DeadlineExceeded {
				ctx.JSON(408, gin.H{"error": "Request Timeout"})
				ctx.Abort()
			}
		default:
			ctx.Next()
		}
	}
}

func (u *MiddlewareCors) CorsWss() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
