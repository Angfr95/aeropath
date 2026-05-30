package postgres

import (
	"os"
	"testing"

	"aeropath/internal/domain"
)

// getTestPool retourne un pool de test. Si DATABASE_URL n'est pas défini, les tests sont skip.
func getTestPool(t *testing.T) *TestPool {
	t.Helper()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL non défini, skip test d'intégration")
	}

	pool, err := NewTestPool(dsn)
	if err != nil {
		t.Fatalf("NewTestPool: %v", err)
	}

	return pool
}

func TestStudentRepo(t *testing.T) {
	pool := getTestPool(t)
	defer pool.Close()

	repo := NewStudentRepo(pool.Pool)

	t.Run("créer un étudiant", func(t *testing.T) {
		student := &domain.Student{
			Email:        "test-" + randString(6) + "@test.com",
			PasswordHash: "hashed-password",
			Lang:         "fr",
		}

		err := repo.Create(student)
		if err != nil {
			t.Fatalf("Create a échoué: %v", err)
		}
		if student.ID == "" {
			t.Fatal("ID ne devrait pas être vide après création")
		}
	})

	t.Run("email dupliqué", func(t *testing.T) {
		email := "duplicate-" + randString(6) + "@test.com"
		s1 := &domain.Student{Email: email, PasswordHash: "pass", Lang: "fr"}
		s2 := &domain.Student{Email: email, PasswordHash: "pass", Lang: "fr"}

		if err := repo.Create(s1); err != nil {
			t.Fatalf("première création a échoué: %v", err)
		}
		if err := repo.Create(s2); err == nil {
			t.Fatal("deuxième création avec le même email devrait échouer")
		}
	})

	t.Run("trouver par email", func(t *testing.T) {
		email := "findbyemail-" + randString(6) + "@test.com"
		original := &domain.Student{Email: email, PasswordHash: "pass", Lang: "en"}
		if err := repo.Create(original); err != nil {
			t.Fatalf("création a échoué: %v", err)
		}

		found, err := repo.FindByEmail(email)
		if err != nil {
			t.Fatalf("FindByEmail a échoué: %v", err)
		}
		if found.Email != email {
			t.Fatalf("email attendu '%s', got '%s'", email, found.Email)
		}
		if found.Lang != "en" {
			t.Fatalf("lang attendu 'en', got '%s'", found.Lang)
		}
	})

	t.Run("trouver par email inexistant", func(t *testing.T) {
		_, err := repo.FindByEmail("nonexistent-" + randString(6) + "@test.com")
		if err == nil {
			t.Fatal("devrait retourner une erreur pour email inexistant")
		}
	})

	t.Run("trouver par ID", func(t *testing.T) {
		email := "findbyid-" + randString(6) + "@test.com"
		original := &domain.Student{Email: email, PasswordHash: "pass", Lang: "fr"}
		if err := repo.Create(original); err != nil {
			t.Fatalf("création a échoué: %v", err)
		}

		found, err := repo.FindByID(original.ID)
		if err != nil {
			t.Fatalf("FindByID a échoué: %v", err)
		}
		if found.ID != original.ID {
			t.Fatalf("ID attendu '%s', got '%s'", original.ID, found.ID)
		}
	})

	t.Run("mettre à jour la langue", func(t *testing.T) {
		email := "updatelang-" + randString(6) + "@test.com"
		original := &domain.Student{Email: email, PasswordHash: "pass", Lang: "fr"}
		if err := repo.Create(original); err != nil {
			t.Fatalf("création a échoué: %v", err)
		}

		err := repo.UpdateLang(original.ID, "en")
		if err != nil {
			t.Fatalf("UpdateLang a échoué: %v", err)
		}

		updated, _ := repo.FindByID(original.ID)
		if updated.Lang != "en" {
			t.Fatalf("lang attendu 'en', got '%s'", updated.Lang)
		}
	})
}
