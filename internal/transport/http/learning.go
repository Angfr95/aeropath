package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"aeropath/internal/domain"
	"aeropath/internal/learning"
)

// 📖 DDIA Chapitre 4 : "Encoding and Evolution"
//    Les handlers HTTP sont le point d'entrée de l'API REST.
//    Ils utilisent Gin (framework HTTP) pour router les requêtes.
//    Chaque handler reçoit un "repository" en paramètre (injection de dépendance).
//    C'est le pattern "Handler Func" du Chapitre 2 : on sépare
//    le transport (HTTP) de la logique métier (domain).

type answerRequest struct {
	QuestionID string `json:"question_id" binding:"required,uuid"`
	Answer     string `json:"answer"      binding:"required"`
}

type submitExamRequest struct {
	Answers []answerRequest `json:"answers" binding:"required,min=1"`
}

// ======================== LEÇONS ========================

// ListStudentsHandler
//
//	@Summary		Liste des étudiants
//	@Description	Retourne tous les étudiants (admin)
//	@Tags			Étudiants
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{array}		domain.Student
//	@Router			/api/students [get]
func ListStudentsHandler(repo domain.StudentRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "non implémenté"})
	}
}

// GetStudentHandler retourne un étudiant par son ID.
func GetStudentHandler(repo domain.StudentRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		student, err := repo.FindByID(c.Param("id"))
		if err != nil || student == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "étudiant introuvable"})
			return
		}
		c.JSON(http.StatusOK, student)
	}
}

// UpdateStudentHandler met à jour un étudiant.
// 🔒 SÉCURITÉ : l'étudiant connecté ne peut modifier que son propre profil.
func UpdateStudentHandler(repo domain.StudentRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		studentID := c.GetString("student_id")

		// 🔒 Vérifier que l'étudiant modifie son propre profil
		if studentID != id {
			c.JSON(http.StatusForbidden, gin.H{"error": "vous ne pouvez modifier que votre propre profil"})
			return
		}

		existing, err := repo.FindByID(id)
		if err != nil || existing == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "étudiant introuvable"})
			return
		}

		var updates struct {
			Email *string `json:"email"`
			Lang  *string `json:"lang"`
		}
		if err := c.ShouldBindJSON(&updates); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if updates.Lang != nil {
			if *updates.Lang != "fr" && *updates.Lang != "en" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "langue non supportée (fr ou en)"})
				return
			}
			if err := repo.UpdateLang(id, *updates.Lang); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "étudiant mis à jour", "id": id})
	}
}

// StudentCountHandler retourne le nombre d'étudiants.
func StudentCountHandler(repo domain.StudentRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		count, err := repo.Count()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"count": count})
	}
}

// DeleteStudentHandler supprime un étudiant.
func DeleteStudentHandler(repo domain.StudentRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := repo.Delete(id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "étudiant introuvable"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "étudiant supprimé", "id": id})
	}
}

// CreateLessonHandler crée une nouvelle leçon.
func CreateLessonHandler(repo domain.LessonRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var l domain.Lesson
		if err := c.ShouldBindJSON(&l); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := repo.Create(&l); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": l.ID, "created_at": l.CreatedAt})
	}
}

// UpdateLessonHandler met à jour une leçon.
func UpdateLessonHandler(repo domain.LessonRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		existing, err := repo.FindByID(id)
		if err != nil || existing == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "leçon introuvable"})
			return
		}

		var updates struct {
			TitleFr    *string        `json:"title_fr"`
			TitleEn    *string        `json:"title_en"`
			ContentFr  *string        `json:"content_fr"`
			ContentEn  *string        `json:"content_en"`
			Theme      *string        `json:"theme"`
			Difficulty *int           `json:"difficulty"`
			OrderIndex *int           `json:"order_index"`
			License    *domain.License `json:"license"`
			Category   *domain.Category `json:"category"`
		}
		if err := c.ShouldBindJSON(&updates); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if updates.TitleFr != nil { existing.TitleFr = *updates.TitleFr }
		if updates.TitleEn != nil { existing.TitleEn = *updates.TitleEn }
		if updates.ContentFr != nil { existing.ContentFr = *updates.ContentFr }
		if updates.ContentEn != nil { existing.ContentEn = *updates.ContentEn }
		if updates.Theme != nil { existing.Theme = *updates.Theme }
		if updates.Difficulty != nil { existing.Difficulty = *updates.Difficulty }
		if updates.OrderIndex != nil { existing.OrderIndex = *updates.OrderIndex }
		if updates.License != nil { existing.License = *updates.License }
		if updates.Category != nil { existing.Category = *updates.Category }

		if err := repo.Update(existing); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "leçon mise à jour", "id": id})
	}
}

// DeleteLessonHandler supprime une leçon.
func DeleteLessonHandler(repo domain.LessonRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := repo.Delete(id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "leçon introuvable"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "leçon supprimée", "id": id})
	}
}

// ListLessonsHandler retourne toutes les leçons avec pagination.
func ListLessonsHandler(repo domain.LessonRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		license := c.Query("license")
		category := c.Query("category")

		// Pagination
		limit := 20
		offset := 0
		if l := c.Query("limit"); l != "" {
			fmt.Sscanf(l, "%d", &limit)
		}
		if o := c.Query("offset"); o != "" {
			fmt.Sscanf(o, "%d", &offset)
		}
		if limit < 1 { limit = 1 }
		if limit > 100 { limit = 100 }
		if offset < 0 { offset = 0 }

		var lessons []*domain.Lesson
		var total int
		var err error

		switch {
		case license != "" && category != "":
			lessons, err = repo.FindByLicenseAndCategory(domain.License(license), domain.Category(category))
			total = len(lessons)
		case license != "":
			lessons, err = repo.FindByLicense(domain.License(license))
			total = len(lessons)
		case category != "":
			lessons, err = repo.FindByCategory(domain.Category(category))
			total = len(lessons)
		default:
			total, err = repo.Count()
			if err == nil {
				lessons, err = repo.FindAllPaginated(limit, offset)
			}
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"total":  total,
			"limit":  limit,
			"offset": offset,
			"data":   lessons,
		})
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

// ListLessonsByLicenseAndCategoryHandler retourne les leçons filtrées par licence et catégorie.
func ListLessonsByLicenseAndCategoryHandler(repo domain.LessonRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		license := domain.License(c.Param("license"))
		category := domain.Category(c.Param("category"))
		lessons, err := repo.FindByLicenseAndCategory(license, category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, lessons)
	}
}

// ListLessonsByDifficultyHandler retourne les leçons filtrées par difficulté.
func ListLessonsByDifficultyHandler(repo domain.LessonRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var difficulty int
		fmt.Sscanf(c.Param("level"), "%d", &difficulty)
		lessons, err := repo.FindByDifficulty(difficulty)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, lessons)
	}
}

// LessonCountHandler retourne le nombre total de leçons.
func LessonCountHandler(repo domain.LessonRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		count, err := repo.Count()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"count": count})
	}
}

// LessonCountByLicenseHandler retourne le nombre de leçons par licence.
func LessonCountByLicenseHandler(repo domain.LessonRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		license := domain.License(c.Param("license"))
		count, err := repo.CountByLicense(license)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"license": license, "count": count})
	}
}

// LessonCountByCategoryHandler retourne le nombre de leçons par catégorie.
func LessonCountByCategoryHandler(repo domain.LessonRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		category := domain.Category(c.Param("category"))
		count, err := repo.CountByCategory(category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"category": category, "count": count})
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
func LessonQuizHandler(engine *learning.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		lessonID := c.Param("id")

		questions, err := engine.LessonQuiz(lessonID, learning.QuizSize)
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
func StartExamHandler(engine *learning.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		studentID := c.GetString("student_id")
		theme := c.Param("theme")

		questions, err := engine.ThemeExam(studentID, theme, learning.ExamSize)
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
func SubmitExamHandler(engine *learning.Engine, historyRepo domain.QuestionHistoryRepository, questionRepo domain.QuestionRepository) gin.HandlerFunc {
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
			"score":   score,
			"correct": correctCount,
			"total":   total,
			"results": results,
		})
	}
}

// ======================== STATISTIQUES ========================

// ======================== QUESTIONS ========================

// ListQuestionsHandler retourne toutes les questions avec filtres optionnels.
func ListQuestionsHandler(repo domain.QuestionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		theme := c.Query("theme")
		license := c.Query("license")
		category := c.Query("category")
		difficulty := c.Query("difficulty")

		var questions []*domain.Question
		var err error

		switch {
		case theme != "":
			questions, err = repo.FindByTheme(theme)
		case license != "" && category != "":
			questions, err = repo.FindByLicenseAndCategory(domain.License(license), domain.Category(category))
		case license != "":
			questions, err = repo.FindByLicense(domain.License(license))
		case category != "":
			questions, err = repo.FindByCategory(domain.Category(category))
		default:
			questions, err = repo.FindAll()
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Filtrer par difficulté si spécifié
		if difficulty != "" {
			var filtered []*domain.Question
			for _, q := range questions {
				if fmt.Sprintf("%d", q.Difficulty) == difficulty {
					filtered = append(filtered, q)
				}
			}
			questions = filtered
		}

		// Ne pas exposer les réponses dans la liste
		var safeQuestions []gin.H
		for _, q := range questions {
			safeQuestions = append(safeQuestions, gin.H{
				"id":          q.ID,
				"lesson_id":   q.LessonID,
				"license":     q.License,
				"category":    q.Category,
				"theme":       q.Theme,
				"subtopic":    q.Subtopic,
				"difficulty":  q.Difficulty,
				"question_fr": q.QuestionFr,
				"question_en": q.QuestionEn,
				"options":     q.Options,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"total":     len(safeQuestions),
			"questions": safeQuestions,
		})
	}
}

// QuestionsByThemeHandler retourne les questions filtrées par thème.
func QuestionsByThemeHandler(repo domain.QuestionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		theme := c.Param("theme")
		questions, err := repo.FindByTheme(theme)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var safeQuestions []gin.H
		for _, q := range questions {
			safeQuestions = append(safeQuestions, gin.H{
				"id":          q.ID,
				"lesson_id":   q.LessonID,
				"license":     q.License,
				"category":    q.Category,
				"theme":       q.Theme,
				"subtopic":    q.Subtopic,
				"difficulty":  q.Difficulty,
				"question_fr": q.QuestionFr,
				"question_en": q.QuestionEn,
				"options":     q.Options,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"theme":     theme,
			"total":     len(safeQuestions),
			"questions": safeQuestions,
		})
	}
}

// QuestionsByCategoryHandler retourne les questions filtrées par catégorie.
func QuestionsByCategoryHandler(repo domain.QuestionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		category := domain.Category(c.Param("category"))
		questions, err := repo.FindByCategory(category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var safeQuestions []gin.H
		for _, q := range questions {
			safeQuestions = append(safeQuestions, gin.H{
				"id":          q.ID,
				"lesson_id":   q.LessonID,
				"license":     q.License,
				"category":    q.Category,
				"theme":       q.Theme,
				"subtopic":    q.Subtopic,
				"difficulty":  q.Difficulty,
				"question_fr": q.QuestionFr,
				"question_en": q.QuestionEn,
				"options":     q.Options,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"category":  category,
			"total":     len(safeQuestions),
			"questions": safeQuestions,
		})
	}
}

// QuestionsByLicenseAndCategoryHandler retourne les questions filtrées par licence et catégorie.
func QuestionsByLicenseAndCategoryHandler(repo domain.QuestionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		license := domain.License(c.Param("license"))
		category := domain.Category(c.Param("category"))
		questions, err := repo.FindByLicenseAndCategory(license, category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var safeQuestions []gin.H
		for _, q := range questions {
			safeQuestions = append(safeQuestions, gin.H{
				"id":          q.ID,
				"lesson_id":   q.LessonID,
				"license":     q.License,
				"category":    q.Category,
				"theme":       q.Theme,
				"subtopic":    q.Subtopic,
				"difficulty":  q.Difficulty,
				"question_fr": q.QuestionFr,
				"question_en": q.QuestionEn,
				"options":     q.Options,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"license":   license,
			"category":  category,
			"total":     len(safeQuestions),
			"questions": safeQuestions,
		})
	}
}

// CreateQuestionHandler crée une nouvelle question.
func CreateQuestionHandler(repo domain.QuestionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var q domain.Question
		if err := c.ShouldBindJSON(&q); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := repo.Create(&q); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"id":         q.ID,
			"created_at": q.CreatedAt,
		})
	}
}

// UpdateQuestionHandler met à jour une question existante.
func UpdateQuestionHandler(repo domain.QuestionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		existing, err := repo.FindByID(id)
		if err != nil || existing == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "question introuvable"})
			return
		}

		var updates struct {
			QuestionFr    *string          `json:"question_fr"`
			QuestionEn    *string          `json:"question_en"`
			Options       []string         `json:"options"`
			AnswerKey     *string          `json:"answer_key"`
			Difficulty    *int             `json:"difficulty"`
			Theme         *string          `json:"theme"`
			Subtopic      *string          `json:"subtopic"`
			License       *domain.License  `json:"license"`
			Category      *domain.Category `json:"category"`
			ExplanationFr *string          `json:"explanation_fr"`
			ExplanationEn *string          `json:"explanation_en"`
		}
		if err := c.ShouldBindJSON(&updates); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if updates.QuestionFr != nil { existing.QuestionFr = *updates.QuestionFr }
		if updates.QuestionEn != nil { existing.QuestionEn = *updates.QuestionEn }
		if updates.Options != nil { existing.Options = updates.Options }
		if updates.AnswerKey != nil { existing.AnswerKey = *updates.AnswerKey }
		if updates.Difficulty != nil { existing.Difficulty = *updates.Difficulty }
		if updates.Theme != nil { existing.Theme = *updates.Theme }
		if updates.Subtopic != nil { existing.Subtopic = *updates.Subtopic }
		if updates.License != nil { existing.License = *updates.License }
		if updates.Category != nil { existing.Category = *updates.Category }
		if updates.ExplanationFr != nil { existing.ExplanationFr = *updates.ExplanationFr }
		if updates.ExplanationEn != nil { existing.ExplanationEn = *updates.ExplanationEn }

		if err := repo.Update(existing); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "question mise à jour", "id": id})
	}
}

// DeleteQuestionHandler supprime une question.
func DeleteQuestionHandler(repo domain.QuestionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := repo.Delete(id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "question introuvable"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "question supprimée", "id": id})
	}
}

// QuestionsByDifficultyHandler retourne les questions filtrées par difficulté.
func QuestionsByDifficultyHandler(repo domain.QuestionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var difficulty int
		fmt.Sscanf(c.Param("level"), "%d", &difficulty)
		questions, err := repo.FindByDifficulty(difficulty)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var safeQuestions []gin.H
		for _, q := range questions {
			safeQuestions = append(safeQuestions, gin.H{
				"id": q.ID, "lesson_id": q.LessonID, "license": q.License,
				"category": q.Category, "theme": q.Theme, "subtopic": q.Subtopic,
				"difficulty": q.Difficulty, "question_fr": q.QuestionFr,
				"question_en": q.QuestionEn, "options": q.Options,
			})
		}
		c.JSON(http.StatusOK, gin.H{"difficulty": difficulty, "total": len(safeQuestions), "questions": safeQuestions})
	}
}

// QuestionsBySubtopicHandler retourne les questions filtrées par sous-thème.
func QuestionsBySubtopicHandler(repo domain.QuestionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		subtopic := c.Param("subtopic")
		questions, err := repo.FindBySubtopic(subtopic)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var safeQuestions []gin.H
		for _, q := range questions {
			safeQuestions = append(safeQuestions, gin.H{
				"id": q.ID, "lesson_id": q.LessonID, "license": q.License,
				"category": q.Category, "theme": q.Theme, "subtopic": q.Subtopic,
				"difficulty": q.Difficulty, "question_fr": q.QuestionFr,
				"question_en": q.QuestionEn, "options": q.Options,
			})
		}
		c.JSON(http.StatusOK, gin.H{"subtopic": subtopic, "total": len(safeQuestions), "questions": safeQuestions})
	}
}

// QuestionCountHandler retourne le nombre total de questions.
func QuestionCountHandler(repo domain.QuestionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		count, err := repo.Count()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"count": count})
	}
}

// QuestionCountByLicenseHandler retourne le nombre de questions par licence.
func QuestionCountByLicenseHandler(repo domain.QuestionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		license := domain.License(c.Param("license"))
		count, err := repo.CountByLicense(license)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"license": license, "count": count})
	}
}

// QuestionCountByCategoryHandler retourne le nombre de questions par catégorie.
func QuestionCountByCategoryHandler(repo domain.QuestionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		category := domain.Category(c.Param("category"))
		count, err := repo.CountByCategory(category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"category": category, "count": count})
	}
}

// QuestionCountByThemeHandler retourne le nombre de questions par thème.
func QuestionCountByThemeHandler(repo domain.QuestionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		theme := c.Param("theme")
		count, err := repo.CountByTheme(theme)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"theme": theme, "count": count})
	}
}

// QuestionsByLicenseHandler retourne les questions filtrées par licence.
func QuestionsByLicenseHandler(repo domain.QuestionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		license := domain.License(c.Param("license"))
		questions, err := repo.FindByLicense(license)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var safeQuestions []gin.H
		for _, q := range questions {
			safeQuestions = append(safeQuestions, gin.H{
				"id":          q.ID,
				"lesson_id":   q.LessonID,
				"license":     q.License,
				"category":    q.Category,
				"theme":       q.Theme,
				"subtopic":    q.Subtopic,
				"difficulty":  q.Difficulty,
				"question_fr": q.QuestionFr,
				"question_en": q.QuestionEn,
				"options":     q.Options,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"license":   license,
			"total":     len(safeQuestions),
			"questions": safeQuestions,
		})
	}
}

// RandomQuestionHandler retourne une ou plusieurs questions aléatoires.
// Query params: limit (défaut: 1), license, category, theme (optionnels)
func RandomQuestionHandler(repo domain.QuestionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit := 1
		if l := c.Query("limit"); l != "" {
			fmt.Sscanf(l, "%d", &limit)
		}
		if limit < 1 {
			limit = 1
		}
		if limit > 50 {
			limit = 50
		}

		var license *domain.License
		if l := c.Query("license"); l != "" {
			v := domain.License(l)
			license = &v
		}
		var category *domain.Category
		if cat := c.Query("category"); cat != "" {
			v := domain.Category(cat)
			category = &v
		}
		var theme *string
		if t := c.Query("theme"); t != "" {
			theme = &t
		}

		questions, err := repo.FindRandom(limit, license, category, theme)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if len(questions) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "aucune question trouvée"})
			return
		}

		// Ne pas exposer les réponses
		var safeQuestions []gin.H
		for _, q := range questions {
			safeQuestions = append(safeQuestions, gin.H{
				"id":          q.ID,
				"lesson_id":   q.LessonID,
				"license":     q.License,
				"category":    q.Category,
				"theme":       q.Theme,
				"subtopic":    q.Subtopic,
				"difficulty":  q.Difficulty,
				"question_fr": q.QuestionFr,
				"question_en": q.QuestionEn,
				"options":     q.Options,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"total":     len(safeQuestions),
			"questions": safeQuestions,
		})
	}
}

// SearchQuestionsHandler cherche des questions par texte libre.
// Utilise la recherche full-text PostgreSQL (tsvector) + ILIKE.
// Query params: q (requis), license, category, difficulty (optionnels)
func SearchQuestionsHandler(repo domain.QuestionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Query("q")
		if query == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "paramètre 'q' requis"})
			return
		}

		questions, err := repo.Search(query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Filtres supplémentaires côté serveur
		license := c.Query("license")
		category := c.Query("category")
		difficulty := c.Query("difficulty")

		var filtered []*domain.Question
		for _, q := range questions {
			if license != "" && string(q.License) != license {
				continue
			}
			if category != "" && string(q.Category) != category {
				continue
			}
			if difficulty != "" && fmt.Sprintf("%d", q.Difficulty) != difficulty {
				continue
			}
			filtered = append(filtered, q)
		}

		// Ne pas exposer les réponses
		var safeQuestions []gin.H
		for _, q := range filtered {
			safeQuestions = append(safeQuestions, gin.H{
				"id":          q.ID,
				"lesson_id":   q.LessonID,
				"license":     q.License,
				"category":    q.Category,
				"theme":       q.Theme,
				"subtopic":    q.Subtopic,
				"difficulty":  q.Difficulty,
				"question_fr": q.QuestionFr,
				"question_en": q.QuestionEn,
				"options":     q.Options,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"query":     query,
			"total":     len(safeQuestions),
			"questions": safeQuestions,
		})
	}
}

// AnswerQuestionHandler vérifie une réponse à une question et enregistre dans l'historique.
// Body: {"question_id": "...", "answer": "A"}
func AnswerQuestionHandler(engine *learning.Engine, historyRepo domain.QuestionHistoryRepository, questionRepo domain.QuestionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		studentID := c.GetString("student_id")

		var req struct {
			QuestionID string `json:"question_id" binding:"required"`
			Answer     string `json:"answer" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "question_id et answer requis"})
			return
		}

		question, err := questionRepo.FindByID(req.QuestionID)
		if err != nil || question == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "question introuvable"})
			return
		}

		wasCorrect := engine.EvaluateAnswer(question, req.Answer)

		// Enregistrer dans l'historique
		_ = historyRepo.RecordAnswer(studentID, req.QuestionID, wasCorrect)

		c.JSON(http.StatusOK, gin.H{
			"question_id":    req.QuestionID,
			"correct":        wasCorrect,
			"correct_answer": question.AnswerKey,
			"explanation_fr": question.ExplanationFr,
			"explanation_en": question.ExplanationEn,
		})
	}
}

// GetQuestionHandler retourne une question par son ID (sans la réponse).
func GetQuestionHandler(repo domain.QuestionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		question, err := repo.FindByID(c.Param("id"))
		if err != nil || question == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "question introuvable"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":            question.ID,
			"lesson_id":     question.LessonID,
			"license":       question.License,
			"category":      question.Category,
			"theme":         question.Theme,
			"subtopic":      question.Subtopic,
			"difficulty":    question.Difficulty,
			"question_fr":   question.QuestionFr,
			"question_en":   question.QuestionEn,
			"options":       question.Options,
			"explanation_fr": question.ExplanationFr,
			"explanation_en": question.ExplanationEn,
		})
	}
}

// HistoryHandler retourne l'historique des réponses de l'étudiant.
// Query params: limit (défaut: 50), offset (défaut: 0), correct (true/false, optionnel)
func HistoryHandler(historyRepo domain.QuestionHistoryRepository, questionRepo domain.QuestionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		studentID := c.GetString("student_id")

		history, err := historyRepo.GetHistory(studentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Filtre optionnel par correct/incorrect
		correctFilter := c.Query("correct")
		if correctFilter != "" {
			var filtered []*domain.QuestionHistory
			wantCorrect := correctFilter == "true"
			for _, h := range history {
				if h.WasCorrect == wantCorrect {
					filtered = append(filtered, h)
				}
			}
			history = filtered
		}

		// Pagination
		limit := 50
		if l := c.Query("limit"); l != "" {
			fmt.Sscanf(l, "%d", &limit)
		}
		offset := 0
		if o := c.Query("offset"); o != "" {
			fmt.Sscanf(o, "%d", &offset)
		}

		total := len(history)

		// Trier du plus récent au plus ancien
		for i := 0; i < len(history); i++ {
			for j := i + 1; j < len(history); j++ {
				if history[j].SeenAt.After(history[i].SeenAt) {
					history[i], history[j] = history[j], history[i]
				}
			}
		}

		if offset > len(history) {
			offset = len(history)
		}
		end := offset + limit
		if end > len(history) {
			end = len(history)
		}
		page := history[offset:end]

		// Enrichir avec les questions (sans la réponse)
		type HistoryEntry struct {
			QuestionID  string   `json:"question_id"`
			Theme       string   `json:"theme"`
			Subtopic    string   `json:"subtopic,omitempty"`
			Difficulty  int      `json:"difficulty"`
			QuestionFr  string   `json:"question_fr"`
			QuestionEn  string   `json:"question_en"`
			Options     []string `json:"options"`
			WasCorrect  bool     `json:"was_correct"`
			SeenAt      time.Time `json:"seen_at"`
		}

		var entries []HistoryEntry
		for _, h := range page {
			q, err := questionRepo.FindByID(h.QuestionID)
			if err != nil {
				continue
			}
			entries = append(entries, HistoryEntry{
				QuestionID: h.QuestionID,
				Theme:      q.Theme,
				Subtopic:   q.Subtopic,
				Difficulty: q.Difficulty,
				QuestionFr: q.QuestionFr,
				QuestionEn: q.QuestionEn,
				Options:    q.Options,
				WasCorrect: h.WasCorrect,
				SeenAt:     h.SeenAt,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"total":   total,
			"limit":   limit,
			"offset":  offset,
			"entries": entries,
		})
	}
}

// AdminStatsHandler retourne les statistiques globales pour l'administration.
func AdminStatsHandler(studentRepo domain.StudentRepository, questionRepo domain.QuestionRepository, historyRepo domain.QuestionHistoryRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		studentCount, err := studentRepo.Count()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "impossible de compter les étudiants"})
			return
		}

		questionCount, err := questionRepo.Count()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "impossible de compter les questions"})
			return
		}

		answerCount, err := historyRepo.Count()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "impossible de compter les réponses"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"students":  studentCount,
			"questions": questionCount,
			"answers":   answerCount,
		})
	}
}

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
