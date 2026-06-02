package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Essaie d'abord le cookie httpOnly, puis le header Authorization
		tokenStr, err := c.Cookie("token")
		if err != nil {
			tokenStr = c.GetHeader("Authorization")
			if tokenStr != "" && len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
				tokenStr = tokenStr[7:]
			} else {
				c.AbortWithStatusJSON(401, gin.H{"error": "token manquant"})
				return
			}
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signature invalide")
			}
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "token invalide"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "token invalide"})
			return
		}

		c.Set("student_id", claims["student_id"])
		c.Next()
	}
}
