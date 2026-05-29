package monitoring

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler retourne l'état de santé du service.
func HealthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
