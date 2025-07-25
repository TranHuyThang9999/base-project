package middlewares

import (
	"demo_time_sheet_server/core/services"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type MiddlewareJwt struct {
	jwtService *services.JwtService
	user       *services.UserService
}

func NewMiddlewareJwt(jwtService *services.JwtService, user *services.UserService) *MiddlewareJwt {
	return &MiddlewareJwt{
		jwtService: jwtService,
		user:       user,
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
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		user, err := m.user.Profile(c, claims.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}
		//check user name changle password

		if user.UpdatedAt.UTC().Format(time.RFC3339) != claims.UpdatedAccountUser.UTC().Format(time.RFC3339) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User information has been updated. Please log in again."})
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
