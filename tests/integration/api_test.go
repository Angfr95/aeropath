//go:build integration

package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"aeropath/internal/admin"
	"aeropath/internal/auth"
	"aeropath/internal/domain"
	"aeropath/internal/learning"
	"aeropath/internal/persistence/postgres"
	"aeropath/internal/persistence/redis"
	"aeropath/internal/transport/http"
)

// Tests d'intégration complets nécessitant PostgreSQL et Redis.
// Exécution : go test -tags=integration ./tests/integration/

var (
	dbPool  *pgxpool.Pool
	rdb     *redis.Client
	router  *gin.Engine
	studID  string
	token   string
)

func TestMain(m *testing.M) {
	// Connexion à la base de test
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://aeropath:aeropath@localhost:5432/aeropath_test?sslmode=disable"
	}

	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://localhost:6379/1"
	}

	var err error
	dbPool, err = pgxpool.New(nil, dbURL)
	if err != nil {
		fmt.Printf("⚠️  PostgreSQL non disponible (DATABASE_URL=%s): %v\n", dbURL, err)
		fmt.Println("Les tests d'intégration sont ignorés.")
		fmt.Println("Pour lancer : docker compose up -d postgres redis")
		os.Exit(0)
	}
	defer dbPool.Close()

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		fmt.Printf("⚠️  Redis non disponible: %v\n", err)
		os.Exit(0)
	}
	rdb = redis.NewClient(opt)
	defer rdb.Close()

	// Setup
	setupTestData()
	setupRouter()

	code := m.Run()
	teardownTestData()
	os.Exit(code)
}

func setupTestData() {
	// Créer un étudiant de test
	err := dbPool.QueryRow(nil, `
		INSERT INTO students (id, email, lang, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		ON CONFLICT (id) DO NOTHING
		RETURNING id
	`, "test-student-001", "test@aeropath.com", "fr").Scan(&studID)
	if err != nil {
		studID = "test-student-001"
	}

	// Créer des questions de test
	for i := 1; i <= 5; i++ {
		opts := []string{"Option A", "Option B", "Option C", "Option D"}
		optsJSON, _ := json.Marshal(opts)
		dbPool.Exec(nil, `
			INSERT INTO questions (id, question_fr, question_en, options, answer_key,
				explanation_fr, explanation_en, license, category, theme, subtopic, difficulty, reference)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
			ON CONFLICT (id) DO NOTHING
		`, fmt.Sprintf("test-q-%d", i),
			fmt.Sprintf("Question test %d FR", i),
			fmt.Sprintf("Question test %d EN", i),
			optsJSON, "A",
			"Explication FR", "Explication EN",
			"PPL", "Navigation", "Météorologie", "Thème A",
			i, "REF-001")
	}

	// Créer des leçons de test
	for i := 1; i <= 3; i++ {
		dbPool.Exec(nil, `
			INSERT INTO lessons (id, title_fr, title_en, content_fr, content_en,
				license, category, theme, difficulty, order_index)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
			ON CONFLICT (id) DO NOTHING
		`, fmt.Sprintf("test-l-%d", i),
			fmt.Sprintf("Leçon test %d FR", i),
			fmt.Sprintf("Leçon test %d EN", i),
			"Contenu FR", "Contenu EN",
			"PPL", "Navigation", "Météorologie", i, i)
	}

	// Créer un historique de test
	dbPool.Exec(nil, `
		INSERT INTO answer_history (id, student_id, question_id, was_correct, answered_at)
		VALUES ($1,$2,$3,$4,$5)
		ON CONFLICT (id) DO NOTHING
	`, "test-ah-001", studID, "test-q-1", true, time.Now().Add(-24*time.Hour))
	dbPool.Exec(nil, `
		INSERT INTO answer_history (id, student_id, question_id, was_correct, answered_at)
		VALUES ($1,$2,$3,$4,$5)
		ON CONFLICT (id) DO NOTHING
	`, "test-ah-002", studID, "test-q-2", false, time.Now().Add(-12*time.Hour))
}

func setupRouter() {
	gin.SetMode(gin.TestMode)
	router = gin.New()

	// Repositories
	studentRepo := postgres.NewStudentRepository(dbPool)
	questionRepo := postgres.NewQuestionRepository(dbPool)
	lessonRepo := postgres.NewLessonRepository(dbPool)
	adminRepo := admin.NewPostgresRepository(dbPool)
	cache := redis.NewCache(rdb)

	// Services
	authService := auth.NewService(studentRepo, "test-secret-key-1234567890")
	learningEngine := learning.NewEngine(questionRepo, studentRepo, cache)

	// Handlers
	authHandler := http.NewAuthHandler(authService)
	learningHandler := http.NewLearningHandler(learningEngine, studentRepo)
	adminHandler := admin.NewHandler(adminRepo)

	// Routes
	api := router.Group("/api")
	authHandler.RegisterRoutes(api)
	learningHandler.RegisterRoutes(api)
	adminHandler.RegisterRoutes(api)

	// Auth middleware pour les routes protégées
	api.Use(func(c *gin.Context) {
		c.Set("student_id", studID)
		c.Next()
	})

	// Générer un token de test
	var err error
	token, err = authService.GenerateToken(nil, "test-student-001", "test@aeropath.com")
	if err != nil {
		panic(err)
	}
}

func teardownTestData() {
	dbPool.Exec(nil, "DELETE FROM answer_history WHERE id LIKE 'test-%'")
	dbPool.Exec(nil, "DELETE FROM questions WHERE id LIKE 'test-%'")
	dbPool.Exec(nil, "DELETE FROM lessons WHERE id LIKE 'test-%'")
	dbPool.Exec(nil, "DELETE FROM students WHERE id LIKE 'test-%'")
}

// ======================== TESTS ========================

func TestAuthFlow(t *testing.T) {
	t.Run("Register", func(t *testing.T) {
		body := `{"email":"new@test.com","password":"password123","lang":"fr"}`
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NotEmpty(t, resp["token"])
		assert.NotEmpty(t, resp["student_id"])
	})

	t.Run("Login", func(t *testing.T) {
		body := `{"email":"test@aeropath.com","password":"password123"}`
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NotEmpty(t, resp["token"])
	})
}

func TestQuestionsFlow(t *testing.T) {
	t.Run("ListQuestions", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/questions?page=1&page_size=10", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NotEmpty(t, resp["data"])
		assert.GreaterOrEqual(t, resp["total"], float64(5))
	})

	t.Run("GetQuestion", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/questions/test-q-1", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var q domain.Question
		json.Unmarshal(w.Body.Bytes(), &q)
		assert.Equal(t, "test-q-1", q.ID)
		assert.Contains(t, q.QuestionFr, "Question test")
	})

	t.Run("AnswerQuestion", func(t *testing.T) {
		body := `{"question_id":"test-q-1","answer":"A"}`
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/questions/answer", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Equal(t, true, resp["was_correct"])
	})
}

func TestLessonsFlow(t *testing.T) {
	t.Run("ListLessons", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/lessons?page=1&page_size=10", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NotEmpty(t, resp["data"])
		assert.GreaterOrEqual(t, resp["total"], float64(3))
	})

	t.Run("GetLesson", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/lessons/test-l-1", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var l domain.Lesson
		json.Unmarshal(w.Body.Bytes(), &l)
		assert.Equal(t, "test-l-1", l.ID)
	})
}

func TestHistoryFlow(t *testing.T) {
	t.Run("GetHistory", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/history?page=1&page_size=10", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NotEmpty(t, resp["data"])
	})

	t.Run("GetStats", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/stats", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NotNil(t, resp["total_answers"])
	})
}

func TestRecommendationsFlow(t *testing.T) {
	t.Run("GetRecommendations", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/recommendations", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NotNil(t, resp["weak_topics"])
		assert.NotNil(t, resp["due_cards"])
	})
}

func TestAdminFlow(t *testing.T) {
	t.Run("GetStats", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/admin/stats", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.GreaterOrEqual(t, resp["total_students"], float64(1))
		assert.GreaterOrEqual(t, resp["total_questions"], float64(5))
	})

	t.Run("ListStudents", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/admin/students?page=1&page_size=10", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NotEmpty(t, resp["data"])
	})

	t.Run("CRUDQuestion", func(t *testing.T) {
		// Create
		body := `{"question_fr":"Nouvelle Q FR","question_en":"New Q EN","options":["A","B","C"],"answer_key":"A","license":"PPL","category":"Test","theme":"Test","difficulty":1}`
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/admin/questions", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)
		require.Equal(t, http.StatusCreated, w.Code)

		var q map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &q)
		qID := q["id"].(string)

		// Delete
		w = httptest.NewRecorder()
		req, _ = http.NewRequest("DELETE", "/api/admin/questions/"+qID, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("CRUDLesson", func(t *testing.T) {
		// Create
		body := `{"title_fr":"Nouvelle L FR","title_en":"New L EN","content_fr":"Contenu","content_en":"Content","license":"PPL","category":"Test","theme":"Test","difficulty":1}`
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/admin/lessons", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)
		require.Equal(t, http.StatusCreated, w.Code)

		var l map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &l)
		lID := l["id"].(string)

		// Delete
		w = httptest.NewRecorder()
		req, _ = http.NewRequest("DELETE", "/api/admin/lessons/"+lID, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestQuizFlow(t *testing.T) {
	t.Run("GetQuiz", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/quiz?count=5&license=PPL", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp []domain.Question
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Len(t, resp, 5)
	})
}

func TestPagination(t *testing.T) {
	t.Run("QuestionsPagination", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/questions?page=1&page_size=2", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Len(t, resp["data"], 2)
		assert.Equal(t, float64(1), resp["page"])
		assert.Equal(t, float64(2), resp["page_size"])
	})

	t.Run("LessonsPagination", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/lessons?page=1&page_size=2", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Len(t, resp["data"], 2)
	})
}

func TestErrorCases(t *testing.T) {
	t.Run("Unauthorized", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/questions", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("InvalidToken", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/questions", nil)
		req.Header.Set("Authorization", "Bearer invalid-token")
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("NotFound", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/questions/nonexistent", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("InvalidAnswer", func(t *testing.T) {
		body := `{"question_id":"test-q-1","answer":"Z"}`
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/questions/answer", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Equal(t, false, resp["was_correct"])
	})
}
