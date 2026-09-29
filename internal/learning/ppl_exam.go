package learning

import (
	"fmt"
	"math/rand"
	"slices"

	"aeropath/internal/domain"
)

// French PPL(A) format, DSAC notice of 3 October 2023, effective 4 December 2023.
// These are practice papers, not official questions or a licence recommendation.
type PPLPaper struct {
	ID      string `json:"id"`
	TitleFr string `json:"title_fr"`
	TitleEn string `json:"title_en"`
	Count   int    `json:"count"`
	Minutes int    `json:"minutes"`
}

var PPLPapers = []PPLPaper{
	{"airlaw", "Droit aérien", "Air law", 20, 30},
	{"human_performance", "Performance humaine", "Human performance", 12, 18},
	{"meteorology", "Météorologie", "Meteorology", 20, 30},
	{"communications", "Communications", "Communications", 8, 12},
	{"navigation", "Navigation", "Navigation", 20, 50},
	{"principles_of_flight", "Principes du vol", "Principles of flight", 12, 18},
	{"operational_procedures", "Procédures opérationnelles", "Operational procedures", 12, 18},
	{"flight_performance_planning", "Performances et préparation du vol", "Flight performance and planning", 12, 18},
	{"aircraft_general", "Connaissance générale de l’aéronef", "Aircraft general knowledge", 12, 18},
}

func PPLPaperByID(id string) (PPLPaper, bool) {
	for _, p := range PPLPapers {
		if p.ID == id {
			return p, true
		}
	}
	return PPLPaper{}, false
}

// Round-robin across shuffled lessons prevents one large lesson dominating a paper.
// Eligibility is explicit: four options alone do not qualify imported content.
func (e *Engine) PPLExam(id string) (PPLPaper, []*domain.Question, error) {
	paper, ok := PPLPaperByID(id)
	if !ok {
		return paper, nil, fmt.Errorf("matière PPL inconnue")
	}
	all, err := e.questionRepo.FindAll()
	if err != nil {
		return paper, nil, err
	}
	groups := map[string][]*domain.Question{}
	seen := map[string]bool{}
	for _, q := range all {
		if q == nil || seen[q.ID] || q.License != "PPL" || len(q.Options) != 4 ||
			!slices.Contains(q.Tags, "exam-mcq") || !slices.Contains(q.Tags, "exam-subject:"+id) {
			continue
		}
		seen[q.ID] = true
		groups[q.LessonID] = append(groups[q.LessonID], q)
	}
	keys := make([]string, 0, len(groups))
	for key, questions := range groups {
		keys = append(keys, key)
		rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })
	}
	rand.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })
	result := make([]*domain.Question, 0, paper.Count)
	for round := 0; len(result) < paper.Count; round++ {
		added := false
		for _, key := range keys {
			if round < len(groups[key]) {
				result = append(result, groups[key][round])
				added = true
				if len(result) == paper.Count {
					break
				}
			}
		}
		if !added {
			return paper, nil, fmt.Errorf("banque insuffisante pour cette épreuve")
		}
	}
	return paper, result, nil
}
