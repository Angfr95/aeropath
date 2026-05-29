package learning

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"aeropath/internal/domain"
)

type answerRequest struct {
	QuestionID string `json:"question_id" binding:"required,uuid"`
	Answer     string `json:"answer"      binding:"required"`
}

type submitExamRequest struct {
	Answers []answerRequest `json:"answers" binding:"required,min=1"`
}

// ======================== LEÇONS ========================

// ListLessonsHandler retourne toutes les leçons.
func ListLessonsHandler(repo domain.LessonRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		license := c.Query("license")
		category := c.Query("category")

		var lessons []*domain.Lesson
		var err error

		switch {
		case license != "" && category != "":
			lessons, err = repo.FindByLicenseAndCategory(domain.License(license), domain.Category(category))
		case license != "":
			lessons, err = repo.FindByLicense(domain.License(license))
		case category != "":
			lessons, err = repo.FindByCategory(domain.Category(category))
		default:
			lessons, err = repo.FindAll()
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, lessons)
	}
}

// GetLessonHandler retourne une leçon par son ID.
func GetLessonHandler(repo domain.LessonRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		lesson, err := repo.FindByID(c.Param("id"))
		if err != nil || lesson == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "leçon introuvable"})
			return
		}
		c.JSON(http.StatusOK, lesson)
	}
}

// ListLessonsByThemeHandler retourne les leçons d'un thème.
func ListLessonsByThemeHandler(repo domain.LessonRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		lessons, err := repo.FindByTheme(c.Param("theme"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, lessons)
	}
}

// ListLessonsByLicenseHandler retourne les leçons filtrées par licence.
func ListLessonsByLicenseHandler(repo domain.LessonRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		license := domain.License(c.Param("license"))
		lessons, err := repo.FindByLicense(license)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, lessons)
	}
}

// ListLessonsByCategoryHandler retourne les leçons filtrées par catégorie.
func ListLessonsByCategoryHandler(repo domain.LessonRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		category := domain.Category(c.Param("category"))
		lessons, err := repo.FindByCategory(category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, lessons)
	}
}


// ======================== QUIZ DE LEÇON ========================

// LessonQuizHandler génère un quiz de 3 questions pour une leçon.
func LessonQuizHandler(engine *Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		lessonID := c.Param("id")

		questions, err := engine.LessonQuiz(lessonID, QuizSize)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		// On ne renvoie pas les réponses
		var safeQuestions []gin.H
		for _, q := range questions {
			safeQuestions = append(safeQuestions, gin.H{
				"id":          q.ID,
				"theme":       q.Theme,
				"difficulty":  q.Difficulty,
				"question_fr": q.QuestionFr,
				"question_en": q.QuestionEn,
				"options":     q.Options,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"lesson_id": lessonID,
			"questions": safeQuestions,
		})
	}
}

// ======================== EXAMEN DE THÈME ========================

// StartExamHandler génère un examen complet sur un thème.
func StartExamHandler(engine *Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		studentID := c.GetString("student_id")
		theme := c.Param("theme")

		questions, err := engine.ThemeExam(studentID, theme, ExamSize)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		// On ne renvoie pas les réponses
		var safeQuestions []gin.H
		for _, q := range questions {
			safeQuestions = append(safeQuestions, gin.H{
				"id":          q.ID,
				"theme":       q.Theme,
				"difficulty":  q.Difficulty,
				"question_fr": q.QuestionFr,
				"question_en": q.QuestionEn,
				"options":     q.Options,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"theme":     theme,
			"total":     len(questions),
			"questions": safeQuestions,
		})
	}
}

// SubmitExamHandler soumet toutes les réponses d'un examen et retourne la correction.
func SubmitExamHandler(engine *Engine, historyRepo domain.QuestionHistoryRepository, questionRepo domain.QuestionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		studentID := c.GetString("student_id")

		var req submitExamRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		type result struct {
			QuestionID    string `json:"question_id"`
			Correct       bool   `json:"correct"`
			CorrectAnswer string `json:"correct_answer"`
			ExplanationFr string `json:"explanation_fr,omitempty"`
			ExplanationEn string `json:"explanation_en,omitempty"`
		}

		var results []result
		correctCount := 0

		for _, ans := range req.Answers {
			question, err := questionRepo.FindByID(ans.QuestionID)
			if err != nil {
				continue
			}

			wasCorrect := engine.EvaluateAnswer(question, ans.Answer)

			// Enregistrer dans l'historique
			_ = historyRepo.RecordAnswer(studentID, ans.QuestionID, wasCorrect)

			if wasCorrect {
				correctCount++
			}

			results = append(results, result{
				QuestionID:    ans.QuestionID,
				Correct:       wasCorrect,
				CorrectAnswer: question.AnswerKey,
				ExplanationFr: question.ExplanationFr,
				ExplanationEn: question.ExplanationEn,
			})
		}

		total := len(req.Answers)
		score := 0.0
		if total > 0 {
			score = float64(correctCount) / float64(total) * 100
		}

		c.JSON(http.StatusOK, gin.H{
			"score":    score,
			"correct":  correctCount,
			"total":    total,
			"results":  results,
		})
	}
}

// ======================== STATISTIQUES ========================

// StatsHandler retourne les statistiques de l'étudiant.
func StatsHandler(historyRepo domain.QuestionHistoryRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		studentID := c.GetString("student_id")

		stats, err := historyRepo.GetStats(studentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, stats)
	}
}
