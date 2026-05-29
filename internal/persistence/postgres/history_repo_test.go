package postgres

import (
	"testing"

	"aeropath/internal/domain"
)

func TestHistoryRepo(t *testing.T) {
	pool := getTestPool(t)
	defer pool.Close()

	studentRepo := NewStudentRepo(pool.Pool)
	questionRepo := NewQuestionRepo(pool.Pool)
	historyRepo := NewHistoryRepo(pool.Pool)

	// Créer un étudiant et une question pour les tests
	student := &domain.Student{Email: "history-" + randString(6) + "@test.com", PasswordHash: "pass", Lang: "fr"}
	_ = studentRepo.Create(student)

	question := &domain.Question{
		Theme:      "test",
		Difficulty: 1,
		QuestionFr: "Test?",
		QuestionEn: "Test?",
		Options:    []string{"A", "B"},
		AnswerKey:  "A",
	}
	_ = questionRepo.Create(question)

	t.Run("enregistrer une réponse correcte", func(t *testing.T) {
		err := historyRepo.RecordAnswer(student.ID, question.ID, true)
		if err != nil {
			t.Fatalf("RecordAnswer a échoué: %v", err)
		}
	})

	t.Run("enregistrer une réponse incorrecte", func(t *testing.T) {
		err := historyRepo.RecordAnswer(student.ID, question.ID, false)
		if err != nil {
			t.Fatalf("RecordAnswer a échoué: %v", err)
		}
	})

	t.Run("mettre à jour une réponse existante (upsert)", func(t *testing.T) {
		// Déjà enregistré comme false, on met à jour en true
		err := historyRepo.RecordAnswer(student.ID, question.ID, true)
		if err != nil {
			t.Fatalf("RecordAnswer (upsert) a échoué: %v", err)
		}
	})

	t.Run("récupérer l'historique", func(t *testing.T) {
		history, err := historyRepo.GetHistory(student.ID)
		if err != nil {
			t.Fatalf("GetHistory a échoué: %v", err)
		}
		if len(history) == 0 {
			t.Fatal("devrait avoir au moins 1 entrée dans l'historique")
		}
	})

	t.Run("récupérer les statistiques", func(t *testing.T) {
		stats, err := historyRepo.GetStats(student.ID)
		if err != nil {
			t.Fatalf("GetStats a échoué: %v", err)
		}
		if stats.TotalQuestions == 0 {
			t.Fatal("TotalQuestions ne devrait pas être 0")
		}
	})

	t.Run("récupérer les IDs des questions vues", func(t *testing.T) {
		ids, err := historyRepo.GetSeenQuestionIDs(student.ID)
		if err != nil {
			t.Fatalf("GetSeenQuestionIDs a échoué: %v", err)
		}
		if len(ids) == 0 {
			t.Fatal("devrait avoir au moins 1 ID de question vue")
		}
	})

	t.Run("historique vide pour un étudiant sans activité", func(t *testing.T) {
		emptyStudent := &domain.Student{Email: "empty-" + randString(6) + "@test.com", PasswordHash: "pass", Lang: "fr"}
		_ = studentRepo.Create(emptyStudent)

		history, err := historyRepo.GetHistory(emptyStudent.ID)
		if err != nil {
			t.Fatalf("GetHistory a échoué: %v", err)
		}
		if len(history) != 0 {
			t.Fatalf("devrait retourner 0 entrées, got %d", len(history))
		}
	})
}
