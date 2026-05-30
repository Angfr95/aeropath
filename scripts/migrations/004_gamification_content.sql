-- ============================================================================
-- MIGRATION 004 : Gamification + Contenu enrichi + Config
-- Architecture : licence + langue séparées en BDD distinctes
-- ============================================================================

-- ============================================================================
-- 1. TABLE USER_GAMIFICATION (dans aero_gamification)
-- ============================================================================
CREATE TABLE IF NOT EXISTS user_gamification (
    user_id            UUID PRIMARY KEY,
    hearts             INT NOT NULL DEFAULT 5 CHECK (hearts BETWEEN 0 AND 5),
    xp                 INT NOT NULL DEFAULT 0,
    streak             INT NOT NULL DEFAULT 0,
    level              INT NOT NULL DEFAULT 1,
    preferred_language TEXT NOT NULL DEFAULT 'fr',
    preferred_license  TEXT NOT NULL DEFAULT 'PPL',
    last_active_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    created_at         TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- ============================================================================
-- 2. TABLE USER_PROGRESS (dans aero_gamification)
-- ============================================================================
CREATE TABLE IF NOT EXISTS user_progress (
    user_id             UUID NOT NULL,
    license             TEXT NOT NULL DEFAULT 'PPL',
    language            TEXT NOT NULL DEFAULT 'fr',
    subject_id          TEXT NOT NULL,  -- ex: "010_airlaw"
    level_1_completed   BOOLEAN NOT NULL DEFAULT FALSE,
    level_2_completed   BOOLEAN NOT NULL DEFAULT FALSE,
    level_3_completed   BOOLEAN NOT NULL DEFAULT FALSE,
    score_avg_level1    FLOAT DEFAULT 0,
    score_avg_level2    FLOAT DEFAULT 0,
    score_avg_level3    FLOAT DEFAULT 0,
    questions_attempted INT NOT NULL DEFAULT 0,
    last_seen_date      TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (user_id, license, language, subject_id)
);

CREATE INDEX IF NOT EXISTS idx_user_progress_subject ON user_progress(subject_id);

-- ============================================================================
-- 3. TABLE ACHIEVEMENTS (dans aero_gamification)
-- ============================================================================
CREATE TABLE IF NOT EXISTS achievements (
    user_id     UUID NOT NULL,
    badge_id    TEXT NOT NULL,  -- ex: "airlaw_master", "7day_streak"
    level       INT NOT NULL DEFAULT 1,
    unlocked_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (user_id, badge_id)
);

-- ============================================================================
-- 4. TABLE LEADERBOARD_WEEKLY (dans aero_gamification)
-- ============================================================================
CREATE TABLE IF NOT EXISTS leaderboard_weekly (
    user_id         UUID NOT NULL,
    week_start_date DATE NOT NULL,  -- YYYY-MM-DD
    xp_earned_week  INT NOT NULL DEFAULT 0,
    rank            INT,
    tier            TEXT NOT NULL DEFAULT 'bronze',  -- bronze, silver, gold, diamond
    PRIMARY KEY (user_id, week_start_date)
);

CREATE INDEX IF NOT EXISTS idx_leaderboard_week ON leaderboard_weekly(week_start_date, rank);

-- ============================================================================
-- 5. TABLE USER_QUESTS (dans aero_gamification)
-- ============================================================================
CREATE TABLE IF NOT EXISTS user_quests (
    user_id      UUID NOT NULL,
    quest_id     TEXT NOT NULL,  -- ex: "3_lessons_day", "100q_no_error"
    progress     INT NOT NULL DEFAULT 0,
    max_progress INT NOT NULL DEFAULT 1,
    reward_xp    INT NOT NULL DEFAULT 0,
    completed    BOOLEAN NOT NULL DEFAULT FALSE,
    completed_at TIMESTAMPTZ,
    PRIMARY KEY (user_id, quest_id)
);

-- ============================================================================
-- 6. TABLE SPACED_REPETITION (dans aero_gamification)
-- ============================================================================
CREATE TABLE IF NOT EXISTS spaced_repetition (
    user_id        UUID NOT NULL,
    question_id    TEXT NOT NULL,
    level          INT NOT NULL DEFAULT 1,  -- 1=jour1, 2=jour3, 3=jour7, 4=jour14, 5=jour30
    next_review_at TIMESTAMPTZ NOT NULL,
    last_seen_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    times_correct  INT NOT NULL DEFAULT 0,
    times_wrong    INT NOT NULL DEFAULT 0,
    PRIMARY KEY (user_id, question_id)
);

CREATE INDEX IF NOT EXISTS idx_spaced_review ON spaced_repetition(next_review_at) WHERE next_review_at <= now();

-- ============================================================================
-- 7. TABLE CONFIG (dans aero_config)
-- ============================================================================
CREATE TABLE IF NOT EXISTS app_config (
    config_key   TEXT PRIMARY KEY,  -- ex: "adaptive_learning", "gamification_rules", "exam_config"
    config_value JSONB NOT NULL,
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Insertion des valeurs par défaut
INSERT INTO app_config (config_key, config_value) VALUES
('adaptive_learning', '{
    "level_1": {"pass_threshold": 0.80, "hint_enabled": true, "time_limit_sec": null, "immediate_feedback": true},
    "level_2": {"pass_threshold": 0.75, "hint_enabled": false, "time_limit_sec": null, "immediate_feedback": true},
    "level_3": {"pass_threshold": 0.75, "hint_enabled": false, "time_limit_sec": 45, "immediate_feedback": false}
}')
ON CONFLICT (config_key) DO NOTHING;

INSERT INTO app_config (config_key, config_value) VALUES
('gamification_rules', '{
    "hearts_enabled": true, "hearts_total": 5,
    "heart_penalty_level2": 1, "heart_penalty_level3": 2,
    "xp_per_lesson": 10, "xp_per_qcm_level1": 20, "xp_per_qcm_level2": 30, "xp_per_qcm_level3": 50,
    "streak_enabled": true, "leaderboard_weekly": true
}')
ON CONFLICT (config_key) DO NOTHING;

INSERT INTO app_config (config_key, config_value) VALUES
('exam_config', '{
    "ppl.air_law": {"questions_count": 16, "time_minutes": 30, "passing_score_percent": 75, "randomized": true, "no_backward": true},
    "ppl.meteorology": {"questions_count": 16, "time_minutes": 30, "passing_score_percent": 75, "randomized": true, "no_backward": true}
}')
ON CONFLICT (config_key) DO NOTHING;

-- ============================================================================
-- 8. ENRICHISSEMENT TABLE QUESTIONS EXISTANTE (dans aero_ppl_fr, aero_ppl_en, etc.)
-- ============================================================================

-- Ajouter les colonnes enrichies aux questions
ALTER TABLE questions ADD COLUMN IF NOT EXISTS level INT DEFAULT 1 CHECK (level IN (1, 2, 3));
ALTER TABLE questions ADD COLUMN IF NOT EXISTS faa_note_fr TEXT;
ALTER TABLE questions ADD COLUMN IF NOT EXISTS faa_note_en TEXT;
ALTER TABLE questions ADD COLUMN IF NOT EXISTS tags JSONB DEFAULT '[]'::jsonb;
ALTER TABLE questions ADD COLUMN IF NOT EXISTS difficulty_score FLOAT CHECK (difficulty_score IS NULL OR (difficulty_score >= 0 AND difficulty_score <= 1));
ALTER TABLE questions ADD COLUMN IF NOT EXISTS distractors_rationale JSONB;

-- Index pour les nouveaux champs
CREATE INDEX IF NOT EXISTS idx_questions_level ON questions(level);
CREATE INDEX IF NOT EXISTS idx_questions_tags ON questions USING GIN(tags);
CREATE INDEX IF NOT EXISTS idx_questions_difficulty_score ON questions(difficulty_score);

-- ============================================================================
-- 9. ENRICHISSEMENT TABLE LESSONS EXISTANTE
-- ============================================================================
ALTER TABLE lessons ADD COLUMN IF NOT EXISTS level INT DEFAULT 1 CHECK (level IN (1, 2, 3));
ALTER TABLE lessons ADD COLUMN IF NOT EXISTS duration_minutes INT DEFAULT 0;
ALTER TABLE lessons ADD COLUMN IF NOT EXISTS tags JSONB DEFAULT '[]'::jsonb;
ALTER TABLE lessons ADD COLUMN IF NOT EXISTS learning_objectives JSONB DEFAULT '[]'::jsonb;

-- ============================================================================
-- 10. ENRICHISSEMENT TABLE STUDENTS
-- ============================================================================
ALTER TABLE students ADD COLUMN IF NOT EXISTS preferred_license TEXT NOT NULL DEFAULT 'PPL';
ALTER TABLE students ADD COLUMN IF NOT EXISTS hearts INT NOT NULL DEFAULT 5 CHECK (hearts BETWEEN 0 AND 5);
ALTER TABLE students ADD COLUMN IF NOT EXISTS xp INT NOT NULL DEFAULT 0;
ALTER TABLE students ADD COLUMN IF NOT EXISTS streak INT NOT NULL DEFAULT 0;
ALTER TABLE students ADD COLUMN IF NOT EXISTS user_level INT NOT NULL DEFAULT 1;
