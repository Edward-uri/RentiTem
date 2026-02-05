package middleware

import (
	"net/http"
	"strings"

	"main/src/features/auth/infraestructure/services"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(jwtSvc *services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		tokenStr := header
		if strings.HasPrefix(header, "Bearer ") {
			tokenStr = strings.TrimPrefix(header, "Bearer ")
		}

		userID, email, role, err := jwtSvc.Validate(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("userID", userID)
		c.Set("userEmail", email)
		c.Set("userRole", role)
		c.Next()
	}
}
