package domain

// AdaptiveConfig contient les paramètres de l'algorithme adaptatif.
type AdaptiveConfig struct {
	Level1 struct {
		PassThreshold float64 `yaml:"pass_threshold" json:"pass_threshold"` // 0.80 = 80%
		HintEnabled   bool    `yaml:"hint_enabled"   json:"hint_enabled"`   // true
		TimeLimitSec  *int    `yaml:"time_limit_sec" json:"time_limit_sec"` // null = pas de limite
		ImmediateFeedback bool `yaml:"immediate_feedback" json:"immediate_feedback"`
	} `yaml:"level_1" json:"level_1"`
	Level2 struct {
		PassThreshold float64 `yaml:"pass_threshold" json:"pass_threshold"` // 0.75
		HintEnabled   bool    `yaml:"hint_enabled"   json:"hint_enabled"`
		TimeLimitSec  *int    `yaml:"time_limit_sec" json:"time_limit_sec"`
		ImmediateFeedback bool `yaml:"immediate_feedback" json:"immediate_feedback"`
	} `yaml:"level_2" json:"level_2"`
	Level3 struct {
		PassThreshold float64 `yaml:"pass_threshold" json:"pass_threshold"` // 0.75
		HintEnabled   bool    `yaml:"hint_enabled"   json:"hint_enabled"`
		TimeLimitSec  int     `yaml:"time_limit_sec" json:"time_limit_sec"` // 45s
		ImmediateFeedback bool `yaml:"immediate_feedback" json:"immediate_feedback"`
	} `yaml:"level_3" json:"level_3"`
}

// GamificationConfig contient les paramètres de gamification.
type GamificationConfig struct {
	HeartsEnabled    bool `yaml:"hearts_enabled"    json:"hearts_enabled"`
	HeartsTotal      int  `yaml:"hearts_total"      json:"hearts_total"`       // 5
	HeartPenaltyL2   int  `yaml:"heart_penalty_level2" json:"heart_penalty_level2"` // 1
	HeartPenaltyL3   int  `yaml:"heart_penalty_level3" json:"heart_penalty_level3"` // 2
	XPPerLesson      int  `yaml:"xp_per_lesson"     json:"xp_per_lesson"`      // 10
	XPPerQCMLevel1   int  `yaml:"xp_per_qcm_level1" json:"xp_per_qcm_level1"` // 20
	XPPerQCMLevel2   int  `yaml:"xp_per_qcm_level2" json:"xp_per_qcm_level2"` // 30
	XPPerQCMLevel3   int  `yaml:"xp_per_qcm_level3" json:"xp_per_qcm_level3"` // 50
	StreakEnabled    bool `yaml:"streak_enabled"     json:"streak_enabled"`    // true
	LeaderboardWeekly bool `yaml:"leaderboard_weekly" json:"leaderboard_weekly"` // true
}

// ExamConfigItem contient la configuration d'un examen.
type ExamConfigItem struct {
	QuestionsCount int     `yaml:"questions_count" json:"questions_count"`
	TimeMinutes    int     `yaml:"time_minutes"    json:"time_minutes"`
	PassingScore   float64 `yaml:"passing_score_percent" json:"passing_score_percent"`
	Randomized     bool    `yaml:"randomized"      json:"randomized"`
	NoBackward     bool    `yaml:"no_backward"     json:"no_backward"`
}

// ExamConfig regroupe la config des examens par matière.
type ExamConfig struct {
	PPLAirLaw        ExamConfigItem `yaml:"ppl.air_law"        json:"ppl_air_law"`
	PPLMeteorology   ExamConfigItem `yaml:"ppl.meteorology"    json:"ppl_meteorology"`
	PPLNavigation    ExamConfigItem `yaml:"ppl.navigation"     json:"ppl_navigation"`
	PPLPerformance   ExamConfigItem `yaml:"ppl.performance"    json:"ppl_performance"`
	ATPLAirLaw       ExamConfigItem `yaml:"atpl.air_law"       json:"atpl_air_law"`
	ATPLMeteorology  ExamConfigItem `yaml:"atpl.meteorology"   json:"atpl_meteorology"`
	ATPLNavigation   ExamConfigItem `yaml:"atpl.navigation"    json:"atpl_navigation"`
	ATPLPerformance  ExamConfigItem `yaml:"atpl.performance"   json:"atpl_performance"`
}

// AppConfig est la configuration complète de l'application.
type AppConfig struct {
	Adaptive    AdaptiveConfig    `yaml:"adaptive_learning" json:"adaptive_learning"`
	Gamification GamificationConfig `yaml:"gamification_rules" json:"gamification_rules"`
	Exam        ExamConfig        `yaml:"exam_config"       json:"exam_config"`
}

// DefaultExamConfig retourne la config par défaut pour PPL Air Law.
func DefaultExamConfig() ExamConfigItem {
	return ExamConfigItem{
		QuestionsCount: 16,
		TimeMinutes:    30,
		PassingScore:   75.0,
		Randomized:     true,
		NoBackward:     true,
	}
}

// DefaultAdaptiveConfig retourne la config adaptative par défaut.
func DefaultAdaptiveConfig() AdaptiveConfig {
	cfg := AdaptiveConfig{}
	cfg.Level1.PassThreshold = 0.80
	cfg.Level1.HintEnabled = true
	cfg.Level1.ImmediateFeedback = true

	cfg.Level2.PassThreshold = 0.75
	cfg.Level2.HintEnabled = false
	cfg.Level2.ImmediateFeedback = true

	cfg.Level3.PassThreshold = 0.75
	cfg.Level3.HintEnabled = false
	cfg.Level3.TimeLimitSec = 45
	cfg.Level3.ImmediateFeedback = false

	return cfg
}

// DefaultGamificationConfig retourne la config gamification par défaut.
func DefaultGamificationConfig() GamificationConfig {
	return GamificationConfig{
		HeartsEnabled:    true,
		HeartsTotal:      5,
		HeartPenaltyL2:   1,
		HeartPenaltyL3:   2,
		XPPerLesson:      10,
		XPPerQCMLevel1:   20,
		XPPerQCMLevel2:   30,
		XPPerQCMLevel3:   50,
		StreakEnabled:    true,
		LeaderboardWeekly: true,
	}
}
