package admin

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 📖 DDIA Chapitre 11 : "Stream Processing"
//    L'admin handler expose des endpoints pour la supervision
//    et la gestion de la plateforme. Les opérations d'écriture
//    (création/modification/suppression) sont journalisées
//    dans NATS pour traçabilité.

// --- Types ---

type Student struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Lang      string    `json:"lang"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Question struct {
	ID            string   `json:"id"`
	QuestionFr    string   `json:"question_fr"`
	QuestionEn    string   `json:"question_en"`
	Options       []string `json:"options"`
	AnswerKey     string   `json:"answer_key"`
	ExplanationFr string   `json:"explanation_fr"`
	ExplanationEn string   `json:"explanation_en"`
	License       string   `json:"license"`
	Category      string   `json:"category"`
	Theme         string   `json:"theme"`
	Subtopic      string   `json:"subtopic"`
	Difficulty    int      `json:"difficulty"`
	Reference     string   `json:"reference"`
}

type Lesson struct {
	ID         string `json:"id"`
	TitleFr    string `json:"title_fr"`
	TitleEn    string `json:"title_en"`
	ContentFr  string `json:"content_fr"`
	ContentEn  string `json:"content_en"`
	License    string `json:"license"`
	Category   string `json:"category"`
	Theme      string `json:"theme"`
	Difficulty int    `json:"difficulty"`
	OrderIndex int    `json:"order_index"`
}

type Stats struct {
	TotalStudents       int                `json:"total_students"`
	TotalQuestions      int                `json:"total_questions"`
	TotalLessons        int                `json:"total_lessons"`
	TotalAnswers        int                `json:"total_answers"`
	CorrectAnswers      int                `json:"correct_answers"`
	GlobalAccuracy      float64            `json:"global_accuracy"`
	QuestionsByLicense  map[string]int     `json:"questions_by_license"`
	QuestionsByCategory map[string]int     `json:"questions_by_category"`
	StudentsByLang      map[string]int     `json:"students_by_lang"`
	AnswersLast7Days    []DailyCount       `json:"answers_last_7_days"`
	RegistrationsByDay  []DailyCount       `json:"registrations_by_day"`
	TopWeakTopics       []TopicStat        `json:"top_weak_topics"`
}

type DailyCount struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type TopicStat struct {
	Theme string  `json:"theme"`
	Score float64 `json:"score"`
	Count int     `json:"count"`
}

// --- Repository interface ---

type Repository interface {
	CountStudents(ctx context.Context) (int, error)
	CountQuestions(ctx context.Context) (int, error)
	CountLessons(ctx context.Context) (int, error)
	CountAnswers(ctx context.Context) (int, error)
	CountCorrectAnswers(ctx context.Context) (int, error)
	GetQuestionsByLicense(ctx context.Context) (map[string]int, error)
	GetQuestionsByCategory(ctx context.Context) (map[string]int, error)
	GetStudentsByLang(ctx context.Context) (map[string]int, error)
	GetAnswersLast7Days(ctx context.Context) ([]DailyCount, error)
	GetRegistrationsByDay(ctx context.Context, days int) ([]DailyCount, error)
	GetTopWeakTopics(ctx context.Context, limit int) ([]TopicStat, error)
	ListStudents(ctx context.Context, offset, limit int) ([]Student, int, error)
	GetStudentByID(ctx context.Context, id string) (*Student, error)
	DeleteStudent(ctx context.Context, id string) error
	CreateQuestion(ctx context.Context, q *Question) error
	UpdateQuestion(ctx context.Context, q *Question) error
	DeleteQuestion(ctx context.Context, id string) error
	CreateLesson(ctx context.Context, l *Lesson) error
	UpdateLesson(ctx context.Context, l *Lesson) error
	DeleteLesson(ctx context.Context, id string) error
}

// --- Handler ---

type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

// RegisterRoutes enregistre les routes admin sur un groupe Gin
func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	admin := rg.Group("/admin")
	{
		admin.GET("/stats", h.GetStats)
		admin.GET("/students", h.ListStudents)
		admin.GET("/students/:id", h.GetStudent)
		admin.DELETE("/students/:id", h.DeleteStudent)
		admin.POST("/questions", h.CreateQuestion)
		admin.PUT("/questions/:id", h.UpdateQuestion)
		admin.DELETE("/questions/:id", h.DeleteQuestion)
		admin.POST("/lessons", h.CreateLesson)
		admin.PUT("/lessons/:id", h.UpdateLesson)
		admin.DELETE("/lessons/:id", h.DeleteLesson)
	}
}

// --- Handlers ---

// GetStats retourne les statistiques globales de la plateforme
func (h *Handler) GetStats(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		totalStudents, totalQuestions, totalLessons int
		totalAnswers, correctAnswers                int
		questionsByLicense, questionsByCategory      map[string]int
		studentsByLang                               map[string]int
		answersLast7Days                             []DailyCount
		registrationsByDay                           []DailyCount
		topWeakTopics                                []TopicStat
	)

	// Requêtes concurrentes
	type result struct {
		name string
		val  interface{}
		err  error
	}
	ch := make(chan result, 11)

	go func() { v, e := h.repo.CountStudents(ctx); ch <- result{"students", v, e} }()
	go func() { v, e := h.repo.CountQuestions(ctx); ch <- result{"questions", v, e} }()
	go func() { v, e := h.repo.CountLessons(ctx); ch <- result{"lessons", v, e} }()
	go func() { v, e := h.repo.CountAnswers(ctx); ch <- result{"answers", v, e} }()
	go func() { v, e := h.repo.CountCorrectAnswers(ctx); ch <- result{"correct", v, e} }()
	go func() { v, e := h.repo.GetQuestionsByLicense(ctx); ch <- result{"qByLicense", v, e} }()
	go func() { v, e := h.repo.GetQuestionsByCategory(ctx); ch <- result{"qByCategory", v, e} }()
	go func() { v, e := h.repo.GetStudentsByLang(ctx); ch <- result{"studentsByLang", v, e} }()
	go func() { v, e := h.repo.GetAnswersLast7Days(ctx); ch <- result{"answers7d", v, e} }()
	go func() { v, e := h.repo.GetRegistrationsByDay(ctx, 30); ch <- result{"registrations", v, e} }()
	go func() { v, e := h.repo.GetTopWeakTopics(ctx, 10); ch <- result{"weakTopics", v, e} }()

	for i := 0; i < 11; i++ {
		r := <-ch
		if r.err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": r.err.Error()})
			return
		}
		switch r.name {
		case "students":
			totalStudents = r.val.(int)
		case "questions":
			totalQuestions = r.val.(int)
		case "lessons":
			totalLessons = r.val.(int)
		case "answers":
			totalAnswers = r.val.(int)
		case "correct":
			correctAnswers = r.val.(int)
		case "qByLicense":
			questionsByLicense = r.val.(map[string]int)
		case "qByCategory":
			questionsByCategory = r.val.(map[string]int)
		case "studentsByLang":
			studentsByLang = r.val.(map[string]int)
		case "answers7d":
			answersLast7Days = r.val.([]DailyCount)
		case "registrations":
			registrationsByDay = r.val.([]DailyCount)
		case "weakTopics":
			topWeakTopics = r.val.([]TopicStat)
		}
	}

	globalAccuracy := 0.0
	if totalAnswers > 0 {
		globalAccuracy = float64(correctAnswers) / float64(totalAnswers)
	}

	c.JSON(http.StatusOK, Stats{
		TotalStudents:       totalStudents,
		TotalQuestions:      totalQuestions,
		TotalLessons:        totalLessons,
		TotalAnswers:        totalAnswers,
		CorrectAnswers:      correctAnswers,
		GlobalAccuracy:      globalAccuracy,
		QuestionsByLicense:  questionsByLicense,
		QuestionsByCategory: questionsByCategory,
		StudentsByLang:      studentsByLang,
		AnswersLast7Days:    answersLast7Days,
		RegistrationsByDay:  registrationsByDay,
		TopWeakTopics:       topWeakTopics,
	})
}

// ListStudents retourne la liste paginée des étudiants
func (h *Handler) ListStudents(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	students, total, err := h.repo.ListStudents(c.Request.Context(), offset, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := (total + pageSize - 1) / pageSize
	if students == nil {
		students = []Student{}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":        students,
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": totalPages,
	})
}

// GetStudent retourne les détails d'un étudiant
func (h *Handler) GetStudent(c *gin.Context) {
	id := c.Param("id")
	student, err := h.repo.GetStudentByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Étudiant non trouvé"})
		return
	}
	c.JSON(http.StatusOK, student)
}

// DeleteStudent supprime un étudiant
func (h *Handler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DeleteStudent(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Étudiant supprimé"})
}

// CreateQuestion crée une nouvelle question
func (h *Handler) CreateQuestion(c *gin.Context) {
	var q Question
	if err := c.ShouldBindJSON(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.CreateQuestion(c.Request.Context(), &q); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, q)
}

// UpdateQuestion met à jour une question existante
func (h *Handler) UpdateQuestion(c *gin.Context) {
	id := c.Param("id")
	var q Question
	if err := c.ShouldBindJSON(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	q.ID = id
	if err := h.repo.UpdateQuestion(c.Request.Context(), &q); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, q)
}

// DeleteQuestion supprime une question
func (h *Handler) DeleteQuestion(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DeleteQuestion(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Question supprimée"})
}

// CreateLesson crée une nouvelle leçon
func (h *Handler) CreateLesson(c *gin.Context) {
	var l Lesson
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.CreateLesson(c.Request.Context(), &l); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, l)
}

// UpdateLesson met à jour une leçon existante
func (h *Handler) UpdateLesson(c *gin.Context) {
	id := c.Param("id")
	var l Lesson
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	l.ID = id
	if err := h.repo.UpdateLesson(c.Request.Context(), &l); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, l)
}

// DeleteLesson supprime une leçon
func (h *Handler) DeleteLesson(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DeleteLesson(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Leçon supprimée"})
}
