package clickhouse

import (
	"os"
	"testing"
	"time"
)

// Test d'intégration pour ClickHouse Analytics.
//
// Prérequis : ClickHouse doit tourner sur localhost:9000
// Lancement : docker compose up clickhouse -d
//
// Ces tests sont désactivés par défaut car ils nécessitent ClickHouse.
// Pour les lancer : go test -tags=integration ./internal/persistence/clickhouse/
//
// Si ClickHouse n'est pas disponible, les tests sont skip automatiquement.

func getClickHouseHost() string {
	host := os.Getenv("CLICKHOUSE_HOST")
	if host == "" {
		host = "localhost:9000"
	}
	return host
}

func skipIfClickHouseUnavailable(t *testing.T) *Analytics {
	t.Helper()

	analytics, err := NewAnalytics(getClickHouseHost())
	if err != nil {
		t.Skipf("ClickHouse non disponible (%v), skip test d'intégration", err)
	}

	if err := analytics.EnsureTable(); err != nil {
		analytics.Close()
		t.Skipf("Impossible de créer la table ClickHouse (%v), skip", err)
	}

	return analytics
}

func TestIntegrationRecordAndGetStats(t *testing.T) {
	analytics := skipIfClickHouseUnavailable(t)
	defer analytics.Close()

	studentID := "int-test-student-" + time.Now().Format("150405.000")
	questionID := "int-test-question-1"

	// Enregistrer quelques réponses
	answers := []struct {
		correct bool
	}{
		{true}, {true}, {false}, {true}, {false},
	}

	for _, a := range answers {
		err := analytics.RecordAnswer(studentID, questionID, a.correct)
		if err != nil {
			t.Fatalf("RecordAnswer a échoué: %v", err)
		}
	}

	// Attendre que ClickHouse indexe les données (async insert)
	time.Sleep(2 * time.Second)

	// Lire les stats
	stats, err := analytics.GetStudentStats(studentID)
	if err != nil {
		t.Fatalf("GetStudentStats a échoué: %v", err)
	}

	if stats.TotalAnswers != 5 {
		t.Fatalf("TotalAnswers attendu 5, got %d", stats.TotalAnswers)
	}
	if stats.CorrectAnswers != 3 {
		t.Fatalf("CorrectAnswers attendu 3, got %d", stats.CorrectAnswers)
	}
	if stats.SuccessRate != 60.0 {
		t.Fatalf("SuccessRate attendu 60.0, got %f", stats.SuccessRate)
	}
}

func TestIntegrationGetStatsEmptyStudent(t *testing.T) {
	analytics := skipIfClickHouseUnavailable(t)
	defer analytics.Close()

	studentID := "int-test-empty-" + time.Now().Format("150405.000")

	// Un étudiant sans réponses devrait retourner des stats à zéro
	stats, err := analytics.GetStudentStats(studentID)
	if err != nil {
		t.Fatalf("GetStudentStats pour étudiant vide a échoué: %v", err)
	}

	if stats.TotalAnswers != 0 {
		t.Fatalf("TotalAnswers attendu 0, got %d", stats.TotalAnswers)
	}
	if stats.CorrectAnswers != 0 {
		t.Fatalf("CorrectAnswers attendu 0, got %d", stats.CorrectAnswers)
	}
	if stats.SuccessRate != 0 {
		t.Fatalf("SuccessRate attendu 0, got %f", stats.SuccessRate)
	}
}

func TestIntegrationDailyStats(t *testing.T) {
	analytics := skipIfClickHouseUnavailable(t)
	defer analytics.Close()

	studentID := "int-test-daily-" + time.Now().Format("150405.000")

	// Enregistrer des réponses
	for i := 0; i < 10; i++ {
		err := analytics.RecordAnswer(studentID, "q-daily", i%2 == 0)
		if err != nil {
			t.Fatalf("RecordAnswer a échoué: %v", err)
		}
	}

	// Attendre l'indexation
	time.Sleep(2 * time.Second)

	// Lire les stats quotidiennes
	stats, err := analytics.GetDailyStats(studentID, 7)
	if err != nil {
		t.Fatalf("GetDailyStats a échoué: %v", err)
	}

	if len(stats) == 0 {
		t.Fatal("devrait retourner au moins 1 jour de stats")
	}

	// Vérifier le jour le plus récent
	lastDay := stats[len(stats)-1]
	if lastDay.Total != 10 {
		t.Fatalf("Total attendu 10, got %d", lastDay.Total)
	}
	if lastDay.Correct != 5 {
		t.Fatalf("Correct attendu 5, got %d", lastDay.Correct)
	}
}

func TestIntegrationMultipleStudents(t *testing.T) {
	analytics := skipIfClickHouseUnavailable(t)
	defer analytics.Close()

	ts := time.Now().Format("150405.000")
	studentA := "int-test-multi-a-" + ts
	studentB := "int-test-multi-b-" + ts

	// Student A : 3 correctes sur 4
	for i := 0; i < 4; i++ {
		analytics.RecordAnswer(studentA, "q-multi", i != 1)
	}

	// Student B : 1 correcte sur 4
	for i := 0; i < 4; i++ {
		analytics.RecordAnswer(studentB, "q-multi", i == 0)
	}

	time.Sleep(2 * time.Second)

	// Vérifier les stats de A
	statsA, err := analytics.GetStudentStats(studentA)
	if err != nil {
		t.Fatalf("GetStudentStats A a échoué: %v", err)
	}
	if statsA.CorrectAnswers != 3 {
		t.Fatalf("CorrectAnswers A attendu 3, got %d", statsA.CorrectAnswers)
	}

	// Vérifier les stats de B
	statsB, err := analytics.GetStudentStats(studentB)
	if err != nil {
		t.Fatalf("GetStudentStats B a échoué: %v", err)
	}
	if statsB.CorrectAnswers != 1 {
		t.Fatalf("CorrectAnswers B attendu 1, got %d", statsB.CorrectAnswers)
	}
}

func TestIntegrationEnsureTableIdempotent(t *testing.T) {
	analytics := skipIfClickHouseUnavailable(t)
	defer analytics.Close()

	// EnsureTable doit être appelable plusieurs fois sans erreur
	err := analytics.EnsureTable()
	if err != nil {
		t.Fatalf("EnsureTable (1er appel) a échoué: %v", err)
	}

	err = analytics.EnsureTable()
	if err != nil {
		t.Fatalf("EnsureTable (2ème appel) a échoué: %v", err)
	}
}
