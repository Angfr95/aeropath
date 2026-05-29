package recommendation

// MasteryScore représente le niveau de maîtrise d'un étudiant sur un sujet.
//
// 📖 DDIA Chapitre 12 : "The Future of Data Systems"
//    Le score de maîtrise est une "donnée dérivée" (derived data).
//    On part de données brutes (l'historique des réponses) et on
//    les transforme en une métrique utile (le score).
//
//    C'est le même principe qu'une "materialized view" en SQL :
//    au lieu de recalculer à chaque fois, on pourrait stocker
//    le score dans une table dédiée et le mettre à jour
//    à chaque nouvelle réponse (via un trigger ou un worker NATS).
type MasteryScore struct {
	StudentID     string  `json:"student_id"`
	Theme         string  `json:"theme"`
	Category      string  `json:"category,omitempty"`
	TotalAttempts int     `json:"total_attempts"`
	CorrectCount  int     `json:"correct_count"`
	WrongCount    int     `json:"wrong_count"`
	Score         float64 `json:"score"`         // 0.0 - 1.0
	Confidence    float64 `json:"confidence"`    // 0.0 - 1.0 (basé sur le nombre de tentatives)
}

// CalculateMastery calcule le score de maîtrise (0.0 - 1.0).
// Utilise une moyenne pondérée qui favorise les résultats récents.
//
// 📖 DDIA Chapitre 1 : "Reliability"
//    La formule est déterministe : pour les mêmes entrées,
//    on obtient toujours le même résultat. C'est important
//    pour la reproductibilité et le débogage.
func CalculateMastery(correct, wrong int) float64 {
	total := correct + wrong
	if total == 0 {
		return 0.0
	}

	// Pénalité pour les erreurs : chaque erreur compte double
	weightedCorrect := float64(correct)
	weightedWrong := float64(wrong) * 2.0
	weightedTotal := weightedCorrect + weightedWrong

	if weightedTotal == 0 {
		return 0.0
	}

	return weightedCorrect / weightedTotal
}

// CalculateConfidence calcule le niveau de confiance dans le score.
// Plus l'étudiant a répondu à des questions, plus la confiance est élevée.
//
// 📖 DDIA Chapitre 12 : "Trust in Derived Data"
//    La confiance est importante : un score basé sur 2 réponses
//    n'a pas la même valeur qu'un score basé sur 100 réponses.
//    C'est le même principe que les "error bars" en science :
//    plus on a de données, plus on est confiant.
func CalculateConfidence(totalAttempts int) float64 {
	// La confiance augmente avec le nombre de tentatives, asymptotiquement vers 1.0
	// 10 tentatives → ~0.5, 50 tentatives → ~0.83, 100+ → ~0.91
	return 1.0 - (1.0 / (1.0 + float64(totalAttempts)/10.0))
}

// NeedsReview retourne true si le score de maîtrise est inférieur au seuil.
func NeedsReview(score float64, threshold float64) bool {
	return score < threshold
}

// RecommendedDifficulty calcule la difficulté recommandée pour les prochaines questions.
// Si la maîtrise est bonne, on augmente la difficulté.
//
// 📖 DDIA Chapitre 1 : "Scalability"
//    Cette fonction adapte la charge de travail à l'étudiant.
//    C'est un exemple de "backpressure" : si l'étudiant réussit,
//    on augmente la difficulté (plus de charge). S'il échoue,
//    on diminue (moins de charge).
func RecommendedDifficulty(masteryScore float64) int {
	switch {
	case masteryScore >= 0.9:
		return 5
	case masteryScore >= 0.75:
		return 4
	case masteryScore >= 0.6:
		return 3
	case masteryScore >= 0.4:
		return 2
	default:
		return 1
	}
}
