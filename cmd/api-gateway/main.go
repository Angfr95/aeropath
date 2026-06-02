package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"aeropath/internal/auth"
	"aeropath/internal/domain"
	"aeropath/internal/learning"
	"aeropath/internal/monitoring"
	"aeropath/internal/persistence/postgres"
	"aeropath/internal/recommendation"
	httptransport "aeropath/internal/transport/http"
	"aeropath/internal/transport/middleware"
	wstransport "aeropath/internal/transport/ws"
)

func main() {
	_ = godotenv.Load()

	pool, err := postgres.NewPool(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("postgres: %v", err)
	}
	defer pool.Close()

	// Repositories
	studentRepo := postgres.NewStudentRepo(pool)
	questionRepo := postgres.NewQuestionRepo(pool)
	historyRepo := postgres.NewHistoryRepo(pool)
	lessonRepo := postgres.NewLessonRepo(pool)

	// Services
	authSvc := auth.NewService(studentRepo, os.Getenv("JWT_SECRET"))
	engine := learning.NewEngine(questionRepo, historyRepo)
	adaptiveEngine := recommendation.New(questionRepo, historyRepo, lessonRepo)

	// WebSocket Hub
	wsHub := wstransport.NewHub()
	go wsHub.Run()

	r := gin.Default()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Fichiers statiques (PWA — build Vite)
	r.StaticFile("/manifest.json", "./web/manifest.json")
	r.StaticFile("/manifest.webmanifest", "./web/manifest.webmanifest")
	r.StaticFile("/sw.js", "./web/sw.js")
	r.StaticFile("/registerSW.js", "./web/registerSW.js")
	r.StaticFile("/favicon.svg", "./web/favicon.svg")
	r.Static("/assets/", "./web/assets/")
	r.StaticFile("/workbox-9c191d2f.js", "./web/workbox-9c191d2f.js")
	// SPA - toutes les routes non-API renvoient index.html (React Router)
	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method == "GET" && !isAPIRoute(c.Request.URL.Path) {
			c.File("./web/index.html")
			return
		}
		c.JSON(http.StatusNotFound, gin.H{
			"error": "route introuvable",
		})
	})

	// Routes publiques
	r.GET("/health", monitoring.HealthHandler())

	r.POST("/auth/register", httptransport.RegisterHandler(authSvc))
	r.POST("/auth/login", httptransport.LoginHandler(authSvc))

	// Routes protégées (nécessitent un JWT valide)
	protected := r.Group("/api", middleware.RequireAuth(os.Getenv("JWT_SECRET")))
	{
		// Profil
		protected.GET("/me", func(c *gin.Context) {
			student, err := studentRepo.FindByID(c.GetString("student_id"))
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			c.JSON(200, student)
		})
		protected.PATCH("/me/lang", httptransport.UpdateLangHandler(authSvc))

		// Étudiants (admin)
		protected.GET("/students", httptransport.ListStudentsHandler(studentRepo))
		protected.GET("/students/count", httptransport.StudentCountHandler(studentRepo))
		protected.GET("/students/:id", httptransport.GetStudentHandler(studentRepo))
		protected.PUT("/students/:id", httptransport.UpdateStudentHandler(studentRepo))
		protected.DELETE("/students/:id", httptransport.DeleteStudentHandler(studentRepo))

		// Leçons
		protected.GET("/lessons", httptransport.ListLessonsHandler(lessonRepo))
		protected.GET("/lessons/count", httptransport.LessonCountHandler(lessonRepo))
		protected.GET("/lessons/count/by-license/:license", httptransport.LessonCountByLicenseHandler(lessonRepo))
		protected.GET("/lessons/count/by-category/:category", httptransport.LessonCountByCategoryHandler(lessonRepo))
		protected.GET("/lessons/theme/:theme", httptransport.ListLessonsByThemeHandler(lessonRepo))
		protected.GET("/lessons/license/:license", httptransport.ListLessonsByLicenseHandler(lessonRepo))
		protected.GET("/lessons/category/:category", httptransport.ListLessonsByCategoryHandler(lessonRepo))
		protected.GET("/lessons/by-license/:license/category/:category", httptransport.ListLessonsByLicenseAndCategoryHandler(lessonRepo))
		protected.GET("/lessons/by-difficulty/:level", httptransport.ListLessonsByDifficultyHandler(lessonRepo))
		protected.GET("/lessons/:id", httptransport.GetLessonHandler(lessonRepo))
		protected.POST("/lessons", httptransport.CreateLessonHandler(lessonRepo))
		protected.PUT("/lessons/:id", httptransport.UpdateLessonHandler(lessonRepo))
		protected.DELETE("/lessons/:id", httptransport.DeleteLessonHandler(lessonRepo))

		// Quiz de leçon (3 questions pour vérifier la compréhension)
		protected.GET("/lessons/:id/quiz", httptransport.LessonQuizHandler(engine))

		// Examen de thème (10 questions)
		protected.GET("/exam/:theme", httptransport.StartExamHandler(engine))
		protected.POST("/exam/submit", httptransport.SubmitExamHandler(engine, historyRepo, questionRepo))

		// Examen par licence et catégorie
		protected.GET("/exam/license/:license/category/:category", func(c *gin.Context) {
			license := domain.License(c.Param("license"))
			category := domain.Category(c.Param("category"))
			questions, err := engine.ExamByLicenseAndCategory(license, category, learning.ExamSize)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			var safeQuestions []gin.H
			for _, q := range questions {
				safeQuestions = append(safeQuestions, gin.H{
					"id":          q.ID,
					"license":     q.License,
					"category":    q.Category,
					"theme":       q.Theme,
					"difficulty":  q.Difficulty,
					"question_fr": q.QuestionFr,
					"question_en": q.QuestionEn,
					"options":     q.Options,
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"license":   license,
				"category":  category,
				"total":     len(questions),
				"questions": safeQuestions,
			})
		})

		// Questions
		protected.GET("/questions", httptransport.ListQuestionsHandler(questionRepo))
		protected.GET("/questions/count", httptransport.QuestionCountHandler(questionRepo))
		protected.GET("/questions/count/by-license/:license", httptransport.QuestionCountByLicenseHandler(questionRepo))
		protected.GET("/questions/count/by-category/:category", httptransport.QuestionCountByCategoryHandler(questionRepo))
		protected.GET("/questions/count/by-theme/:theme", httptransport.QuestionCountByThemeHandler(questionRepo))
		protected.GET("/questions/by-license/:license", httptransport.QuestionsByLicenseHandler(questionRepo))
		protected.GET("/questions/by-category/:category", httptransport.QuestionsByCategoryHandler(questionRepo))
		protected.GET("/questions/by-theme/:theme", httptransport.QuestionsByThemeHandler(questionRepo))
		protected.GET("/questions/by-license/:license/category/:category", httptransport.QuestionsByLicenseAndCategoryHandler(questionRepo))
		protected.GET("/questions/by-difficulty/:level", httptransport.QuestionsByDifficultyHandler(questionRepo))
		protected.GET("/questions/by-subtopic/:subtopic", httptransport.QuestionsBySubtopicHandler(questionRepo))
		protected.GET("/questions/random", httptransport.RandomQuestionHandler(questionRepo))
		protected.GET("/questions/search", httptransport.SearchQuestionsHandler(questionRepo))
		protected.GET("/questions/:id", httptransport.GetQuestionHandler(questionRepo))
		protected.POST("/questions", httptransport.CreateQuestionHandler(questionRepo))
		protected.PUT("/questions/:id", httptransport.UpdateQuestionHandler(questionRepo))
		protected.DELETE("/questions/:id", httptransport.DeleteQuestionHandler(questionRepo))
		protected.POST("/questions/answer", httptransport.AnswerQuestionHandler(engine, historyRepo, questionRepo))

		// Historique
		protected.GET("/history", httptransport.HistoryHandler(historyRepo, questionRepo))

		// Statistiques
		protected.GET("/stats", httptransport.StatsHandler(historyRepo))

		// Administration
		protected.GET("/admin/stats", httptransport.AdminStatsHandler(studentRepo, questionRepo, historyRepo))

		// Recommandations adaptatives
		protected.GET("/recommendations", httptransport.RecommendationHandler(adaptiveEngine))
		protected.GET("/recommendations/weak-topics", httptransport.WeakTopicsHandler(adaptiveEngine))
		protected.GET("/recommendations/progression", httptransport.ProgressionHandler(adaptiveEngine))
		protected.GET("/recommendations/due-cards", httptransport.DueCardsHandler(adaptiveEngine))

		// WebSocket temps réel
		protected.GET("/ws", wstransport.ServeWS(wsHub))
		protected.GET("/ws/info", wstransport.WsInfoHandler(wsHub))
	}

	log.Printf("démarré sur :%s\n", os.Getenv("PORT"))
	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("serveur arrêté: %v", err)
	}
}

// isAPIRoute vérifie si le chemin correspond à une route API.
func isAPIRoute(path string) bool {
	apiPrefixes := []string{
		"/api/",
		"/auth/",
		"/health",
		"/swagger/",
		"/manifest.json",
		"/manifest.webmanifest",
		"/sw.js",
		"/registerSW.js",
		"/favicon.svg",
		"/assets/",
		"/workbox-",
	}
	for _, prefix := range apiPrefixes {
		if path == prefix || (len(path) >= len(prefix) && path[:len(prefix)] == prefix) {
			return true
		}
	}
	return false
}
