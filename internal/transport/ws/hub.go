package ws

import (
	"encoding/json"
	"log"
	"sync"
)

// 📖 DDIA Chapitre 9 : "Consistency and Consensus"
//    Le WebSocket Hub est un "pub-sub" pattern :
//    - Les clients s'abonnent à un sujet (ex: "student:123")
//    - Le hub distribue les messages aux abonnés
//    - C'est un système "at-most-once" : pas de garantie de delivery
//
//    Utilisation : notifications en temps réel
//    - Résultat d'examen
//    - Mise à jour des recommandations
//    - Alertes de progression

// MessageType représente le type d'un message WebSocket
type MessageType string

const (
	MsgPing         MessageType = "ping"
	MsgPong         MessageType = "pong"
	MsgNotification MessageType = "notification"
	MsgError        MessageType = "error"
)

// Message représente un message WebSocket structuré
type Message struct {
	Type    MessageType     `json:"type"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

// mustMarshal marshale une valeur en JSON ou panique
func mustMarshal(v interface{}) json.RawMessage {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return data
}

// Hub gère l'ensemble des connexions WebSocket
type Hub struct {
	mu         sync.RWMutex
	clients    map[*Client]bool
	rooms      map[string]map[string]bool // room -> studentID -> true
	register   chan *Client
	unregister chan *Client
}

// NewHub crée un nouveau Hub WebSocket
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		rooms:      make(map[string]map[string]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// Run démarre la boucle principale du hub
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()
			log.Printf("🟢 WebSocket client connecté: %s", client.StudentID)

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}
			h.mu.Unlock()
			log.Printf("🔴 WebSocket client déconnecté: %s", client.StudentID)
		}
	}
}

// ConnectedCount retourne le nombre de clients connectés
func (h *Hub) ConnectedCount() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.clients)
}

// JoinRoom abonne un étudiant à une room
func (h *Hub) JoinRoom(studentID string, room string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.rooms[room] == nil {
		h.rooms[room] = make(map[string]bool)
	}
	h.rooms[room][studentID] = true
}

// LeaveRoom désabonne un étudiant d'une room
func (h *Hub) LeaveRoom(studentID string, room string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if students, ok := h.rooms[room]; ok {
		delete(students, studentID)
		if len(students) == 0 {
			delete(h.rooms, room)
		}
	}
}

// BroadcastToRoom envoie un message à tous les clients d'une room
func (h *Hub) BroadcastToRoom(room string, message interface{}) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Printf("❌ Erreur de marshaling WebSocket: %v", err)
		return
	}

	h.mu.RLock()
	students := h.rooms[room]
	clients := make([]*Client, 0)
	for c := range h.clients {
		if students[c.StudentID] {
			clients = append(clients, c)
		}
	}
	h.mu.RUnlock()

	for _, client := range clients {
		select {
		case client.Send <- data:
		default:
			// Client trop lent, on ferme la connexion
			h.mu.Lock()
			delete(h.clients, client)
			close(client.Send)
			h.mu.Unlock()
		}
	}
}

// BroadcastToUser envoie un message à un utilisateur spécifique
func (h *Hub) BroadcastToUser(userID string, message interface{}) {
	h.BroadcastToRoom("user:"+userID, message)
}

// FindClientByStudentID retourne le premier client trouvé pour un étudiant
func (h *Hub) FindClientByStudentID(studentID string) *Client {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for c := range h.clients {
		if c.StudentID == studentID {
			return c
		}
	}
	return nil
}
