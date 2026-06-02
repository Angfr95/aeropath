package auth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupAuthContext(method, path, body string) (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	return c, w
}

func TestRegisterHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Setenv("JWT_SECRET", "test-secret")

	repo := newMockStudentRepo()
	svc := NewService(repo, "test-secret")

	t.Run("inscription réussie", func(t *testing.T) {
		c, w := setupAuthContext("POST", "/auth/register", `{"email":"test@test.com","password":"password123","lang":"fr"}`)
		RegisterHandler(svc)(c)

		if w.Code != http.StatusCreated {
			t.Fatalf("code attendu 201, got %d", w.Code)
		}

		var response map[string]string
		_ = json.Unmarshal(w.Body.Bytes(), &response)
		if response["token"] == "" {
			t.Fatal("token ne devrait pas être vide")
		}
	})

	t.Run("inscription sans langue (défaut fr)", func(t *testing.T) {
		c, w := setupAuthContext("POST", "/auth/register", `{"email":"default@test.com","password":"password123"}`)
		RegisterHandler(svc)(c)

		if w.Code != http.StatusCreated {
			t.Fatalf("code attendu 201, got %d", w.Code)
		}
	})

	t.Run("inscription avec email invalide", func(t *testing.T) {
		c, w := setupAuthContext("POST", "/auth/register", `{"email":"invalid","password":"password123"}`)
		RegisterHandler(svc)(c)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("code attendu 400, got %d", w.Code)
		}
	})

	t.Run("inscription avec mot de passe trop court", func(t *testing.T) {
		c, w := setupAuthContext("POST", "/auth/register", `{"email":"test@test.com","password":"short"}`)
		RegisterHandler(svc)(c)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("code attendu 400, got %d", w.Code)
		}
	})

	t.Run("inscription avec langue invalide", func(t *testing.T) {
		c, w := setupAuthContext("POST", "/auth/register", `{"email":"test@test.com","password":"password123","lang":"de"}`)
		RegisterHandler(svc)(c)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("code attendu 400, got %d", w.Code)
		}
	})

	t.Run("email déjà utilisé", func(t *testing.T) {
		c, w := setupAuthContext("POST", "/auth/register", `{"email":"test@test.com","password":"password123"}`)
		RegisterHandler(svc)(c)

		if w.Code != http.StatusConflict {
			t.Fatalf("code attendu 409, got %d", w.Code)
		}
	})
}

func TestLoginHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Setenv("JWT_SECRET", "test-secret")

	repo := newMockStudentRepo()
	svc := NewService(repo, "test-secret")

	// Créer un étudiant pour tester
	_, _ = svc.Register("login@test.com", "password123", "fr")

	t.Run("connexion réussie", func(t *testing.T) {
		c, w := setupAuthContext("POST", "/auth/login", `{"email":"login@test.com","password":"password123"}`)
		LoginHandler(svc)(c)

		if w.Code != http.StatusOK {
			t.Fatalf("code attendu 200, got %d", w.Code)
		}

		var response map[string]string
		_ = json.Unmarshal(w.Body.Bytes(), &response)
		if response["token"] == "" {
			t.Fatal("token ne devrait pas être vide")
		}
	})

	t.Run("mauvais mot de passe", func(t *testing.T) {
		c, w := setupAuthContext("POST", "/auth/login", `{"email":"login@test.com","password":"wrong"}`)
		LoginHandler(svc)(c)

		if w.Code != http.StatusUnauthorized {
			t.Fatalf("code attendu 401, got %d", w.Code)
		}
	})

	t.Run("email inexistant", func(t *testing.T) {
		c, w := setupAuthContext("POST", "/auth/login", `{"email":"unknown@test.com","password":"password123"}`)
		LoginHandler(svc)(c)

		if w.Code != http.StatusUnauthorized {
			t.Fatalf("code attendu 401, got %d", w.Code)
		}
	})

	t.Run("email invalide", func(t *testing.T) {
		c, w := setupAuthContext("POST", "/auth/login", `{"email":"invalid","password":"password123"}`)
		LoginHandler(svc)(c)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("code attendu 400, got %d", w.Code)
		}
	})
}

func TestUpdateLangHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Setenv("JWT_SECRET", "test-secret")

	repo := newMockStudentRepo()
	svc := NewService(repo, "test-secret")

	// Créer un étudiant
	token, _ := svc.Register("lang@test.com", "password123", "fr")
	student, _ := repo.FindByEmail("lang@test.com")

	t.Run("changement de langue réussi", func(t *testing.T) {
		c, w := setupAuthContext("PATCH", "/api/me/lang", `{"lang":"en"}`)
		c.Set("student_id", student.ID)
		UpdateLangHandler(svc)(c)

		if w.Code != http.StatusOK {
			t.Fatalf("code attendu 200, got %d", w.Code)
		}

		var response map[string]string
		_ = json.Unmarshal(w.Body.Bytes(), &response)
		if response["lang"] != "en" {
			t.Fatalf("lang attendu 'en', got '%s'", response["lang"])
		}
	})

	t.Run("langue invalide", func(t *testing.T) {
		c, w := setupAuthContext("PATCH", "/api/me/lang", `{"lang":"de"}`)
		c.Set("student_id", student.ID)
		UpdateLangHandler(svc)(c)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("code attendu 400, got %d", w.Code)
		}
	})

	t.Run("sans student_id dans le contexte", func(t *testing.T) {
		c, w := setupAuthContext("PATCH", "/api/me/lang", `{"lang":"fr"}`)
		// Ne pas set student_id
		UpdateLangHandler(svc)(c)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("code attendu 400, got %d", w.Code)
		}
	})

	// Utiliser token pour éviter l'erreur unused
	_ = token
}
