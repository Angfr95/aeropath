package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"aeropath/internal/domain"
	"aeropath/internal/learning"
	"aeropath/internal/transport/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type practiceQuestions struct{ domain.QuestionRepository }

func (practiceQuestions) FindByID(id string) (*domain.Question, error) {
	return &domain.Question{ID: id, Options: []string{"yes", "no"}, AnswerKey: "yes"}, nil
}

type practiceHistory struct {
	domain.QuestionHistoryRepository
	students []string
}

func (h *practiceHistory) RecordAnswer(student, question string, correct bool) error {
	h.students = append(h.students, student)
	return nil
}

func TestPracticeAuthenticationAndTracking(t *testing.T) {
	gin.SetMode(gin.TestMode)
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"student_id": "student-1", "exp": time.Now().Add(time.Hour).Unix(),
	}).SignedString([]byte("test-secret"))
	if err != nil {
		t.Fatal(err)
	}
	for _, endpoint := range []string{"answer", "exam"} {
		for _, tc := range []struct {
			name, token     string
			status, records int
		}{
			{"guest", "", http.StatusOK, 0},
			{"signed-in", token, http.StatusOK, 1},
			{"invalid-session", "invalid", http.StatusUnauthorized, 0},
		} {
			t.Run(endpoint+"/"+tc.name, func(t *testing.T) {
				questions := practiceQuestions{}
				history := &practiceHistory{}
				engine := learning.NewEngine(questions, history)
				router := gin.New()
				practice := router.Group("/", middleware.OptionalAuth("test-secret"))
				practice.POST("/answer", AnswerQuestionHandler(engine, history, questions))
				practice.POST("/exam", SubmitExamHandler(engine, history, questions))
				body := `{"question_id":"q1","answer":"A","student_id":"spoofed"}`
				if endpoint == "exam" {
					body = `{"answers":[{"question_id":"q1","answer":"A"}],"student_id":"spoofed"}`
				}
				req := httptest.NewRequest(http.MethodPost, "/"+endpoint, strings.NewReader(body))
				req.Header.Set("Content-Type", "application/json")
				if tc.token != "" {
					req.AddCookie(&http.Cookie{Name: "token", Value: tc.token})
				}
				response := httptest.NewRecorder()
				router.ServeHTTP(response, req)
				if response.Code != tc.status {
					t.Fatalf("status %d: %s", response.Code, response.Body.String())
				}
				if len(history.students) != tc.records {
					t.Fatalf("recorded %d answers, want %d", len(history.students), tc.records)
				}
				if tc.records > 0 && history.students[0] != "student-1" {
					t.Fatal("result attributed to wrong student")
				}
				if tc.status == http.StatusOK && !strings.Contains(response.Body.String(), `"correct":true`) {
					t.Fatalf("missing correction: %s", response.Body.String())
				}
			})
		}
	}
}
