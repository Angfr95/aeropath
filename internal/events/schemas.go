package events

import "time"

// EventType représente les types d'événements du système
type EventType string

const (
	// Événements d'apprentissage
	EventQuestionAnswered EventType = "learning.question.answered"
	EventExamStarted      EventType = "learning.exam.started"
	EventExamCompleted    EventType = "learning.exam.completed"
	EventLessonViewed     EventType = "learning.lesson.viewed"

	// Événements utilisateur
	EventStudentRegistered EventType = "student.registered"
	EventStudentLoggedIn   EventType = "student.logged_in"
	EventStudentUpdated    EventType = "student.updated"

	// Événements de recommandation
	EventRecommendationsUpdated EventType = "recommendation.updated"
	EventWeakTopicsDetected     EventType = "recommendation.weak_topics"
	EventMilestoneReached       EventType = "recommendation.milestone"

	// Événements système
	EventServiceHealthCheck EventType = "system.healthcheck"
	EventServiceDegraded    EventType = "system.service.degraded"
	EventServiceDown        EventType = "system.service.down"
)

// Event est la structure générique d'un événement
type Event struct {
	ID        string            `json:"id"`
	Type      EventType         `json:"type"`
	Source    string            `json:"source"`
	Timestamp time.Time         `json:"timestamp"`
	StudentID string            `json:"student_id,omitempty"`
	SessionID string            `json:"session_id,omitempty"`
	Data      map[string]any    `json:"data"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

// QuestionAnsweredEvent est l'événement émis quand un étudiant répond à une question
type QuestionAnsweredEvent struct {
	QuestionID string `json:"question_id"`
	Answer     string `json:"answer"`
	Correct    bool   `json:"correct"`
	License    string `json:"license"`
	Category   string `json:"category"`
	Theme      string `json:"theme"`
	Difficulty int    `json:"difficulty"`
	LatencyMs  int64  `json:"latency_ms"`
}

// ExamCompletedEvent est l'événement émis quand un examen est terminé
type ExamCompletedEvent struct {
	ExamID      string  `json:"exam_id"`
	License     string  `json:"license"`
	Category    string  `json:"category"`
	Total       int     `json:"total"`
	Correct     int     `json:"correct"`
	Score       float64 `json:"score"`
	Passed      bool    `json:"passed"`
	DurationSec int     `json:"duration_sec"`
}
