package recommendation

import (
	"sort"

	"aeropath/internal/domain"
)

// WeakTopic représente un sujet où l'étudiant a des difficultés.
//
// 📖 DDIA Chapitre 12 : "The Future of Data Systems"
//    Les sujets faibles sont identifiés en comparant les scores
//    de maîtrise entre différents thèmes. C'est un exemple de
//    "data integration" : on combine l'historique (QuestionHistory)
//    avec les métadonnées (Question.Theme) pour créer une nouvelle
//    vue des données.
type WeakTopic struct {
	Theme        string  `json:"theme"`
	Category     string  `json:"category,omitempty"`
	MasteryScore float64 `json:"mastery_score"`
	TotalQuestions int   `json:"total_questions"`
	Priority     int     `json:"priority"` // 1 = priorité maximale
}

// IdentifyWeakTopics identifie les sujets faibles d'un étudiant.
// Retourne une liste triée par priorité (les plus faibles en premier).
//
// 📖 DDIA Chapitre 2 : "Data Models and Query Languages"
//    Cette fonction fait du "map-reduce" à la main :
//    1. Map : on groupe les réponses par thème
//    2. Reduce : on calcule le score de maîtrise pour chaque thème
//    3. Sort : on trie par score croissant
//    C'est exactement le même pattern que MapReduce (Google, 2004).
func IdentifyWeakTopics(stats []*domain.StudentStats, history []*domain.QuestionHistory, questions []*domain.Question) []*WeakTopic {
	// Grouper les questions par thème
	themeStats := make(map[string]*struct {
		correct int
		wrong   int
		total   int
	})

	// Compter les réponses par thème
	questionTheme := make(map[string]string) // questionID -> theme
	for _, q := range questions {
		questionTheme[q.ID] = q.Theme
	}

	for _, h := range history {
		theme, ok := questionTheme[h.QuestionID]
		if !ok {
			continue
		}

		ts, exists := themeStats[theme]
		if !exists {
			ts = &struct {
				correct int
				wrong   int
				total   int
			}{}
			themeStats[theme] = ts
		}

		ts.total++
		if h.WasCorrect {
			ts.correct++
		} else {
			ts.wrong++
		}
	}

	// Calculer les scores de maîtrise par thème
	var weakTopics []*WeakTopic
	for theme, ts := range themeStats {
		score := CalculateMastery(ts.correct, ts.wrong)
		weakTopics = append(weakTopics, &WeakTopic{
			Theme:         theme,
			MasteryScore:  score,
			TotalQuestions: ts.total,
		})
	}

	// Trier par score de maîtrise croissant (les plus faibles en premier)
	sort.Slice(weakTopics, func(i, j int) bool {
		return weakTopics[i].MasteryScore < weakTopics[j].MasteryScore
	})

	// Assigner les priorités
	for i := range weakTopics {
		weakTopics[i].Priority = i + 1
	}

	return weakTopics
}

// NextRecommendedTheme retourne le thème recommandé pour l'étudiant.
// Priorise les sujets faibles non vus récemment.
//
// 📖 DDIA Chapitre 1 : "Maintainability"
//    Cette fonction implémente une règle métier simple :
//    "travaille d'abord tes points faibles, mais alterne
//    pour ne pas te lasser". C'est facile à comprendre et
//    à modifier si le besoin change.
func NextRecommendedTheme(weakTopics []*WeakTopic, recentThemes []string) *WeakTopic {
	if len(weakTopics) == 0 {
		return nil
	}

	// Chercher le sujet le plus faible qui n'a pas été vu récemment
	recent := make(map[string]bool)
	for _, t := range recentThemes {
		recent[t] = true
	}

	for _, wt := range weakTopics {
		if !recent[wt.Theme] {
			return wt
		}
	}

	// Si tous les sujets faibles ont été vus récemment, prendre le plus faible
	return weakTopics[0]
}
