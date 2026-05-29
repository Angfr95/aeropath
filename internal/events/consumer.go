package events

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

// ============================================================
//  📥 CONSUMER — Recevoir et traiter les événements (NATS)
// ============================================================
// 📖 DDIA Chapitre 11 : "Stream Processing" (Consumers)
//
// ❓ C'EST QUOI UN CONSUMER ?
//    Si le Producer est le "micro", le Consumer est le "haut-parleur".
//    Il écoute les événements et exécute des actions en réponse.
//
//    Exemple : Quand un étudiant répond à une question :
//    1. Le Producer publie "question.answered"
//    2. NATS distribue le message à TOUS les consommateurs abonnés
//    3. Le Consumer Analytics reçoit → met à jour les stats ClickHouse
//    4. Le Consumer Recommandation reçoit → recalcule le score de maîtrise
//    5. Le Consumer Historique reçoit → enregistre dans PostgreSQL
//
// 🧱 COMMENT ÇA MARCHE ?
//    - On s'abonne à un sujet NATS (ex: "aeroforge.learning.>")
//    - NATS nous envoie tous les messages de ce sujet
//    - On traite le message (handler)
//    - On accuse réception (ACK) pour dire "traitement OK"
//    - Si on ne ACK pas, NATS renvoie le message (retry)
//
// 🔗 LIENS UTILES :
//    - https://nats.io/documentation/concepts/jetstream/consumers/
//    - github.com/nats-io/nats.go
// ============================================================

// Consumer reçoit et traite les événements du bus NATS.
// C'est lui qui "écoute" ce qui se passe et réagit.
type Consumer struct {
	nc      *nats.Conn
	js      nats.JetStreamContext
	sub     *nats.Subscription
	handler EventHandler
}

// EventHandler est une fonction qui traite un événement reçu.
// Retourne une erreur si le traitement a échoué (NATS renverra le message).
type EventHandler func(event Event) error

// NewConsumer crée un nouveau consumer connecté à NATS.
//
// 🛡️ PRODUCTION :
//    - Reconnexion automatique si NATS redémarre
//    - Abonnement durable (Durable) : si le consumer plante,
//      NATS garde les messages jusqu'à son retour
//    - ACK manuel : on décide quand le message est traité
//    - Panic recovery : si le handler panique, on catch et on NAK
//
// 📝 EXEMPLE D'UTILISATION :
//    consumer, err := events.NewConsumer("nats://localhost:4222",
//        "aeroforge.learning.>",
//        func(event events.Event) error {
//            fmt.Printf("Reçu: %+v\n", event)
//            return nil
//        })
func NewConsumer(natsURL, subject string, handler EventHandler) (*Consumer, error) {
	nc, err := nats.Connect(natsURL,
		nats.Timeout(5*time.Second),
		nats.ReconnectWait(1*time.Second),
		nats.MaxReconnects(-1),
		nats.PingInterval(10*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("connexion NATS échouée: %w", err)
	}

	js, err := nc.JetStream()
	if err != nil {
		nc.Close()
		return nil, fmt.Errorf("JetStream échoué: %w", err)
	}

	// S'abonner au sujet avec JetStream (persistant)
	// JetStream garantit que les messages ne sont pas perdus
	// même si le consumer redémarre.
	sub, err := js.Subscribe(subject, func(msg *nats.Msg) {
		// 🛡️ PROTECTION CONTRE LES PANICS
		// Si le handler fait panic (ex: nil pointer, division par zéro),
		// on ne veut pas que tout le programme crashe.
		// On récupère la panic et on NAK le message pour retry.
		defer func() {
			if r := recover(); r != nil {
				log.Printf("❌ PANIC dans le handler: %v", r)
				msg.Nak() // Nak = NATS renverra le message
			}
		}()

		var event Event
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Printf("❌ Erreur désérialisation: %v", err)
			msg.Ack() // On ACK même en erreur pour éviter les boucles
			return
		}

		log.Printf("📥 Événement reçu: %s", event.Type)

		if err := handler(event); err != nil {
			log.Printf("❌ Erreur traitement %s: %v", event.Type, err)
			msg.Nak() // Nak = Negative ACK → NATS renverra le message
			return
		}

		msg.Ack() // ACK = tout s'est bien passé
	}, nats.Durable("aeroforge-worker"), nats.ManualAck())

	if err != nil {
		nc.Close()
		return nil, fmt.Errorf("abonnement échoué: %w", err)
	}

	log.Printf("📥 Consumer NATS abonné à: %s", subject)
	return &Consumer{nc: nc, js: js, sub: sub, handler: handler}, nil
}

// Close ferme la connexion au broker.
// 🛡️ PRODUCTION : on ferme d'abord l'abonnement, puis la connexion.
func (c *Consumer) Close() error {
	if c.sub != nil {
		if err := c.sub.Unsubscribe(); err != nil {
			return fmt.Errorf("désabonnement échoué: %w", err)
		}
	}
	if c.nc != nil {
		c.nc.Close()
	}
	return nil
}

// ============================================================
//  🧪 TESTER LE CONSUMER
// ============================================================
// 1. Démarre NATS : docker compose up nats -d
// 2. Publie un message : nats pub "aeroforge.learning.question.answered" '{"type":"question.answered","student_id":"123"}'
// 3. Lance le worker : go run cmd/worker/main.go
// 4. Tu devrais voir "📥 Événement reçu: question.answered"
// ============================================================
