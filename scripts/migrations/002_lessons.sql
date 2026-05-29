-- Leçons : contenu éducatif organisé par thème
CREATE TABLE lessons (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    theme       TEXT NOT NULL,          -- ex: 'meteorology', 'navigation'
    title_fr    TEXT NOT NULL,
    title_en    TEXT NOT NULL,
    content_fr  TEXT NOT NULL,          -- Contenu markdown de la leçon en français
    content_en  TEXT NOT NULL,          -- Contenu markdown de la leçon en anglais
    difficulty  SMALLINT NOT NULL DEFAULT 1 CHECK (difficulty BETWEEN 1 AND 5),
    order_index INT NOT NULL DEFAULT 0, -- Ordre d'affichage dans le thème
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Lier les questions aux leçons (optionnel : une question peut être liée à une leçon)
ALTER TABLE questions ADD COLUMN lesson_id UUID REFERENCES lessons(id) ON DELETE SET NULL;

-- Index pour retrouver les leçons par thème
CREATE INDEX idx_lessons_theme ON lessons(theme, order_index);
