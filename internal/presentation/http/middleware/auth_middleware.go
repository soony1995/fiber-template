package middleware

import (
	"login_module/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenValid() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no authentication token provided"})
			c.Abort()
			return
		}

		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}

		claims, err := jwt.VerifyJWTToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("username", claims.(jwt.CustomClaims).Username)
		c.Next()
	}
}
