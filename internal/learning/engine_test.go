package learning

import (
	"errors"
	"testing"

	"aeropath/internal/domain"
)

// mocks
type mockQuestionRepo struct {
	questions []*domain.Question
}

func (m *mockQuestionRepo) Create(q *domain.Question) error {
	return nil
}

func (m *mockQuestionRepo) Update(q *domain.Question) error {
	return nil
}

func (m *mockQuestionRepo) FindByID(id string) (*domain.Question, error) {
	for _, q := range m.questions {
		if q.ID == id {
			return q, nil
		}
	}
	return nil, errors.New("question introuvable")
}

func (m *mockQuestionRepo) FindAll() ([]*domain.Question, error) {
	return m.questions, nil
}

func (m *mockQuestionRepo) FindAllPaginated(limit, offset int) ([]*domain.Question, error) {
	if offset > len(m.questions) {
		return []*domain.Question{}, nil
	}
	end := offset + limit
	if end > len(m.questions) {
		end = len(m.questions)
	}
	return m.questions[offset:end], nil
}

func (m *mockQuestionRepo) FindByTheme(theme string) ([]*domain.Question, error) {
	var filtered []*domain.Question
	for _, q := range m.questions {
		if q.Theme == theme {
			filtered = append(filtered, q)
		}
	}
	return filtered, nil
}

func (m *mockQuestionRepo) FindByLicense(license domain.License) ([]*domain.Question, error) {
	var filtered []*domain.Question
	for _, q := range m.questions {
		if q.License == license {
			filtered = append(filtered, q)
		}
	}
	return filtered, nil
}

func (m *mockQuestionRepo) FindByCategory(category domain.Category) ([]*domain.Question, error) {
	var filtered []*domain.Question
	for _, q := range m.questions {
		if q.Category == category {
			filtered = append(filtered, q)
		}
	}
	return filtered, nil
}

func (m *mockQuestionRepo) FindByLicenseAndCategory(license domain.License, category domain.Category) ([]*domain.Question, error) {
	var filtered []*domain.Question
	for _, q := range m.questions {
		if q.License == license && q.Category == category {
			filtered = append(filtered, q)
		}
	}
	return filtered, nil
}

func (m *mockQuestionRepo) FindByDifficulty(difficulty int) ([]*domain.Question, error) {
	var filtered []*domain.Question
	for _, q := range m.questions {
		if q.Difficulty == difficulty {
			filtered = append(filtered, q)
		}
	}
	return filtered, nil
}

func (m *mockQuestionRepo) FindBySubtopic(subtopic string) ([]*domain.Question, error) {
	var filtered []*domain.Question
	for _, q := range m.questions {
		if q.Subtopic == subtopic {
			filtered = append(filtered, q)
		}
	}
	return filtered, nil
}

func (m *mockQuestionRepo) CountByLicense(license domain.License) (int, error) {
	count := 0
	for _, q := range m.questions {
		if q.License == license {
			count++
		}
	}
	return count, nil
}

func (m *mockQuestionRepo) CountByCategory(category domain.Category) (int, error) {
	count := 0
	for _, q := range m.questions {
		if q.Category == category {
			count++
		}
	}
	return count, nil
}

func (m *mockQuestionRepo) CountByTheme(theme string) (int, error) {
	count := 0
	for _, q := range m.questions {
		if q.Theme == theme {
			count++
		}
	}
	return count, nil
}

func (m *mockQuestionRepo) Delete(id string) error {
	for i, q := range m.questions {
		if q.ID == id {
			m.questions = append(m.questions[:i], m.questions[i+1:]...)
			return nil
		}
	}
	return errors.New("question introuvable")
}

func (m *mockQuestionRepo) Count() (int, error) {
	return len(m.questions), nil
}

func (m *mockQuestionRepo) FindRandom(limit int, license *domain.License, category *domain.Category, theme *string) ([]*domain.Question, error) {
	// Copier et mélanger
	pool := make([]*domain.Question, len(m.questions))
	copy(pool, m.questions)

	// Filtrer
	var filtered []*domain.Question
	for _, q := range pool {
		if license != nil && q.License != *license {
			continue
		}
		if category != nil && q.Category != *category {
			continue
		}
		if theme != nil && *theme != "" && q.Theme != *theme {
			continue
		}
		filtered = append(filtered, q)
	}

	if limit > len(filtered) {
		limit = len(filtered)
	}
	return filtered[:limit], nil
}

func (m *mockQuestionRepo) Search(query string) ([]*domain.Question, error) {
	var filtered []*domain.Question
	for _, q := range m.questions {
		if contains(q.QuestionFr, query) || contains(q.QuestionEn, query) || contains(q.Theme, query) {
			filtered = append(filtered, q)
		}
	}
	return filtered, nil
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && searchString(s, substr)
}

func searchString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		match := true
		for j := 0; j < len(substr); j++ {
			ci := s[i+j]
			cj := substr[j]
			// case-insensitive
			if ci >= 'A' && ci <= 'Z' {
				ci += 32
			}
			if cj >= 'A' && cj <= 'Z' {
				cj += 32
			}
			if ci != cj {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

type mockHistoryRepo struct {
	history []*domain.QuestionHistory
}

func (m *mockHistoryRepo) RecordAnswer(studentID, questionID string, wasCorrect bool) error {
	m.history = append(m.history, &domain.QuestionHistory{
		StudentID:  studentID,
		QuestionID: questionID,
		WasCorrect: wasCorrect,
	})
	return nil
}

func (m *mockHistoryRepo) GetHistory(studentID string) ([]*domain.QuestionHistory, error) {
	var filtered []*domain.QuestionHistory
	for _, h := range m.history {
		if h.StudentID == studentID {
			filtered = append(filtered, h)
		}
	}
	return filtered, nil
}

func (m *mockHistoryRepo) GetStats(studentID string) (*domain.StudentStats, error) {
	var correct, wrong int
	for _, h := range m.history {
		if h.StudentID == studentID {
			if h.WasCorrect {
				correct++
			} else {
				wrong++
			}
		}
	}
	total := correct + wrong
	rate := 0.0
	if total > 0 {
		rate = float64(correct) / float64(total) * 100
	}
	return &domain.StudentStats{
		TotalQuestions: total,
		CorrectAnswers: correct,
		WrongAnswers:   wrong,
		SuccessRate:    rate,
		CurrentLevel:   1,
	}, nil
}

func (m *mockHistoryRepo) Count() (int, error) {
	return len(m.history), nil
}

func (m *mockHistoryRepo) GetSeenQuestionIDs(studentID string) ([]string, error) {
	var ids []string
	for _, h := range m.history {
		if h.StudentID == studentID {
			ids = append(ids, h.QuestionID)
		}
	}
	return ids, nil
}

func newTestQuestions() []*domain.Question {
	return []*domain.Question{
		{ID: "q1", LessonID: "lesson-1", Theme: "meteorology", Difficulty: 1, QuestionFr: "Question facile?", QuestionEn: "Easy question?", Options: []string{"A", "B", "C"}, AnswerKey: "A"},
		{ID: "q2", LessonID: "lesson-1", Theme: "meteorology", Difficulty: 2, QuestionFr: "Question moyenne?", QuestionEn: "Medium question?", Options: []string{"A", "B", "C"}, AnswerKey: "B"},
		{ID: "q3", LessonID: "lesson-1", Theme: "meteorology", Difficulty: 3, QuestionFr: "Question difficile?", QuestionEn: "Hard question?", Options: []string{"A", "B", "C"}, AnswerKey: "C"},
		{ID: "q4", LessonID: "lesson-2", Theme: "navigation", Difficulty: 1, QuestionFr: "Navigation facile?", QuestionEn: "Easy nav?", Options: []string{"A", "B", "C"}, AnswerKey: "A"},
		{ID: "q5", LessonID: "lesson-2", Theme: "navigation", Difficulty: 2, QuestionFr: "Navigation moyenne?", QuestionEn: "Medium nav?", Options: []string{"A", "B", "C"}, AnswerKey: "B"},
	}
}

func TestLessonQuiz(t *testing.T) {
	qr := &mockQuestionRepo{questions: newTestQuestions()}
	hr := &mockHistoryRepo{}
	engine := NewEngine(qr, hr)

	t.Run("quiz pour une leçon existante", func(t *testing.T) {
		questions, err := engine.LessonQuiz("lesson-1", QuizSize)
		if err != nil {
			t.Fatalf("LessonQuiz a échoué: %v", err)
		}
		if len(questions) != 3 {
			t.Fatalf("devrait retourner 3 questions, got %d", len(questions))
		}
		for _, q := range questions {
			if q.LessonID != "lesson-1" {
				t.Fatalf("question %s ne fait pas partie de la leçon lesson-1", q.ID)
			}
		}
	})

	t.Run("quiz pour une leçon avec moins de 3 questions", func(t *testing.T) {
		questions, err := engine.LessonQuiz("lesson-2", QuizSize)
		if err != nil {
			t.Fatalf("LessonQuiz a échoué: %v", err)
		}
		if len(questions) != 2 {
			t.Fatalf("devrait retourner 2 questions (max disponibles), got %d", len(questions))
		}
	})

	t.Run("quiz pour une leçon inexistante", func(t *testing.T) {
		_, err := engine.LessonQuiz("inexistant", QuizSize)
		if err == nil {
			t.Fatal("devrait retourner une erreur pour leçon inexistante")
		}
	})
}

func TestThemeExam(t *testing.T) {
	qr := &mockQuestionRepo{questions: newTestQuestions()}
	hr := &mockHistoryRepo{}
	engine := NewEngine(qr, hr)

	t.Run("examen pour un thème existant", func(t *testing.T) {
		questions, err := engine.ThemeExam("student-1", "meteorology", ExamSize)
		if err != nil {
			t.Fatalf("ThemeExam a échoué: %v", err)
		}
		if len(questions) != 3 {
			t.Fatalf("devrait retourner 3 questions (max disponibles), got %d", len(questions))
		}
		for _, q := range questions {
			if q.Theme != "meteorology" {
				t.Fatalf("question %s ne fait pas partie du thème meteorology", q.ID)
			}
		}
	})

	t.Run("examen pour un thème inexistant", func(t *testing.T) {
		_, err := engine.ThemeExam("student-1", "inexistant", ExamSize)
		if err == nil {
			t.Fatal("devrait retourner une erreur pour thème inexistant")
		}
	})

	t.Run("examen avec un nombre personnalisé de questions", func(t *testing.T) {
		questions, err := engine.ThemeExam("student-1", "navigation", 1)
		if err != nil {
			t.Fatalf("ThemeExam a échoué: %v", err)
		}
		if len(questions) != 1 {
			t.Fatalf("devrait retourner 1 question, got %d", len(questions))
		}
	})
}

func TestEvaluateAnswer(t *testing.T) {
	qr := &mockQuestionRepo{questions: newTestQuestions()}
	hr := &mockHistoryRepo{}
	engine := NewEngine(qr, hr)

	t.Run("réponse correcte", func(t *testing.T) {
		question := &domain.Question{AnswerKey: "A"}
		if !engine.EvaluateAnswer(question, "A") {
			t.Fatal("'A' devrait être correct")
		}
	})

	t.Run("réponse incorrecte", func(t *testing.T) {
		question := &domain.Question{AnswerKey: "A"}
		if engine.EvaluateAnswer(question, "B") {
			t.Fatal("'B' ne devrait pas être correct")
		}
	})

	t.Run("réponse case-sensitive", func(t *testing.T) {
		question := &domain.Question{AnswerKey: "A"}
		if engine.EvaluateAnswer(question, "a") {
			t.Fatal("'a' ne devrait pas être correct (case-sensitive)")
		}
	})
}
