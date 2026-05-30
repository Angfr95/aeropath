package domain

// Concept représente un concept/cours théorique.
type Concept struct {
	ID             string   `json:"id"`              // ex: "PPL_AIRLAW_001_C01"
	ModuleID       string   `json:"module_id"`       // ex: "PPL_AIRLAW_001"
	License        License  `json:"license"`
	Category       Category `json:"category"`
	TitleFr        string   `json:"title_fr"`
	TitleEn        string   `json:"title_en"`
	ContentFr      string   `json:"content_fr"`
	ContentEn      string   `json:"content_en"`
	Difficulty     string   `json:"difficulty"`      // "easy", "medium", "hard"
	Level          int      `json:"level"`           // 1=Basic, 2=Intermediate, 3=Advanced
	RiskLevel      string   `json:"risk_level"`      // "low", "medium", "high"
	PhaseOfFlight  []string `json:"phase_of_flight"` // "preflight", "enroute", "postflight"
	Tags           []string `json:"tags"`
	FaaNoteFr      string   `json:"faa_note_fr,omitempty"`
	FaaNoteEn      string   `json:"faa_note_en,omitempty"`
	CommonErrors   []string `json:"common_errors,omitempty"`
	RelatedConcepts []string `json:"related_concepts,omitempty"`
}

// Flashcard représente une carte de révision rapide.
type Flashcard struct {
	ID        string   `json:"id"`
	ConceptID string   `json:"concept_id"`
	License   License  `json:"license"`
	Category  Category `json:"category"`
	FrontFr   string   `json:"front_fr"`
	FrontEn   string   `json:"front_en"`
	BackFr    string   `json:"back_fr"`
	BackEn    string   `json:"back_en"`
	Tags      []string `json:"tags,omitempty"`
}

// Scenario représente un scénario réaliste (cas pratique).
type Scenario struct {
	ID               string            `json:"id"`
	ModuleID         string            `json:"module_id"`
	License          License           `json:"license"`
	Category         Category          `json:"category"`
	TitleFr          string            `json:"title_fr"`
	TitleEn          string            `json:"title_en"`
	ContextFr        string            `json:"context_fr"`
	ContextEn        string            `json:"context_en"`
	SituationFr      string            `json:"situation_fr"`
	SituationEn      string            `json:"situation_en"`
	DecisionFr       string            `json:"decision_fr"`
	DecisionEn       string            `json:"decision_en"`
	RisksFr          []string          `json:"risks_fr"`
	RisksEn          []string          `json:"risks_en"`
	SolutionFr       string            `json:"solution_fr"`
	SolutionEn       string            `json:"solution_en"`
	ExplanationFr    string            `json:"explanation_fr"`
	ExplanationEn    string            `json:"explanation_en"`
}

// OperationalRule représente une règle opérationnelle.
type OperationalRule struct {
	ID         string   `json:"id"`
	ModuleID   string   `json:"module_id"`
	TitleFr    string   `json:"title_fr"`
	TitleEn    string   `json:"title_en"`
	ProcedureFr string  `json:"procedure_fr"`
	ProcedureEn string  `json:"procedure_en"`
	Conditions []string `json:"conditions"`
	CommonErrors []string `json:"common_errors,omitempty"`
	SafetyNotes []string `json:"safety_notes,omitempty"`
}

// Remediation représente une analyse d'erreur pour l'apprentissage adaptatif.
type Remediation struct {
	ID                 string            `json:"id"`
	ModuleID           string            `json:"module_id"`
	StudentLevel       string            `json:"student_level"`
	ChapterTitleFr     string            `json:"chapter_title_fr"`
	ChapterTitleEn     string            `json:"chapter_title_en"`
	ConceptID          string            `json:"concept_id"`
	QuestionID         string            `json:"question_id"`
	SelectedAnswer     string            `json:"selected_answer"`
	CorrectAnswer      string            `json:"correct_answer"`
	ErrorCategory      string            `json:"error_category"` // "misunderstanding", "terminology_confusion", "procedural_error"
	ErrorDescriptionFr string            `json:"error_description_fr"`
	ErrorDescriptionEn string            `json:"error_description_en"`
	MisconceptionFr    string            `json:"misconception_fr"`
	MisconceptionEn    string            `json:"misconception_en"`
	Severity           string            `json:"severity"`
	OperationalRisk    string            `json:"operational_risk"`
	RecommendedConcepts []string         `json:"recommended_concepts"`
	RecommendedFlashcards []string       `json:"recommended_flashcards"`
	RecommendedScenarios []string         `json:"recommended_scenarios"`
	InstructorFeedbackFr string          `json:"instructor_feedback_fr"`
	InstructorFeedbackEn string          `json:"instructor_feedback_en"`
	NextLearningStepFr  string           `json:"next_learning_step_fr"`
	NextLearningStepEn  string           `json:"next_learning_step_en"`
}
