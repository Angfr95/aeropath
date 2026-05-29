package postgres

import (
	"testing"

	"aeropath/internal/domain"
)

func TestQuestionRepo(t *testing.T) {
	pool := getTestPool(t)
	defer pool.Close()

	repo := NewQuestionRepo(pool.Pool)

	t.Run("créer une question", func(t *testing.T) {
		q := &domain.Question{
			Theme:      "meteorology",
			Subtopic:   "clouds",
			Difficulty: 2,
			QuestionFr: "Quel nuage est associé au beau temps ?",
			QuestionEn: "Which cloud is associated with fair weather?",
			Options:    []string{"Cumulus", "Stratus", "Cumulonimbus", "Nimbostratus"},
			AnswerKey:  "Cumulus",
		}

		err := repo.Create(q)
		if err != nil {
			t.Fatalf("Create a échoué: %v", err)
		}
		if q.ID == "" {
			t.Fatal("ID ne devrait pas être vide après création")
		}
	})

	t.Run("trouver par ID", func(t *testing.T) {
		q := &domain.Question{
			Theme:      "navigation",
			Difficulty: 1,
			QuestionFr: "Test question",
			QuestionEn: "Test question",
			Options:    []string{"A", "B", "C"},
			AnswerKey:  "A",
		}
		repo.Create(q)

		found, err := repo.FindByID(q.ID)
		if err != nil {
			t.Fatalf("FindByID a échoué: %v", err)
		}
		if found.ID != q.ID {
			t.Fatalf("ID attendu '%s', got '%s'", q.ID, found.ID)
		}
		if found.QuestionFr != q.QuestionFr {
			t.Fatalf("QuestionFr attendu '%s', got '%s'", q.QuestionFr, found.QuestionFr)
		}
	})

	t.Run("trouver par ID inexistant", func(t *testing.T) {
		_, err := repo.FindByID("00000000-0000-0000-0000-000000000000")
		if err == nil {
			t.Fatal("devrait retourner une erreur pour ID inexistant")
		}
	})

	t.Run("trouver toutes les questions", func(t *testing.T) {
		// Créer quelques questions
		for i := 0; i < 3; i++ {
			q := &domain.Question{
				Theme:      "test",
				Difficulty: 1,
				QuestionFr: "Question " + string(rune('A'+i)),
				QuestionEn: "Question " + string(rune('A'+i)),
				Options:    []string{"A", "B"},
				AnswerKey:  "A",
			}
			repo.Create(q)
		}

		questions, err := repo.FindAll()
		if err != nil {
			t.Fatalf("FindAll a échoué: %v", err)
		}
		if len(questions) < 3 {
			t.Fatalf("devrait avoir au moins 3 questions, got %d", len(questions))
		}
	})

	t.Run("trouver par thème", func(t *testing.T) {
		theme := "theme_" + randString(4)
		for i := 0; i < 2; i++ {
			q := &domain.Question{
				Theme:      theme,
				Difficulty: 1,
				QuestionFr: "Question " + string(rune('A'+i)),
				QuestionEn: "Question " + string(rune('A'+i)),
				Options:    []string{"A", "B"},
				AnswerKey:  "A",
			}
			repo.Create(q)
		}

		questions, err := repo.FindByTheme(theme)
		if err != nil {
			t.Fatalf("FindByTheme a échoué: %v", err)
		}
		if len(questions) != 2 {
			t.Fatalf("devrait retourner 2 questions, got %d", len(questions))
		}
	})

	t.Run("options JSON correctement sérialisées/désérialisées", func(t *testing.T) {
		options := []string{"Option A", "Option B", "Option C", "Option D"}
		q := &domain.Question{
			Theme:      "test_json",
			Difficulty: 3,
			QuestionFr: "Test JSON options",
			QuestionEn: "Test JSON options",
			Options:    options,
			AnswerKey:  "Option A",
		}
		repo.Create(q)

		found, _ := repo.FindByID(q.ID)
		if len(found.Options) != 4 {
			t.Fatalf("devrait avoir 4 options, got %d", len(found.Options))
		}
		if found.Options[0] != "Option A" {
			t.Fatalf("première option attendue 'Option A', got '%s'", found.Options[0])
		}
	})
}
