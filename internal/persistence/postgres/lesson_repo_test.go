package postgres

import (
	"testing"

	"aeropath/internal/domain"
)

func TestLessonRepo(t *testing.T) {
	pool := getTestPool(t)
	defer pool.Close()

	repo := NewLessonRepo(pool.Pool)

	t.Run("créer une leçon", func(t *testing.T) {
		lesson := &domain.Lesson{
			Theme:      "meteorology",
			TitleFr:    "Les nuages",
			TitleEn:    "Clouds",
			ContentFr:  "Les nuages sont formés par...",
			ContentEn:  "Clouds are formed by...",
			Difficulty: 1,
			OrderIndex: 1,
		}

		err := repo.Create(lesson)
		if err != nil {
			t.Fatalf("Create a échoué: %v", err)
		}
		if lesson.ID == "" {
			t.Fatal("ID ne devrait pas être vide après création")
		}
	})

	t.Run("trouver par ID", func(t *testing.T) {
		lesson := &domain.Lesson{
			Theme:      "navigation",
			TitleFr:    "Le VOR",
			TitleEn:    "VOR Navigation",
			ContentFr:  "Le VOR est un système de...",
			ContentEn:  "VOR is a navigation system...",
			Difficulty: 2,
			OrderIndex: 1,
		}
		_ = repo.Create(lesson)

		found, err := repo.FindByID(lesson.ID)
		if err != nil {
			t.Fatalf("FindByID a échoué: %v", err)
		}
		if found.ID != lesson.ID {
			t.Fatalf("ID attendu '%s', got '%s'", lesson.ID, found.ID)
		}
		if found.TitleFr != lesson.TitleFr {
			t.Fatalf("TitleFr attendu '%s', got '%s'", lesson.TitleFr, found.TitleFr)
		}
	})

	t.Run("trouver par ID inexistant", func(t *testing.T) {
		_, err := repo.FindByID("00000000-0000-0000-0000-000000000000")
		if err == nil {
			t.Fatal("devrait retourner une erreur pour ID inexistant")
		}
	})

	t.Run("trouver par thème", func(t *testing.T) {
		theme := "theme_" + randString(4)
		for i := 0; i < 2; i++ {
			lesson := &domain.Lesson{
				Theme:      theme,
				TitleFr:    "Leçon " + string(rune('A'+i)),
				TitleEn:    "Lesson " + string(rune('A'+i)),
				ContentFr:  "Contenu " + string(rune('A'+i)),
				ContentEn:  "Content " + string(rune('A'+i)),
				Difficulty: 1,
				OrderIndex: i,
			}
			_ = repo.Create(lesson)
		}

		lessons, err := repo.FindByTheme(theme)
		if err != nil {
			t.Fatalf("FindByTheme a échoué: %v", err)
		}
		if len(lessons) != 2 {
			t.Fatalf("devrait retourner 2 leçons, got %d", len(lessons))
		}
	})

	t.Run("trouver toutes les leçons", func(t *testing.T) {
		lessons, err := repo.FindAll()
		if err != nil {
			t.Fatalf("FindAll a échoué: %v", err)
		}
		if len(lessons) < 3 {
			t.Fatalf("devrait avoir au moins 3 leçons, got %d", len(lessons))
		}
	})

	t.Run("ordre des leçons par thème", func(t *testing.T) {
		theme := "order_" + randString(4)
		// Créer dans le désordre
		lessons := []*domain.Lesson{
			{Theme: theme, TitleFr: "B", TitleEn: "B", ContentFr: "B", ContentEn: "B", Difficulty: 1, OrderIndex: 2},
			{Theme: theme, TitleFr: "A", TitleEn: "A", ContentFr: "A", ContentEn: "A", Difficulty: 1, OrderIndex: 1},
		}
		for _, l := range lessons {
			_ = repo.Create(l)
		}

		found, _ := repo.FindByTheme(theme)
		if len(found) == 2 {
			if found[0].OrderIndex > found[1].OrderIndex {
				t.Fatal("les leçons devraient être triées par order_index")
			}
		}
	})
}
