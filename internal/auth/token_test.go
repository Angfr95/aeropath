package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGenerateToken(t *testing.T) {
	t.Run("génération de token réussie", func(t *testing.T) {
		token, err := GenerateToken("student-123", "my-secret-key")
		if err != nil {
			t.Fatalf("GenerateToken a échoué: %v", err)
		}
		if token == "" {
			t.Fatal("token ne devrait pas être vide")
		}
	})

	t.Run("token différent pour des étudiants différents", func(t *testing.T) {
		token1, _ := GenerateToken("student-1", "secret")
		token2, _ := GenerateToken("student-2", "secret")
		if token1 == token2 {
			t.Fatal("deux tokens différents ne devraient pas être identiques")
		}
	})

	t.Run("token différent avec des secrets différents", func(t *testing.T) {
		token1, _ := GenerateToken("student-1", "secret-1")
		token2, _ := GenerateToken("student-1", "secret-2")
		if token1 == token2 {
			t.Fatal("deux tokens avec des secrets différents ne devraient pas être identiques")
		}
	})
}

func TestRequireAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("accès sans token", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		req := httptest.NewRequest("GET", "/api/me", nil)
		c.Request = req

		RequireAuth("secret")(c)

		if w.Code != http.StatusUnauthorized {
			t.Fatalf("code attendu 401, got %d", w.Code)
		}
	})

	t.Run("accès avec mauvais format de token", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		req := httptest.NewRequest("GET", "/api/me", nil)
		req.Header.Set("Authorization", "InvalidFormat")
		c.Request = req

		RequireAuth("secret")(c)

		if w.Code != http.StatusUnauthorized {
			t.Fatalf("code attendu 401, got %d", w.Code)
		}
	})

	t.Run("accès avec token invalide", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		req := httptest.NewRequest("GET", "/api/me", nil)
		req.Header.Set("Authorization", "Bearer invalid-token")
		c.Request = req

		RequireAuth("secret")(c)

		if w.Code != http.StatusUnauthorized {
			t.Fatalf("code attendu 401, got %d", w.Code)
		}
	})

	t.Run("accès avec token valide", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		token, _ := GenerateToken("student-123", "my-secret")
		req := httptest.NewRequest("GET", "/api/me", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		c.Request = req

		RequireAuth("my-secret")(c)

		if w.Code != 200 { // Le middleware passe au handler suivant
			t.Fatalf("code attendu 200, got %d", w.Code)
		}

		studentID, exists := c.Get("student_id")
		if !exists {
			t.Fatal("student_id devrait être défini dans le contexte")
		}
		if studentID != "student-123" {
			t.Fatalf("student_id attendu 'student-123', got '%v'", studentID)
		}
	})

	t.Run("token signé avec un secret différent", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		token, _ := GenerateToken("student-123", "wrong-secret")
		req := httptest.NewRequest("GET", "/api/me", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		c.Request = req

		RequireAuth("correct-secret")(c)

		if w.Code != http.StatusUnauthorized {
			t.Fatalf("code attendu 401, got %d", w.Code)
		}
	})
}
