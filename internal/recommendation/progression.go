package recommendation

// Progression représente la progression globale d'un étudiant.
//
// 📖 DDIA Chapitre 12 : "The Future of Data Systems"
//    La progression est une "vue intégrée" (integrated view)
//    qui combine plusieurs sources :
//    - Les scores de maîtrise par thème (MasteryScore)
//    - Le nombre de questions faites (QuestionHistory)
//    - Le nombre total de questions disponibles (Question)
//
//    C'est exactement ce que fait une "data warehouse" :
//    on prend des données de différentes sources, on les
//    transforme, et on crée une vue unifiée pour l'utilisateur.
type Progression struct {
	StudentID        string                  `json:"student_id"`
	TotalLessons     int                     `json:"total_lessons"`
	CompletedLessons int                     `json:"completed_lessons"`
	ProgressPercent  float64                 `json:"progress_percent"`
	ThemeProgress    map[string]ThemeProgress `json:"theme_progress"`
	CurrentLevel     int                     `json:"current_level"` // 1-5
	NextMilestone    string                  `json:"next_milestone"`
}

// ThemeProgress représente la progression sur un thème spécifique.
type ThemeProgress struct {
	Theme           string  `json:"theme"`
	MasteryScore    float64 `json:"mastery_score"`
	QuestionsDone   int     `json:"questions_done"`
	QuestionsTotal  int     `json:"questions_total"`
	PercentComplete float64 `json:"percent_complete"`
}

// CalculateProgression calcule la progression globale d'un étudiant.
//
// 📖 DDIA Chapitre 10 : "Batch Processing"
//    Cette fonction est un "batch processing" miniature :
//    elle prend un lot de données en entrée (les scores par thème)
//    et produit un résultat unique (la progression globale).
//    C'est le même pattern que MapReduce, mais en mémoire.
func CalculateProgression(
	studentID string,
	themeMastery map[string]float64,
	questionsDone map[string]int,
	questionsTotal map[string]int,
) *Progression {
	p := &Progression{
		StudentID:     studentID,
		ThemeProgress: make(map[string]ThemeProgress),
	}

	totalMastery := 0.0
	themeCount := 0

	for theme, mastery := range themeMastery {
		done := questionsDone[theme]
		total := questionsTotal[theme]

		percentComplete := 0.0
		if total > 0 {
			percentComplete = float64(done) / float64(total) * 100
		}

		p.ThemeProgress[theme] = ThemeProgress{
			Theme:           theme,
			MasteryScore:    mastery,
			QuestionsDone:   done,
			QuestionsTotal:  total,
			PercentComplete: percentComplete,
		}

		totalMastery += mastery
		themeCount++
	}

	if themeCount > 0 {
		p.CurrentLevel = RecommendedDifficulty(totalMastery / float64(themeCount))
	}

	return p
}

// Milestone représente un palier de progression.
//
// 📖 DDIA Chapitre 1 : "Maintainability"
//    Les paliers sont définis dans le code, pas dans la base.
//    Pourquoi ? Parce qu'ils changent rarement et qu'ils sont
//    faciles à modifier ici. Si on les mettait dans la base,
//    il faudrait une migration à chaque changement.
//
//    C'est le principe de "configuration over database" :
//    les données qui changent avec le code restent dans le code.
type Milestone struct {
	Name        string  `json:"name"`
	Threshold   float64 `json:"threshold"`
	Description string  `json:"description"`
}

// GetNextMilestone retourne le prochain palier à atteindre.
func GetNextMilestone(currentMastery float64) *Milestone {
	milestones := []Milestone{
		{Name: "Débutant", Threshold: 0.1, Description: "Commencer à étudier"},
		{Name: "Apprenti", Threshold: 0.25, Description: "Maîtriser les bases"},
		{Name: "Intermédiaire", Threshold: 0.5, Description: "Consolider les connaissances"},
		{Name: "Avancé", Threshold: 0.75, Description: "Approfondir les sujets complexes"},
		{Name: "Expert", Threshold: 0.9, Description: "Atteindre l'excellence"},
		{Name: "Maître", Threshold: 1.0, Description: "Maîtrise complète"},
	}

	for _, m := range milestones {
		if currentMastery < m.Threshold {
			return &m
		}
	}

	return &milestones[len(milestones)-1]
}
