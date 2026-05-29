package postgres

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// TestPool est un pool de connexion pour les tests d'intégration.
// Il crée un schéma de test isolé et le nettoie à la fermeture.
type TestPool struct {
	*pgxpool.Pool
	schema string
}

// NewTestPool crée un pool de test avec un schéma isolé.
func NewTestPool(dsn string) (*TestPool, error) {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("new pool: %w", err)
	}

	schema := "test_" + randString(8)

	// Créer un schéma de test isolé
	_, err = pool.Exec(context.Background(), fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schema))
	if err != nil {
		pool.Close()
		return nil, fmt.Errorf("create schema: %w", err)
	}

	// Définir le search_path pour utiliser ce schéma
	_, err = pool.Exec(context.Background(), fmt.Sprintf("SET search_path TO %s, public", schema))
	if err != nil {
		pool.Close()
		return nil, fmt.Errorf("set search_path: %w", err)
	}

	// Copier la structure des tables dans le schéma de test
	tables := []string{
		`CREATE TABLE IF NOT EXISTS students (
			id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			email      TEXT NOT NULL UNIQUE,
			password   TEXT NOT NULL,
			lang       TEXT NOT NULL DEFAULT 'fr',
			created_at TIMESTAMPTZ NOT NULL DEFAULT now()
		)`,
		`CREATE TABLE IF NOT EXISTS questions (
			id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			lesson_id      UUID,
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
		)`,
		`CREATE TABLE IF NOT EXISTS student_question_history (
			student_id  UUID NOT NULL REFERENCES students(id) ON DELETE CASCADE,
			question_id UUID NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
			seen_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
			was_correct BOOLEAN NOT NULL,
			PRIMARY KEY (student_id, question_id)
		)`,
		`CREATE TABLE IF NOT EXISTS lessons (
			id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			theme       TEXT NOT NULL,
			title_fr    TEXT NOT NULL,
			title_en    TEXT NOT NULL,
			content_fr  TEXT NOT NULL,
			content_en  TEXT NOT NULL,
			difficulty  SMALLINT NOT NULL DEFAULT 1 CHECK (difficulty BETWEEN 1 AND 5),
			order_index INT NOT NULL DEFAULT 0,
			created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
		)`,
	}

	for _, ddl := range tables {
		if _, err := pool.Exec(context.Background(), ddl); err != nil {
			pool.Close()
			return nil, fmt.Errorf("create table: %w", err)
		}
	}

	return &TestPool{Pool: pool, schema: schema}, nil
}

// Close ferme le pool et nettoie le schéma de test.
func (tp *TestPool) Close() {
	if tp.schema != "" {
		tp.Exec(context.Background(), fmt.Sprintf("DROP SCHEMA IF EXISTS %s CASCADE", tp.schema))
	}
	tp.Pool.Close()
}

// randString génère une chaîne aléatoire de longueur n.
func randString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rng.Intn(len(letters))]
	}
	return string(b)
}
