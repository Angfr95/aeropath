package ws

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Temps maximum pour écrire un message
	writeWait = 10 * time.Second

	// Temps maximum pour lire un message (pong)
	pongWait = 60 * time.Second

	// Intervalle d'envoi des pings
	pingPeriod = (pongWait * 9) / 10

	// Taille maximum d'un message (1 MB)
	maxMessageSize = 1048576
)

// Client représente une connexion WebSocket.
type Client struct {
	Hub       *Hub
	Conn      *websocket.Conn
	Send      chan []byte
	StudentID string
}

// NewClient crée un nouveau Client.
func NewClient(hub *Hub, conn *websocket.Conn, studentID string) *Client {
	return &Client{
		Hub:       hub,
		Conn:      conn,
		Send:      make(chan []byte, 256),
		StudentID: studentID,
	}
}

// ReadPump lit les messages depuis la connexion WebSocket.
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				log.Printf("[ws] erreur lecture: %v", err)
			}
			break
		}

		// Traiter le message
		c.handleMessage(message)
	}
}

// WritePump écrit les messages vers la connexion WebSocket.
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Vider le buffer
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write([]byte("\n"))
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// SendJSON envoie un message JSON au client.
func (c *Client) SendJSON(msg Message) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("[ws] erreur marshal: %v", err)
		return
	}
	select {
	case c.Send <- data:
	default:
		log.Printf("[ws] buffer plein pour %s, message ignoré", c.StudentID)
	}
}

// handleMessage traite un message reçu du client.
func (c *Client) handleMessage(raw []byte) {
	var msg Message
	if err := json.Unmarshal(raw, &msg); err != nil {
		c.SendJSON(Message{
			Type: MsgError,
			Payload: mustMarshal(map[string]string{
				"error": "message JSON invalide",
			}),
		})
		return
	}

	switch msg.Type {
	case MsgPing:
		c.SendJSON(Message{Type: MsgPong})

	case "room:join":
		var payload struct {
			Room string `json:"room"`
		}
		if err := json.Unmarshal(msg.Payload, &payload); err == nil && payload.Room != "" {
			c.Hub.JoinRoom(c.StudentID, payload.Room)
			c.SendJSON(Message{
				Type: MsgNotification,
				Payload: mustMarshal(map[string]string{
					"message": "rejoint la room: " + payload.Room,
					"room":    payload.Room,
				}),
			})
		}

	case "room:leave":
		var payload struct {
			Room string `json:"room"`
		}
		if err := json.Unmarshal(msg.Payload, &payload); err == nil && payload.Room != "" {
			c.Hub.LeaveRoom(c.StudentID, payload.Room)
		}

	default:
		c.SendJSON(Message{
			Type: MsgError,
			Payload: mustMarshal(map[string]string{
				"error": "type de message inconnu: " + string(msg.Type),
			}),
		})
	}
}
