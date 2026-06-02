package auth

import (
	"errors"
	"testing"

	"aeropath/internal/domain"
)

// mockStudentRepo implémente domain.StudentRepository pour les tests
type mockStudentRepo struct {
	students map[string]*domain.Student // email -> student
	byID     map[string]*domain.Student // id -> student
}

func newMockStudentRepo() *mockStudentRepo {
	return &mockStudentRepo{
		students: make(map[string]*domain.Student),
		byID:     make(map[string]*domain.Student),
	}
}

func (m *mockStudentRepo) Create(s *domain.Student) error {
	if _, exists := m.students[s.Email]; exists {
		return errors.New("email déjà utilisé")
	}
	s.ID = "mock-uuid-" + s.Email
	m.students[s.Email] = s
	m.byID[s.ID] = s
	return nil
}

func (m *mockStudentRepo) FindByEmail(email string) (*domain.Student, error) {
	s, exists := m.students[email]
	if !exists {
		return nil, errors.New("student introuvable")
	}
	return s, nil
}

func (m *mockStudentRepo) FindByID(id string) (*domain.Student, error) {
	s, exists := m.byID[id]
	if !exists {
		return nil, errors.New("student introuvable")
	}
	return s, nil
}

func (m *mockStudentRepo) Count() (int, error) {
	return len(m.students), nil
}

func (m *mockStudentRepo) Delete(id string) error {
	_, exists := m.byID[id]
	if !exists {
		return errors.New("étudiant introuvable")
	}
	s := m.byID[id]
	delete(m.byID, id)
	delete(m.students, s.Email)
	return nil
}

func (m *mockStudentRepo) UpdateLang(id, lang string) error {
	s, exists := m.byID[id]
	if !exists {
		return errors.New("student introuvable")
	}
	s.Lang = lang
	return nil
}

func TestRegister(t *testing.T) {
	t.Setenv("JWT_SECRET", "test-secret")
	repo := newMockStudentRepo()
	svc := NewService(repo, "test-secret")

	t.Run("inscription réussie", func(t *testing.T) {
		token, err := svc.Register("test@test.com", "password123", "fr")
		if err != nil {
			t.Fatalf("Register a échoué: %v", err)
		}
		if token == "" {
			t.Fatal("token ne devrait pas être vide")
		}
	})

	t.Run("inscription avec email déjà utilisé", func(t *testing.T) {
		_, err := svc.Register("test@test.com", "password123", "fr")
		if err == nil {
			t.Fatal("devrait retourner une erreur pour email dupliqué")
		}
	})

	t.Run("inscription en anglais", func(t *testing.T) {
		token, err := svc.Register("english@test.com", "password123", "en")
		if err != nil {
			t.Fatalf("Register a échoué: %v", err)
		}
		if token == "" {
			t.Fatal("token ne devrait pas être vide")
		}
		student, _ := repo.FindByEmail("english@test.com")
		if student.Lang != "en" {
			t.Fatalf("langue attendue 'en', got '%s'", student.Lang)
		}
	})
}

func TestLogin(t *testing.T) {
	t.Setenv("JWT_SECRET", "test-secret")
	repo := newMockStudentRepo()
	svc := NewService(repo, "test-secret")

	// Créer un étudiant pour tester le login
	_, err := svc.Register("login@test.com", "password123", "fr")
	if err != nil {
		t.Fatalf("Register a échoué: %v", err)
	}

	t.Run("connexion réussie", func(t *testing.T) {
		token, err := svc.Login("login@test.com", "password123")
		if err != nil {
			t.Fatalf("Login a échoué: %v", err)
		}
		if token == "" {
			t.Fatal("token ne devrait pas être vide")
		}
	})

	t.Run("mauvais mot de passe", func(t *testing.T) {
		_, err := svc.Login("login@test.com", "wrongpassword")
		if err == nil {
			t.Fatal("devrait retourner une erreur pour mauvais mot de passe")
		}
	})

	t.Run("email inexistant", func(t *testing.T) {
		_, err := svc.Login("unknown@test.com", "password123")
		if err == nil {
			t.Fatal("devrait retourner une erreur pour email inconnu")
		}
	})
}

func TestUpdateLang(t *testing.T) {
	t.Setenv("JWT_SECRET", "test-secret")
	repo := newMockStudentRepo()
	svc := NewService(repo, "test-secret")

	// Créer un étudiant
	_, err := svc.Register("lang@test.com", "password123", "fr")
	if err != nil {
		t.Fatalf("Register a échoué: %v", err)
	}

	student, _ := repo.FindByEmail("lang@test.com")

	t.Run("changement de langue réussi", func(t *testing.T) {
		err := svc.UpdateLang(student.ID, "en")
		if err != nil {
			t.Fatalf("UpdateLang a échoué: %v", err)
		}
		updated, _ := repo.FindByID(student.ID)
		if updated.Lang != "en" {
			t.Fatalf("langue attendue 'en', got '%s'", updated.Lang)
		}
	})

	t.Run("langue invalide", func(t *testing.T) {
		err := svc.UpdateLang(student.ID, "de")
		if err == nil {
			t.Fatal("devrait retourner une erreur pour langue invalide")
		}
	})

	t.Run("étudiant inexistant", func(t *testing.T) {
		err := svc.UpdateLang("invalid-id", "fr")
		if err == nil {
			t.Fatal("devrait retourner une erreur pour étudiant inconnu")
		}
	})
}