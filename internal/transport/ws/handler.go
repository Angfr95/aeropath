package ws

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Permettre toutes les origines en dev
	},
}

// ServeWS gère l'upgrade HTTP → WebSocket.
// Le studentID est extrait du JWT par le middleware RequireAuth.
func ServeWS(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		studentID := c.GetString("student_id")
		if studentID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "non authentifié"})
			return
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("[ws] upgrade échoué: %v", err)
			return
		}

		client := NewClient(hub, conn, studentID)
		hub.register <- client

		// Lancer les goroutines de lecture/écriture
		go client.WritePump()
		go client.ReadPump()
	}
}

// WsInfoHandler retourne les stats du hub WebSocket.
func WsInfoHandler(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"connected_clients": hub.ConnectedCount(),
			"status":            "running",
		})
	}
}
