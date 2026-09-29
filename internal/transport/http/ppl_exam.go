package http

import (
	"aeropath/internal/learning"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PPLExamPapersHandler() gin.HandlerFunc {
	return func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"papers": learning.PPLPapers, "pass_percent": 75}) }
}

func PPLExamHandler(engine *learning.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		paper, questions, err := engine.PPLExam(c.Param("subject"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		safe := make([]gin.H, 0, len(questions))
		for _, q := range questions {
			safe = append(safe, gin.H{"id": q.ID, "lesson_id": q.LessonID, "license": q.License, "category": q.Category,
				"question_fr": q.QuestionFr, "question_en": q.QuestionEn, "question_es": q.QuestionEs,
				"options": q.Options, "options_en": q.OptionsEn, "options_es": q.OptionsEs})
		}
		c.JSON(http.StatusOK, gin.H{"paper": paper, "questions": safe, "pass_percent": 75})
	}
}
