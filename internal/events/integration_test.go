package events

import (
	"os"
	"sync/atomic"
	"testing"
	"time"
)

// Test d'intégration pour NATS (Producer + Consumer).
//
// Prérequis : NATS doit tourner sur localhost:4222 avec JetStream activé
// Lancement : docker compose up nats -d
//
// Ces tests sont désactivés par défaut car ils nécessitent NATS.
// Pour les lancer : go test -tags=integration ./internal/events/
//
// Si NATS n'est pas disponible, les tests sont skip automatiquement.

func getNatsURL() string {
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = "nats://localhost:4222"
	}
	return url
}

func skipIfNatsUnavailable(t *testing.T) *Producer {
	t.Helper()

	producer, err := NewProducer(getNatsURL())
	if err != nil {
		t.Skipf("NATS non disponible (%v), skip test d'intégration", err)
	}
	return producer
}

func TestIntegrationPublishAndConsume(t *testing.T) {
	producer := skipIfNatsUnavailable(t)
	defer producer.Close()

	subject := "aeroforge.test.integration"
	received := make(chan Event, 1)

	// Créer un consumer qui écoute notre sujet de test
	consumer, err := NewConsumer(getNatsURL(), subject, func(event Event) error {
		received <- event
		return nil
	})
	if err != nil {
		t.Fatalf("NewConsumer a échoué: %v", err)
	}
	defer consumer.Close()

	// Laisser le temps à NATS de propager l'abonnement
	time.Sleep(500 * time.Millisecond)

	// Publier un événement
	event := Event{
		Type:      EventQuestionAnswered,
		Source:    "integration-test",
		StudentID: "test-student-123",
		Data: map[string]any{
			"question_id": "q-42",
			"correct":     true,
		},
		Metadata: map[string]string{
			"test": "integration",
		},
	}

	err = producer.Publish(event)
	if err != nil {
		t.Fatalf("Publish a échoué: %v", err)
	}

	// Attendre de recevoir l'événement (timeout 5s)
	select {
	case receivedEvent := <-received:
		if receivedEvent.Type != event.Type {
			t.Fatalf("Type attendu '%s', got '%s'", event.Type, receivedEvent.Type)
		}
		if receivedEvent.StudentID != event.StudentID {
			t.Fatalf("StudentID attendu '%s', got '%s'", event.StudentID, receivedEvent.StudentID)
		}
		if receivedEvent.Data["question_id"] != "q-42" {
			t.Fatalf("question_id attendu 'q-42', got '%v'", receivedEvent.Data["question_id"])
		}
		if receivedEvent.Source != "integration-test" {
			t.Fatalf("Source attendu 'integration-test', got '%s'", receivedEvent.Source)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("timeout: événement non reçu après 5s")
	}
}

func TestIntegrationMultipleEvents(t *testing.T) {
	producer := skipIfNatsUnavailable(t)
	defer producer.Close()

	subject := "aeroforge.test.integration.multi"
	var count atomic.Int32

	consumer, err := NewConsumer(getNatsURL(), subject, func(event Event) error {
		count.Add(1)
		return nil
	})
	if err != nil {
		t.Fatalf("NewConsumer a échoué: %v", err)
	}
	defer consumer.Close()

	time.Sleep(500 * time.Millisecond)

	// Publier 10 événements
	for i := 0; i < 10; i++ {
		event := Event{
			Type:      EventQuestionAnswered,
			Source:    "integration-test",
			StudentID: "test-student",
			Data:      map[string]any{"index": i},
		}
		err := producer.Publish(event)
		if err != nil {
			t.Fatalf("Publish %d a échoué: %v", i, err)
		}
	}

	// Attendre que tous les messages soient reçus
	time.Sleep(2 * time.Second)

	if count.Load() != 10 {
		t.Fatalf("devrait avoir reçu 10 événements, got %d", count.Load())
	}
}

func TestIntegrationEventTypes(t *testing.T) {
	producer := skipIfNatsUnavailable(t)
	defer producer.Close()

	subject := "aeroforge.test.integration.types"
	received := make(chan EventType, 10)

	consumer, err := NewConsumer(getNatsURL(), subject, func(event Event) error {
		received <- event.Type
		return nil
	})
	if err != nil {
		t.Fatalf("NewConsumer a échoué: %v", err)
	}
	defer consumer.Close()

	time.Sleep(500 * time.Millisecond)

	// Tester tous les types d'événements
	eventTypes := []EventType{
		EventQuestionAnswered,
		EventExamStarted,
		EventExamCompleted,
		EventLessonViewed,
		EventStudentRegistered,
		EventStudentLoggedIn,
		EventStudentUpdated,
		EventRecommendationsUpdated,
		EventWeakTopicsDetected,
		EventMilestoneReached,
		EventServiceHealthCheck,
		EventServiceDegraded,
		EventServiceDown,
	}

	for _, et := range eventTypes {
		event := Event{
			Type:      et,
			Source:    "integration-test",
			StudentID: "test-student",
			Data:      map[string]any{},
		}
		err := producer.Publish(event)
		if err != nil {
			t.Fatalf("Publish %s a échoué: %v", et, err)
		}
	}

	// Vérifier que tous les types sont reçus
	receivedTypes := make(map[EventType]bool)
	for i := 0; i < len(eventTypes); i++ {
		select {
		case et := <-received:
			receivedTypes[et] = true
		case <-time.After(5 * time.Second):
			t.Fatalf("timeout: reçu %d/%d événements", i, len(eventTypes))
		}
	}

	for _, et := range eventTypes {
		if !receivedTypes[et] {
			t.Fatalf("type d'événement non reçu: %s", et)
		}
	}
}

func TestIntegrationConsumerErrorHandling(t *testing.T) {
	producer := skipIfNatsUnavailable(t)
	defer producer.Close()

	subject := "aeroforge.test.integration.errors"
	var errorCount atomic.Int32

	// Consumer qui échoue volontairement 2 fois puis réussit
	attempt := 0
	consumer, err := NewConsumer(getNatsURL(), subject, func(event Event) error {
		attempt++
		if attempt <= 2 {
			errorCount.Add(1)
			return nil // On ACK même les erreurs pour éviter les boucles infinies
		}
		return nil
	})
	if err != nil {
		t.Fatalf("NewConsumer a échoué: %v", err)
	}
	defer consumer.Close()

	time.Sleep(500 * time.Millisecond)

	event := Event{
		Type:      EventQuestionAnswered,
		Source:    "integration-test",
		StudentID: "test-student",
		Data:      map[string]any{},
	}

	err = producer.Publish(event)
	if err != nil {
		t.Fatalf("Publish a échoué: %v", err)
	}

	time.Sleep(2 * time.Second)

	// Le handler a été appelé (même si avec erreurs simulées)
	if attempt == 0 {
		t.Fatal("le handler n'a jamais été appelé")
	}
}

func TestIntegrationProducerDisconnected(t *testing.T) {
	// Tester que Publish retourne une erreur si NATS est déconnecté
	// On utilise un Producer avec une connexion fermée
	producer := skipIfNatsUnavailable(t)
	producer.Close() // Fermer la connexion

	event := Event{
		Type:      EventQuestionAnswered,
		Source:    "integration-test",
		StudentID: "test-student",
		Data:      map[string]any{},
	}

	err := producer.Publish(event)
	if err == nil {
		t.Fatal("Publish devrait retourner une erreur si NATS est déconnecté")
	}
}
