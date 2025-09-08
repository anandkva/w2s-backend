package middleware

import (
	"net/http"
	"strings"

	"w2s-backend/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			utils.SendError(c, http.StatusUnauthorized, "authorization header missing")
			c.Abort()
			return
		}

		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			utils.SendError(c, http.StatusUnauthorized, "invalid authorization header")
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			utils.SendError(c, http.StatusUnauthorized, "invalid or expired token")
			c.Abort()
			return
		}

		if uid, ok := claims["userId"].(string); ok {
			c.Set("userId", uid)
		} else {
			utils.SendError(c, http.StatusUnauthorized, "invalid token payload")
			c.Abort()
			return
		}

		c.Next()
	}
}
