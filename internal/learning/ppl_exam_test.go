package learning

import (
	"aeropath/internal/domain"
	"fmt"
	"testing"
)

type pplQuestions struct {
	domain.QuestionRepository
	questions []*domain.Question
}

func (r pplQuestions) FindAll() ([]*domain.Question, error) { return r.questions, nil }

func TestPPLPapersUseFrenchQuotasAndExcludeRevision(t *testing.T) {
	want := map[string][2]int{"airlaw": {20, 30}, "human_performance": {12, 18}, "meteorology": {20, 30}, "communications": {8, 12}, "navigation": {20, 50}, "principles_of_flight": {12, 18}, "operational_procedures": {12, 18}, "flight_performance_planning": {12, 18}, "aircraft_general": {12, 18}}
	var bank []*domain.Question
	for subject := range want {
		for i := 0; i < 30; i++ {
			bank = append(bank, &domain.Question{ID: fmt.Sprintf("%s-%d", subject, i), LessonID: fmt.Sprintf("lesson%d", i%3), License: "PPL", Options: []string{"one", "two", "three", "four"}, Tags: []string{"exam-mcq", "exam-subject:" + subject}})
		}
	}
	bank = append(bank, &domain.Question{ID: "tf", License: "PPL", Options: []string{"yes", "no"}, Tags: []string{"exam-mcq", "exam-subject:navigation"}},
		&domain.Question{ID: "untagged", License: "PPL", Options: []string{"a", "b", "c", "d"}},
		&domain.Question{ID: "atpl", License: "ATPL", Options: []string{"a", "b", "c", "d"}, Tags: []string{"exam-mcq", "exam-subject:navigation"}})
	engine := NewEngine(pplQuestions{questions: bank}, nil)
	total, minutes := 0, 0
	for _, p := range PPLPapers {
		if want[p.ID] != [2]int{p.Count, p.Minutes} {
			t.Fatalf("wrong policy: %+v", p)
		}
		_, qs, err := engine.PPLExam(p.ID)
		if err != nil {
			t.Fatal(err)
		}
		if len(qs) != p.Count {
			t.Fatal("incomplete paper")
		}
		seen := map[string]bool{}
		lessons := map[string]int{}
		for _, q := range qs {
			if seen[q.ID] || q.ID == "tf" || q.ID == "untagged" || q.ID == "atpl" {
				t.Fatalf("invalid selection %s", q.ID)
			}
			seen[q.ID] = true
			lessons[q.LessonID]++
		}
		for _, count := range lessons {
			if count > (p.Count+2)/3 {
				t.Fatal("paper dominated by one lesson")
			}
		}
		total += p.Count
		minutes += p.Minutes
	}
	if total != 128 || minutes != 212 {
		t.Fatalf("format %d questions, %d minutes", total, minutes)
	}
}

func TestPPLPaperFailsClosedForMissingBankAndUnknownSubject(t *testing.T) {
	e := NewEngine(pplQuestions{}, nil)
	for _, id := range []string{"navigation", "unknown"} {
		if _, qs, err := e.PPLExam(id); err == nil || len(qs) != 0 {
			t.Fatal("invalid paper accepted")
		}
	}
}
