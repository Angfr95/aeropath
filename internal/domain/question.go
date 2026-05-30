package domain

import "time"

// Question représente une question d'examen aéronautique enrichie.
//
// 📖 DDIA Chapitre 3 : "Storage and Retrieval"
//    Les questions sont stockées dans PostgreSQL (base relationnelle).
//    On utilise un index sur (license, category) pour les requêtes fréquentes
//    comme "toutes les questions PPL de météo".
type Question struct {
	ID            string    `json:"id"`
	LessonID      string    `json:"lesson_id,omitempty"`
	License       License   `json:"license"`
	Category      Category  `json:"category"`
	Theme         string    `json:"theme"`
	Subtopic      string    `json:"subtopic,omitempty"`
	Difficulty    int       `json:"difficulty"`
	Level         int       `json:"level"`          // 1=Basic, 2=Intermediate, 3=Advanced (Duolingo-style)
	QuestionFr    string    `json:"question_fr"`
	QuestionEn    string    `json:"question_en"`
	Options       []string  `json:"options"`
	AnswerKey     string    `json:"answer_key"`
	ExplanationFr string    `json:"explanation_fr,omitempty"`
	ExplanationEn string    `json:"explanation_en,omitempty"`
	FaaNoteFr     string    `json:"faa_note_fr,omitempty"`     // Différence FAA vs EASA (FR)
	FaaNoteEn     string    `json:"faa_note_en,omitempty"`     // Différence FAA vs EASA (EN)
	Tags          []string  `json:"tags,omitempty"`            // Mots-clés pour classification
	DifficultyScore float64 `json:"difficulty_score,omitempty"` // 0.0-1.0 pour algo adaptatif
	Distractors   map[string]string `json:"distractors_rationale,omitempty"` // Pourquoi chaque distracteur est faux
	CreatedAt     time.Time `json:"created_at"`
}


// QuestionRepository définit le contrat pour accéder aux questions.
//
// 📖 DDIA Chapitre 2 : "Data Models and Query Languages"
//    Le Repository Pattern cache la complexité du stockage.
//    Les handlers HTTP ne savent pas si les questions viennent
//    de PostgreSQL, Redis, ou d'un fichier JSON.
//    C'est le principe de "séparation des préoccupations".
type QuestionRepository interface {
	Create(q *Question) error
	Update(q *Question) error
	FindByID(id string) (*Question, error)
	FindAll() ([]*Question, error)
	FindAllPaginated(limit, offset int) ([]*Question, error)
	FindByTheme(theme string) ([]*Question, error)
	FindByLicense(license License) ([]*Question, error)
	FindByCategory(category Category) ([]*Question, error)
	FindByLicenseAndCategory(license License, category Category) ([]*Question, error)
	FindByDifficulty(difficulty int) ([]*Question, error)
	FindBySubtopic(subtopic string) ([]*Question, error)
	Search(query string) ([]*Question, error)
	FindRandom(limit int, license *License, category *Category, theme *string) ([]*Question, error)
	Count() (int, error)
	CountByLicense(license License) (int, error)
	CountByCategory(category Category) (int, error)
	CountByTheme(theme string) (int, error)
	Delete(id string) error
}
