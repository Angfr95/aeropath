-- Ajouter les colonnes license et category aux tables questions et lessons
-- Ces colonnes permettent de filtrer par licence (PPL, ATPL, etc.) et catégorie (Meteorology, Navigation, etc.)

ALTER TABLE questions ADD COLUMN IF NOT EXISTS license TEXT NOT NULL DEFAULT 'PPL';
ALTER TABLE questions ADD COLUMN IF NOT EXISTS category TEXT NOT NULL DEFAULT 'General';

ALTER TABLE lessons ADD COLUMN IF NOT EXISTS license TEXT NOT NULL DEFAULT 'PPL';
ALTER TABLE lessons ADD COLUMN IF NOT EXISTS category TEXT NOT NULL DEFAULT 'General';

-- Index pour accélérer les recherches par licence et catégorie
CREATE INDEX IF NOT EXISTS idx_questions_license_category ON questions(license, category);
CREATE INDEX IF NOT EXISTS idx_lessons_license_category ON lessons(license, category);
