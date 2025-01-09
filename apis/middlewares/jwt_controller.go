package middlewares

import (
	"net/http"
	"rices/core/services"
	"strings"

	"github.com/gin-gonic/gin"
)

type MiddlewareJwt struct {
	jwtService *services.JwtService
}

func NewMiddlewareJwt(jwtService *services.JwtService) *MiddlewareJwt {
	return &MiddlewareJwt{
		jwtService: jwtService,
	}
}

func (m *MiddlewareJwt) Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Extract the token from the Authorization header
		// Format: "Bearer <token>"
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format. Expected Bearer <token>"})
			c.Abort()
			return
		}

		tokenString := bearerToken[1]

		// Verify the token
		claims, err := m.jwtService.VerifyToken(c, tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Store user information in the context for later use
		c.Set("userId", claims.Id)
		c.Set("userName", claims.UserName)
		c.Set("updatedAt", claims.UpdatedAccountUser)

		c.Next()
	}
}
