package learning

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"aeropath/internal/domain"
)

type mockLessonRepo struct {
	lessons []*domain.Lesson
}

func (m *mockLessonRepo) Create(l *domain.Lesson) error {
	return nil
}

func (m *mockLessonRepo) Update(l *domain.Lesson) error {
	return nil
}

func (m *mockLessonRepo) FindByID(id string) (*domain.Lesson, error) {
	for _, l := range m.lessons {
		if l.ID == id {
			return l, nil
		}
	}
	return nil, nil // Le handler gère le nil comme "introuvable"
}

func (m *mockLessonRepo) FindByTheme(theme string) ([]*domain.Lesson, error) {
	var filtered []*domain.Lesson
	for _, l := range m.lessons {
		if l.Theme == theme {
			filtered = append(filtered, l)
		}
	}
	return filtered, nil
}

func (m *mockLessonRepo) FindAll() ([]*domain.Lesson, error) {
	return m.lessons, nil
}

func (m *mockLessonRepo) FindAllPaginated(limit, offset int) ([]*domain.Lesson, error) {
	if offset > len(m.lessons) {
		return []*domain.Lesson{}, nil
	}
	end := offset + limit
	if end > len(m.lessons) {
		end = len(m.lessons)
	}
	return m.lessons[offset:end], nil
}

func (m *mockLessonRepo) FindByLicense(license domain.License) ([]*domain.Lesson, error) {
	var filtered []*domain.Lesson
	for _, l := range m.lessons {
		if l.License == license {
			filtered = append(filtered, l)
		}
	}
	return filtered, nil
}

func (m *mockLessonRepo) FindByCategory(category domain.Category) ([]*domain.Lesson, error) {
	var filtered []*domain.Lesson
	for _, l := range m.lessons {
		if l.Category == category {
			filtered = append(filtered, l)
		}
	}
	return filtered, nil
}

func (m *mockLessonRepo) FindByLicenseAndCategory(license domain.License, category domain.Category) ([]*domain.Lesson, error) {
	var filtered []*domain.Lesson
	for _, l := range m.lessons {
		if l.License == license && l.Category == category {
			filtered = append(filtered, l)
		}
	}
	return filtered, nil
}

func (m *mockLessonRepo) FindByDifficulty(difficulty int) ([]*domain.Lesson, error) {
	var filtered []*domain.Lesson
	for _, l := range m.lessons {
		if l.Difficulty == difficulty {
			filtered = append(filtered, l)
		}
	}
	return filtered, nil
}

func (m *mockLessonRepo) CountByLicense(license domain.License) (int, error) {
	count := 0
	for _, l := range m.lessons {
		if l.License == license {
			count++
		}
	}
	return count, nil
}

func (m *mockLessonRepo) CountByCategory(category domain.Category) (int, error) {
	count := 0
	for _, l := range m.lessons {
		if l.Category == category {
			count++
		}
	}
	return count, nil
}

func (m *mockLessonRepo) Delete(id string) error {
	for i, l := range m.lessons {
		if l.ID == id {
			m.lessons = append(m.lessons[:i], m.lessons[i+1:]...)
			return nil
		}
	}
	return errors.New("leçon introuvable")
}

func (m *mockLessonRepo) Count() (int, error) {
	return len(m.lessons), nil
}

func setupTestContext(method, path, body string) (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	// Simuler un étudiant connecté
	c.Set("student_id", "test-student-id")

	return c, w
}

func TestListLessonsHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := &mockLessonRepo{
		lessons: []*domain.Lesson{
			{ID: "l1", Theme: "meteorology", TitleFr: "Les nuages", TitleEn: "Clouds"},
			{ID: "l2", Theme: "navigation", TitleFr: "Le VOR", TitleEn: "VOR Navigation"},
		},
	}

	t.Run("liste toutes les leçons", func(t *testing.T) {
		c, w := setupTestContext("GET", "/api/lessons", "")
		ListLessonsHandler(repo)(c)

		if w.Code != http.StatusOK {
			t.Fatalf("code attendu 200, got %d", w.Code)
		}

		var lessons []*domain.Lesson
		if err := json.Unmarshal(w.Body.Bytes(), &lessons); err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}
		if len(lessons) != 2 {
			t.Fatalf("devrait retourner 2 leçons, got %d", len(lessons))
		}
	})
}

func TestGetLessonHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := &mockLessonRepo{
		lessons: []*domain.Lesson{
			{ID: "l1", Theme: "meteorology", TitleFr: "Les nuages", TitleEn: "Clouds"},
		},
	}

	t.Run("leçon existante", func(t *testing.T) {
		c, w := setupTestContext("GET", "/api/lessons/l1", "")
		c.Params = []gin.Param{{Key: "id", Value: "l1"}}
		GetLessonHandler(repo)(c)

		if w.Code != http.StatusOK {
			t.Fatalf("code attendu 200, got %d", w.Code)
		}

		var lesson domain.Lesson
		if err := json.Unmarshal(w.Body.Bytes(), &lesson); err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}
		if lesson.ID != "l1" {
			t.Fatalf("id attendu 'l1', got '%s'", lesson.ID)
		}
	})

	t.Run("leçon inexistante", func(t *testing.T) {
		c, w := setupTestContext("GET", "/api/lessons/unknown", "")
		c.Params = []gin.Param{{Key: "id", Value: "unknown"}}
		GetLessonHandler(repo)(c)

		if w.Code != http.StatusNotFound {
			t.Fatalf("code attendu 404, got %d", w.Code)
		}
	})
}

func TestListLessonsByThemeHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := &mockLessonRepo{
		lessons: []*domain.Lesson{
			{ID: "l1", Theme: "meteorology", TitleFr: "Les nuages", TitleEn: "Clouds"},
			{ID: "l2", Theme: "meteorology", TitleFr: "La pression", TitleEn: "Pressure"},
			{ID: "l3", Theme: "navigation", TitleFr: "Le VOR", TitleEn: "VOR"},
		},
	}

	t.Run("leçons d'un thème existant", func(t *testing.T) {
		c, w := setupTestContext("GET", "/api/lessons/theme/meteorology", "")
		c.Params = []gin.Param{{Key: "theme", Value: "meteorology"}}
		ListLessonsByThemeHandler(repo)(c)

		if w.Code != http.StatusOK {
			t.Fatalf("code attendu 200, got %d", w.Code)
		}

		var lessons []*domain.Lesson
		_ = json.Unmarshal(w.Body.Bytes(), &lessons)
		if len(lessons) != 2 {
			t.Fatalf("devrait retourner 2 leçons, got %d", len(lessons))
		}
	})
}

func TestLessonQuizHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	qr := &mockQuestionRepo{questions: newTestQuestions()}
	hr := &mockHistoryRepo{}
	engine := NewEngine(qr, hr)

	t.Run("quiz pour leçon existante", func(t *testing.T) {
		c, w := setupTestContext("GET", "/api/lessons/lesson-1/quiz", "")
		c.Params = []gin.Param{{Key: "id", Value: "lesson-1"}}
		LessonQuizHandler(engine)(c)

		if w.Code != http.StatusOK {
			t.Fatalf("code attendu 200, got %d", w.Code)
		}

		var response struct {
			LessonID  string `json:"lesson_id"`
			Questions []gin.H `json:"questions"`
		}
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}
		if response.LessonID != "lesson-1" {
			t.Fatalf("lesson_id attendu 'lesson-1', got '%s'", response.LessonID)
		}
		if len(response.Questions) != 3 {
			t.Fatalf("devrait retourner 3 questions, got %d", len(response.Questions))
		}
	})

	t.Run("quiz pour leçon inexistante", func(t *testing.T) {
		c, w := setupTestContext("GET", "/api/lessons/unknown/quiz", "")
		c.Params = []gin.Param{{Key: "id", Value: "unknown"}}
		LessonQuizHandler(engine)(c)

		if w.Code != http.StatusNotFound {
			t.Fatalf("code attendu 404, got %d", w.Code)
		}
	})
}

func TestStartExamHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	qr := &mockQuestionRepo{questions: newTestQuestions()}
	hr := &mockHistoryRepo{}
	engine := NewEngine(qr, hr)

	t.Run("examen pour thème existant", func(t *testing.T) {
		c, w := setupTestContext("GET", "/api/exam/meteorology", "")
		c.Params = []gin.Param{{Key: "theme", Value: "meteorology"}}
		StartExamHandler(engine)(c)

		if w.Code != http.StatusOK {
			t.Fatalf("code attendu 200, got %d", w.Code)
		}

		var response struct {
			Theme     string  `json:"theme"`
			Total     int     `json:"total"`
			Questions []gin.H `json:"questions"`
		}
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}
		if response.Theme != "meteorology" {
			t.Fatalf("theme attendu 'meteorology', got '%s'", response.Theme)
		}
		if response.Total != 3 {
			t.Fatalf("total attendu 3, got %d", response.Total)
		}
	})

	t.Run("examen pour thème inexistant", func(t *testing.T) {
		c, w := setupTestContext("GET", "/api/exam/unknown", "")
		c.Params = []gin.Param{{Key: "theme", Value: "unknown"}}
		StartExamHandler(engine)(c)

		if w.Code != http.StatusNotFound {
			t.Fatalf("code attendu 404, got %d", w.Code)
		}
	})
}

func TestSubmitExamHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	qr := &mockQuestionRepo{questions: newTestQuestions()}
	hr := &mockHistoryRepo{}
	engine := NewEngine(qr, hr)

	t.Run("soumission d'examen avec réponses correctes", func(t *testing.T) {
		body := `{"answers":[{"question_id":"q1","answer":"A"},{"question_id":"q2","answer":"B"}]}`
		c, w := setupTestContext("POST", "/api/exam/submit", body)
		SubmitExamHandler(engine, hr, qr)(c)

		if w.Code != http.StatusOK {
			t.Fatalf("code attendu 200, got %d", w.Code)
		}

		var response struct {
			Score   float64 `json:"score"`
			Correct int     `json:"correct"`
			Total   int     `json:"total"`
		}
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}
		if response.Correct != 2 {
			t.Fatalf("correct attendu 2, got %d", response.Correct)
		}
		if response.Total != 2 {
			t.Fatalf("total attendu 2, got %d", response.Total)
		}
		if response.Score != 100.0 {
			t.Fatalf("score attendu 100, got %f", response.Score)
		}
	})

	t.Run("soumission avec réponses incorrectes", func(t *testing.T) {
		body := `{"answers":[{"question_id":"q1","answer":"B"},{"question_id":"q2","answer":"C"}]}`
		c, w := setupTestContext("POST", "/api/exam/submit", body)
		SubmitExamHandler(engine, hr, qr)(c)

		if w.Code != http.StatusOK {
			t.Fatalf("code attendu 200, got %d", w.Code)
		}

		var response struct {
			Score   float64 `json:"score"`
			Correct int     `json:"correct"`
			Total   int     `json:"total"`
		}
		_ = json.Unmarshal(w.Body.Bytes(), &response)
		if response.Correct != 0 {
			t.Fatalf("correct attendu 0, got %d", response.Correct)
		}
	})

	t.Run("soumission sans réponses", func(t *testing.T) {
		body := `{}`
		c, w := setupTestContext("POST", "/api/exam/submit", body)
		SubmitExamHandler(engine, hr, qr)(c)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("code attendu 400, got %d", w.Code)
		}
	})
}

func TestStatsHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	hr := &mockHistoryRepo{}
	// Ajouter un peu d'historique
	_ = hr.RecordAnswer("test-student-id", "q1", true)
	_ = hr.RecordAnswer("test-student-id", "q2", false)
	_ = hr.RecordAnswer("test-student-id", "q3", true)

	t.Run("stats de l'étudiant", func(t *testing.T) {
		c, w := setupTestContext("GET", "/api/stats", "")
		StatsHandler(hr)(c)

		if w.Code != http.StatusOK {
			t.Fatalf("code attendu 200, got %d", w.Code)
		}

		var stats domain.StudentStats
		if err := json.Unmarshal(w.Body.Bytes(), &stats); err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}
		if stats.TotalQuestions != 3 {
			t.Fatalf("total attendu 3, got %d", stats.TotalQuestions)
		}
		if stats.CorrectAnswers != 2 {
			t.Fatalf("correct attendu 2, got %d", stats.CorrectAnswers)
		}
		if stats.WrongAnswers != 1 {
			t.Fatalf("wrong attendu 1, got %d", stats.WrongAnswers)
		}
	})
}
