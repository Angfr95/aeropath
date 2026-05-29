package learning

import (
	"fmt"
	"math/rand"

	"aeropath/internal/domain"
)

const (
	QuizSize  = 3 // Nombre de questions par quiz de leçon
	ExamSize  = 10 // Nombre de questions par examen de thème
)

// Engine est le moteur d'apprentissage adaptatif.
type Engine struct {
	questionRepo domain.QuestionRepository
	historyRepo  domain.QuestionHistoryRepository
}

func NewEngine(qr domain.QuestionRepository, hr domain.QuestionHistoryRepository) *Engine {
	return &Engine{questionRepo: qr, historyRepo: hr}
}

// LessonQuiz retourne N questions pour un quiz de leçon (vérification rapide).
func (e *Engine) LessonQuiz(lessonID string, count int) ([]*domain.Question, error) {
	allQuestions, err := e.questionRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("find all questions: %w", err)
	}

	// Filtrer les questions de cette leçon
	var lessonQuestions []*domain.Question
	for _, q := range allQuestions {
		if q.LessonID == lessonID {
			lessonQuestions = append(lessonQuestions, q)
		}
	}

	if len(lessonQuestions) == 0 {
		return nil, fmt.Errorf("aucune question pour cette leçon")
	}

	// Mélanger et prendre les N premières
	rand.Shuffle(len(lessonQuestions), func(i, j int) {
		lessonQuestions[i], lessonQuestions[j] = lessonQuestions[j], lessonQuestions[i]
	})

	if count > len(lessonQuestions) {
		count = len(lessonQuestions)
	}
	return lessonQuestions[:count], nil
}

// ThemeExam génère un examen complet sur un thème.
func (e *Engine) ThemeExam(studentID, theme string, count int) ([]*domain.Question, error) {
	allQuestions, err := e.questionRepo.FindByTheme(theme)
	if err != nil {
		return nil, fmt.Errorf("find questions by theme: %w", err)
	}

	if len(allQuestions) == 0 {
		return nil, fmt.Errorf("aucune question pour le thème: %s", theme)
	}

	rand.Shuffle(len(allQuestions), func(i, j int) {
		allQuestions[i], allQuestions[j] = allQuestions[j], allQuestions[i]
	})

	if count > len(allQuestions) {
		count = len(allQuestions)
	}
	return allQuestions[:count], nil
}

// ExamByLicenseAndCategory génère un examen filtré par licence et catégorie.
func (e *Engine) ExamByLicenseAndCategory(license domain.License, category domain.Category, count int) ([]*domain.Question, error) {
	allQuestions, err := e.questionRepo.FindByLicenseAndCategory(license, category)
	if err != nil {
		return nil, fmt.Errorf("find questions by license and category: %w", err)
	}

	if len(allQuestions) == 0 {
		return nil, fmt.Errorf("aucune question pour %s / %s", license, category)
	}

	rand.Shuffle(len(allQuestions), func(i, j int) {
		allQuestions[i], allQuestions[j] = allQuestions[j], allQuestions[i]
	})

	if count > len(allQuestions) {
		count = len(allQuestions)
	}
	return allQuestions[:count], nil
}

// EvaluateAnswer évalue une réponse et retourne si elle est correcte.
func (e *Engine) EvaluateAnswer(question *domain.Question, answer string) bool {
	return question.AnswerKey == answer
}
