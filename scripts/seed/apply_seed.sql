-- ============================================================================
-- SEED COMPLET : Nettoie + Initialise la table app_config + seed
-- ============================================================================

-- Nettoyage des données existantes
DELETE FROM student_question_history;
DELETE FROM questions;
DELETE FROM lessons;
DELETE FROM user_gamification;
DELETE FROM user_progress;
DELETE FROM achievements;
DELETE FROM leaderboard_weekly;
DELETE FROM user_quests;
DELETE FROM spaced_repetition;
DELETE FROM app_config;

-- ============================================================================
-- 1. APP CONFIG (JSONB dans aero_config)
-- ============================================================================
INSERT INTO app_config (config_key, config_value) VALUES
('adaptive_learning', '{
    "level_1": {"pass_threshold": 0.80, "hint_enabled": true, "time_limit_sec": null, "immediate_feedback": true},
    "level_2": {"pass_threshold": 0.75, "hint_enabled": false, "time_limit_sec": null, "immediate_feedback": true},
    "level_3": {"pass_threshold": 0.75, "hint_enabled": false, "time_limit_sec": 45, "immediate_feedback": false}
}');

INSERT INTO app_config (config_key, config_value) VALUES
('gamification_rules', '{
    "hearts_enabled": true, "hearts_total": 5,
    "heart_penalty_level2": 1, "heart_penalty_level3": 2,
    "xp_per_lesson": 10, "xp_per_qcm_level1": 20, "xp_per_qcm_level2": 30, "xp_per_qcm_level3": 50,
    "streak_enabled": true, "leaderboard_weekly": true
}');

INSERT INTO app_config (config_key, config_value) VALUES
('exam_config', '{
    "ppl.air_law": {"questions_count": 16, "time_minutes": 30, "passing_score_percent": 75, "randomized": true, "no_backward": true},
    "ppl.meteorology": {"questions_count": 16, "time_minutes": 30, "passing_score_percent": 75, "randomized": true, "no_backward": true}
}');


