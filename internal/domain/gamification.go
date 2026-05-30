package domain

import "time"

// UserProgress suit la progression d'un étudiant par matière.
type UserProgress struct {
	UserID            string    `json:"user_id"`
	License           License   `json:"license"`
	Language          string    `json:"language"`
	SubjectID         string    `json:"subject_id"` // ex: "010_airlaw"
	Level1Completed   bool      `json:"level_1_completed"`
	Level2Completed   bool      `json:"level_2_completed"`
	Level3Completed   bool      `json:"level_3_completed"`
	ScoreAvgLevel1    float64   `json:"score_avg_level1"`
	ScoreAvgLevel2    float64   `json:"score_avg_level2"`
	ScoreAvgLevel3    float64   `json:"score_avg_level3"`
	QuestionsAttempted int      `json:"questions_attempted"`
	LastSeenDate      time.Time `json:"last_seen_date"`
}

// UserGamification rassemble les données de gamification d'un étudiant.
type UserGamification struct {
	UserID         string    `json:"user_id"`
	Hearts         int       `json:"hearts"`          // max 5
	XP             int       `json:"xp"`
	Streak         int       `json:"streak"`          // jours consécutifs
	Level          int       `json:"level"`           // user level (1-N)
	PreferredLang  string    `json:"preferred_language"`
	PreferredLicense License `json:"preferred_license"`
	LastActiveAt   time.Time `json:"last_active_at"`
	CreatedAt      time.Time `json:"created_at"`
}

// Achievement représente un badge débloqué.
type Achievement struct {
	UserID     string    `json:"user_id"`
	BadgeID    string    `json:"badge_id"` // ex: "airlaw_master", "7day_streak"
	Level      int       `json:"level"`    // 1/2/3
	UnlockedAt time.Time `json:"unlocked_at"`
}

// LeaderboardEntry représente une entrée au classement hebdomadaire.
type LeaderboardEntry struct {
	UserID      string `json:"user_id"`
	WeekStart   string `json:"week_start_date"` // YYYY-MM-DD
	XPEarned    int    `json:"xp_earned_week"`
	Rank        int    `json:"rank"`
	Tier        string `json:"tier"` // 'bronze', 'silver', 'gold', 'diamond'
}

// UserQuest représente une quête active.
type UserQuest struct {
	UserID    string     `json:"user_id"`
	QuestID   string     `json:"quest_id"`   // ex: "3_lessons_day", "100q_no_error"
	Progress  int        `json:"progress"`   // ex: 0/3 ou 0/100
	Max       int        `json:"max"`
	RewardXP  int        `json:"reward_xp"`
	Completed bool       `json:"completed"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

// SpacedRepetitionItem gère la répétition espacée d'une question.
type SpacedRepetitionItem struct {
	UserID       string    `json:"user_id"`
	QuestionID   string    `json:"question_id"`
	Level        int       `json:"level"`         // niveau actuel (1=jour1, 2=jour3, 3=jour7, 4=jour14, 5=jour30)
	NextReviewAt time.Time `json:"next_review_at"`
	LastSeenAt   time.Time `json:"last_seen_at"`
	TimesCorrect int       `json:"times_correct"`
	TimesWrong   int       `json:"times_wrong"`
}

// GamificationRepository définit le contrat pour les données de gamification.
type GamificationRepository interface {
	// Progression
	GetProgress(userID, subjectID string) (*UserProgress, error)
	UpsertProgress(p *UserProgress) error

	// Gamification
	GetGamification(userID string) (*UserGamification, error)
	UpsertGamification(g *UserGamification) error
	UpdateHearts(userID string, hearts int) error
	AddXP(userID string, xp int) error
	UpdateStreak(userID string, streak int) error

	// Achievements
	GetAchievements(userID string) ([]*Achievement, error)
	UnlockAchievement(a *Achievement) error

	// Leaderboard
	GetLeaderboard(weekStart string, limit int) ([]*LeaderboardEntry, error)
	UpsertLeaderboardEntry(e *LeaderboardEntry) error

	// Quests
	GetQuests(userID string) ([]*UserQuest, error)
	UpdateQuestProgress(userID, questID string, progress int) error
	CompleteQuest(userID, questID string) error

	// Spaced Repetition
	GetItemsToReview(userID string, now time.Time) ([]*SpacedRepetitionItem, error)
	UpsertSpacedRepetition(item *SpacedRepetitionItem) error
	RecordAnswer(userID, questionID string, wasCorrect bool) error
}
