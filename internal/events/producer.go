package events

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

// ============================================================
//  📤 PRODUCER — Envoyer des événements (NATS)
// ============================================================
// 📖 DDIA Chapitre 11 : "Stream Processing" (Message Queues)
//
// ❓ C'EST QUOI UN PRODUCER ?
//    Imagine que chaque action sur AeroForge (réponse à une question,
//    inscription, fin d'examen) doit être "annoncée" aux autres
//    services. Le Producer est le "micro" qui annonce ces événements.
//
//    Exemple : Quand un étudiant répond à une question :
//    1. Le Producer crie dans le micro : "Question répondue !"
//    2. NATS transmet le message à tous les intéressés
//    3. Le Consumer Analytics reçoit et met à jour les stats
//    4. Le Consumer Recommandation reçoit et ajuste le score
//
// 🧱 COMMENT ÇA MARCHE ?
//    - NATS est un "message broker" (facteur qui distribue le courrier)
//    - On publie des messages sur des "sujets" (topics)
//    - Les sujets sont organisés comme : aeroforge.learning.question.answered
//    - Les consommateurs s'abonnent aux sujets qui les intéressent
//
// 🔗 LIENS UTILES :
//    - https://nats.io/documentation/
//    - github.com/nats-io/nats.go
// ============================================================

// Producer publie des événements sur le bus NATS.
// C'est lui qui "parle" dans le micro pour que les autres entendent.
type Producer struct {
	nc *nats.Conn
	js nats.JetStreamContext
}

// NewProducer crée un nouveau producer connecté à NATS.
//
// 🛡️ PRODUCTION :
//    - Timeout de connexion de 5 secondes (évite de bloquer indéfiniment)
//    - Reconnexion automatique si NATS redémarre
//    - Toutes les 10 secondes, NATS vérifie que la connexion est vivante (Ping)
//    - Si la connexion est perdue, NATS réessaie toutes les secondes
//
// 📝 EXEMPLE D'UTILISATION :
//    producer, err := events.NewProducer("nats://localhost:4222")
//    if err != nil { log.Fatal(err) }
//    defer producer.Close()
func NewProducer(natsURL string) (*Producer, error) {
	nc, err := nats.Connect(natsURL,
		nats.Timeout(5*time.Second),           // Timeout connexion
		nats.ReconnectWait(1*time.Second),      // Attendre 1s entre chaque reconnexion
		nats.MaxReconnects(-1),                 // Reconnexion infinie (-1 = illimité)
		nats.PingInterval(10*time.Second),      // Ping toutes les 10s
	)
	if err != nil {
		return nil, fmt.Errorf("connexion NATS échouée: %w", err)
	}

	js, err := nc.JetStream()
	if err != nil {
		nc.Close()
		return nil, fmt.Errorf("JetStream échoué: %w", err)
	}

	log.Println("📤 Producer NATS connecté")
	return &Producer{nc: nc, js: js}, nil
}

// Publish publie un événement sur le bus NATS.
//
// 🛡️ PRODUCTION :
//    - Timeout de publication de 5 secondes
//    - Si NATS est déconnecté, la publication échoue immédiatement
//      (pas de file d'attente locale qui pourrait perdre des messages)
//
// 🧠 COMMENT CHOISIR LE SUJET (subject) ?
//    On utilise une convention : aeroforge.<domaine>.<action>
//    Exemples :
//    - aeroforge.learning.question.answered
//    - aeroforge.student.registered
//    - aeroforge.recommendation.updated
func (p *Producer) Publish(event Event) error {
	if !p.nc.IsConnected() {
		return fmt.Errorf("NATS déconnecté, message non publié: %s", event.Type)
	}

	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("sérialisation échouée: %w", err)
	}

	subject := "aeroforge." + string(event.Type)

	_, err = p.js.Publish(subject, data)
	if err != nil {
		return fmt.Errorf("publication échouée: %w", err)
	}

	log.Printf("📤 Événement publié: %s", subject)
	return nil
}

// Close ferme la connexion au broker.
func (p *Producer) Close() error {
	if p.nc != nil {
		p.nc.Close()
	}
	return nil
}

// ============================================================
//  🧪 TESTER LE PRODUCER
// ============================================================
// 1. Démarre NATS : docker compose up nats -d
// 2. Écoute les messages : nats sub "aeroforge.>"
// 3. Lance ton code : go run cmd/worker/main.go
// 4. Tu devrais voir les messages apparaître dans le terminal
// ============================================================
