package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminLookup interface {
	IsAdmin(id string) (bool, error)
}

func RequireAdmin(repo AdminLookup) gin.HandlerFunc {
	return authorize(repo, false)
}

func RequireOwnerOrAdmin(repo AdminLookup) gin.HandlerFunc {
	return authorize(repo, true)
}

func authorize(repo AdminLookup, allowOwner bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetString("student_id")
		if id == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
			return
		}
		if allowOwner && id == c.Param("id") {
			c.Next()
			return
		}
		admin, err := repo.IsAdmin(id)
		if err != nil || !admin {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "access denied"})
			return
		}
		c.Next()
	}
}
