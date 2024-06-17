package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
	"os"
)

func ValidateIDToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
		tokenString, err := c.Cookie(c.Param("provider"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token cookie not found"})
			return
		}
		// id token 앞에 어디서 발급했는 지 확인할 수 있게 추가 정보를 줘야함.
		// ID 토큰 검증
		payload, err := idtoken.Validate(c.Request.Context(), tokenString, os.Getenv("GOOGLE_CLIENT_ID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid ID token"})
			return
		}

		// 페이로드에서 userUUID 추출
		userUUID, ok := payload.Claims["sub"].(string)
		if !ok || userUUID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "userUUID not found in token"})
			return
		}

		// userUUID를 컨텍스트에 설정
		c.Set("userUUID", userUUID)
		c.Next()
	}
}
