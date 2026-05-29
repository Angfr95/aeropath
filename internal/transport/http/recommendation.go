package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"aeropath/internal/recommendation"
)

// 📖 DDIA Chapitre 12 : "The Future of Data Systems"
//    Les handlers de recommandation sont des "read-side" handlers
//    dans le pattern CQRS (Command Query Responsibility Segregation).
//    Ils ne modifient pas les données, ils les lisent et les transforment.
//    La logique métier est dans le package recommendation/.

// RecommendationHandler retourne les recommandations personnalisées pour l'étudiant.
//
// 📖 DDIA Chapitre 1 : "Reliability"
//    Ce handler est "stateless" : il ne stocke rien en mémoire.
//    On peut avoir plusieurs instances sans problème.
//    Si une instance plante, une autre prend le relais.
func RecommendationHandler(engine *recommendation.AdaptiveEngine) gin.HandlerFunc {
	return func(c *gin.Context) {
		studentID := c.GetString("student_id")

		rec, err := engine.GetRecommendations(studentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, rec)
	}
}

// WeakTopicsHandler retourne les sujets faibles de l'étudiant.
func WeakTopicsHandler(engine *recommendation.AdaptiveEngine) gin.HandlerFunc {
	return func(c *gin.Context) {
		studentID := c.GetString("student_id")

		rec, err := engine.GetRecommendations(studentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"weak_topics": rec.WeakTopics,
			"next_theme":  rec.NextTheme,
		})
	}
}

// ProgressionHandler retourne la progression de l'étudiant.
func ProgressionHandler(engine *recommendation.AdaptiveEngine) gin.HandlerFunc {
	return func(c *gin.Context) {
		studentID := c.GetString("student_id")

		rec, err := engine.GetRecommendations(studentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"progression":    rec.Progression,
			"next_milestone": rec.NextMilestone,
		})
	}
}

// DueCardsHandler retourne les cartes dues pour révision.
func DueCardsHandler(engine *recommendation.AdaptiveEngine) gin.HandlerFunc {
	return func(c *gin.Context) {
		studentID := c.GetString("student_id")

		rec, err := engine.GetRecommendations(studentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"due_cards": rec.DueCards,
			"total":     len(rec.DueCards),
		})
	}
}
