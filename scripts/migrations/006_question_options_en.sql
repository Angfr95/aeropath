-- Keep answer positions identical across languages; answer_key remains canonical.
ALTER TABLE questions ADD COLUMN IF NOT EXISTS options_en JSONB NOT NULL DEFAULT '[]'::jsonb;
