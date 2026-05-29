CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE students (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email         TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    lang          TEXT NOT NULL DEFAULT 'fr',  -- 'fr' | 'en'
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE questions (
    id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    theme          TEXT NOT NULL,
    subtopic       TEXT,
    difficulty     SMALLINT NOT NULL CHECK (difficulty BETWEEN 1 AND 5),
    question_fr    TEXT NOT NULL,
    question_en    TEXT NOT NULL,
    options        JSONB NOT NULL,
    answer_key     TEXT NOT NULL,
    explanation_fr TEXT,
    explanation_en TEXT,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE student_question_history (
    student_id  UUID NOT NULL REFERENCES students(id) ON DELETE CASCADE,
    question_id UUID NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    seen_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    was_correct BOOLEAN NOT NULL,
    PRIMARY KEY (student_id, question_id)
);

CREATE INDEX idx_history_student_seen ON student_question_history(student_id, seen_at);