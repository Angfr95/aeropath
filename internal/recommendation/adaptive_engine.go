package recommendation

import (
	"time"

	"aeropath/internal/domain"
)

// AdaptiveEngine est le moteur de recommandation adaptative.
//
// 📖 DDIA Chapitre 12 : "The Future of Data Systems"
//    C'est un exemple de "derived data system" : on combine plusieurs
//    sources de données (historique, questions, leçons) pour créer
//    une nouvelle donnée (les recommandations).
//
//    Le moteur utilise 4 algorithmes :
//    1. Spaced Repetition (SM-2) → quand revoir une question
//    2. Mastery Score → quel est le niveau sur chaque sujet
//    3. Weak Topics → où sont les difficultés
//    4. Progression → où en est l'étudiant dans son parcours
//
// 📖 DDIA Chapitre 1 : "Reliability"
//    Si le moteur de recommandation plante, l'API continue de fonctionner.
//    Les étudiants peuvent toujours répondre à des questions.
//    Simplement, ils n'auront pas de recommandations personnalisées.
//    C'est le principe de "graceful degradation".
type AdaptiveEngine struct {
	questionRepo domain.QuestionRepository
	historyRepo  domain.QuestionHistoryRepository
	lessonRepo   domain.LessonRepository
}

// New crée un nouveau moteur adaptatif.
func New(qr domain.QuestionRepository, hr domain.QuestionHistoryRepository, lr domain.LessonRepository) *AdaptiveEngine {
	return &AdaptiveEngine{
		questionRepo: qr,
		historyRepo:  hr,
		lessonRepo:   lr,
	}
}

// Recommendation contient les recommandations pour un étudiant.
type Recommendation struct {
	StudentID      string        `json:"student_id"`
	NextTheme      *WeakTopic    `json:"next_theme,omitempty"`
	DueCards       []*Card       `json:"due_cards,omitempty"`
	WeakTopics     []*WeakTopic  `json:"weak_topics,omitempty"`
	Progression    *Progression  `json:"progression,omitempty"`
	NextMilestone  string        `json:"next_milestone,omitempty"`
}

// GetRecommendations génère des recommandations personnalisées pour un étudiant.
//
// 📖 DDIA Chapitre 12 : "Integrating Derived Data"
//    Cette fonction est un "materialized view" en temps réel.
//    À chaque appel, elle recalcule les recommandations à partir
//    des données brutes (historique). C'est un compromis :
//    - Avantage : toujours à jour
//    - Inconvénient : plus lent qu'un cache pré-calculé
//
//    Pour améliorer les performances, on pourrait :
//    1. Mettre en cache les résultats pendant 5 minutes (Redis)
//    2. Pré-calculer les recommandations via un worker NATS
//    3. Utiliser une base de données dédiée (ClickHouse)
func (e *AdaptiveEngine) GetRecommendations(studentID string) (*Recommendation, error) {
	// Récupérer l'historique et les questions
	history, err := e.historyRepo.GetHistory(studentID)
	if err != nil {
		return nil, err
	}

	allQuestions, err := e.questionRepo.FindAll()
	if err != nil {
		return nil, err
	}

	// Récupérer les stats
	stats, err := e.historyRepo.GetStats(studentID)
	if err != nil {
		return nil, err
	}

	// Identifier les sujets faibles
	weakTopics := IdentifyWeakTopics([]*domain.StudentStats{stats}, history, allQuestions)

	// Calculer la progression par thème
	themeMastery := make(map[string]float64)
	questionsDone := make(map[string]int)
	questionsTotal := make(map[string]int)

	for _, wt := range weakTopics {
		themeMastery[wt.Theme] = wt.MasteryScore
		questionsDone[wt.Theme] = wt.TotalQuestions
		// Compter le total des questions pour ce thème
		count := 0
		for _, q := range allQuestions {
			if q.Theme == wt.Theme {
				count++
			}
		}
		questionsTotal[wt.Theme] = count
	}

	progression := CalculateProgression(studentID, themeMastery, questionsDone, questionsTotal)

	// Prochain palier
	avgMastery := 0.0
	if len(weakTopics) > 0 {
		for _, wt := range weakTopics {
			avgMastery += wt.MasteryScore
		}
		avgMastery /= float64(len(weakTopics))
	}
	nextMilestone := GetNextMilestone(avgMastery)

	// Prochain thème recommandé
	var recentThemes []string
	for _, h := range history {
		recentThemes = append(recentThemes, h.QuestionID)
	}
	nextTheme := NextRecommendedTheme(weakTopics, recentThemes)

	// Cartes dues pour la répétition espacée
	var dueCards []*Card
	// Construire les cartes à partir de l'historique
	cardMap := make(map[string]*Card)
	for _, h := range history {
		if _, exists := cardMap[h.QuestionID]; !exists {
			cardMap[h.QuestionID] = NewCard(h.QuestionID)
		}
		quality := 5
		if !h.WasCorrect {
			quality = 1
		}
		cardMap[h.QuestionID].Review(quality)
	}
	for _, card := range cardMap {
		dueCards = append(dueCards, card)
	}
	dueCards = DueCards(dueCards, time.Now())
	SortByPriority(dueCards, time.Now())

	milestoneName := ""
	if nextMilestone != nil {
		milestoneName = nextMilestone.Name
	}
	return &Recommendation{
		StudentID:     studentID,
		NextTheme:     nextTheme,
		DueCards:      dueCards,
		WeakTopics:    weakTopics,
		Progression:   progression,
		NextMilestone: milestoneName,
	}, nil
}
