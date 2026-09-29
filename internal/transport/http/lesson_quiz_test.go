package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"aeropath/internal/domain"
	"aeropath/internal/learning"
	"github.com/gin-gonic/gin"
)

type lessonQuizQuestions struct {
	domain.QuestionRepository
	questions []*domain.Question
}

func (r lessonQuizQuestions) FindAll() ([]*domain.Question, error) { return r.questions, nil }

func TestLessonQuizHasFifteenQuestionsAndLessonContext(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var questions []*domain.Question
	for i := 0; i < 18; i++ {
		questions = append(questions, &domain.Question{
			ID: fmt.Sprintf("q%d", i), LessonID: "lesson1", License: "PPL", Category: "navigation",
			QuestionFr: "Question", Options: []string{"Faux", "Juste"}, AnswerKey: "Juste", ExplanationFr: "Correction",
		})
	}
	questions = append(questions, &domain.Question{ID: "other", LessonID: "lesson2"})
	router := gin.New()
	router.GET("/lessons/:id/quiz", LessonQuizHandler(learning.NewEngine(lessonQuizQuestions{questions: questions}, nil)))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/lessons/lesson1/quiz", nil))
	if response.Code != http.StatusOK {
		t.Fatalf("status %d: %s", response.Code, response.Body.String())
	}
	var body struct {
		LessonID  string           `json:"lesson_id"`
		Questions []map[string]any `json:"questions"`
	}
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if body.LessonID != "lesson1" || len(body.Questions) != 15 {
		t.Fatalf("unexpected quiz: %+v", body)
	}
	seen := map[string]bool{}
	for _, q := range body.Questions {
		id := q["id"].(string)
		if id == "other" || seen[id] {
			t.Fatalf("unrelated or duplicate question: %s", id)
		}
		seen[id] = true
		if q["license"] != "PPL" || q["category"] != "navigation" {
			t.Fatalf("missing lesson navigation context: %+v", q)
		}
		for _, key := range []string{"answer_key", "correct_answer", "explanation_fr"} {
			if _, ok := q[key]; ok {
				t.Fatalf("correction exposed before answer: %s", key)
			}
		}
	}
}
