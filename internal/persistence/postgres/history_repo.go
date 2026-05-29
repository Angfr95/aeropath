package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"aeropath/internal/domain"
)

type HistoryRepo struct {
	pool *pgxpool.Pool
}

func NewHistoryRepo(pool *pgxpool.Pool) *HistoryRepo {
	return &HistoryRepo{pool: pool}
}

func (r *HistoryRepo) RecordAnswer(studentID, questionID string, wasCorrect bool) error {
	query := `
		INSERT INTO student_question_history (student_id, question_id, was_correct)
		VALUES ($1, $2, $3)
		ON CONFLICT (student_id, question_id) DO UPDATE SET
			seen_at = now(),
			was_correct = $3`

	_, err := r.pool.Exec(context.Background(), query, studentID, questionID, wasCorrect)
	if err != nil {
		return fmt.Errorf("record answer: %w", err)
	}
	return nil
}

func (r *HistoryRepo) GetHistory(studentID string) ([]*domain.QuestionHistory, error) {
	query := `SELECT student_id, question_id, seen_at, was_correct FROM student_question_history WHERE student_id = $1 ORDER BY seen_at DESC`

	rows, err := r.pool.Query(context.Background(), query, studentID)
	if err != nil {
		return nil, fmt.Errorf("get history: %w", err)
	}
	defer rows.Close()

	var history []*domain.QuestionHistory
	for rows.Next() {
		h := &domain.QuestionHistory{}
		if err := rows.Scan(&h.StudentID, &h.QuestionID, &h.SeenAt, &h.WasCorrect); err != nil {
			return nil, fmt.Errorf("scan history: %w", err)
		}
		history = append(history, h)
	}
	return history, rows.Err()
}

func (r *HistoryRepo) GetStats(studentID string) (*domain.StudentStats, error) {
	query := `
		SELECT
			COUNT(*)::int AS total,
			COALESCE(SUM(CASE WHEN was_correct THEN 1 ELSE 0 END), 0)::int AS correct,
			COALESCE(SUM(CASE WHEN NOT was_correct THEN 1 ELSE 0 END), 0)::int AS wrong
		FROM student_question_history
		WHERE student_id = $1`

	stats := &domain.StudentStats{}
	err := r.pool.QueryRow(context.Background(), query, studentID).
		Scan(&stats.TotalQuestions, &stats.CorrectAnswers, &stats.WrongAnswers)
	if err != nil {
		return nil, fmt.Errorf("get stats: %w", err)
	}

	if stats.TotalQuestions > 0 {
		stats.SuccessRate = float64(stats.CorrectAnswers) / float64(stats.TotalQuestions) * 100
	}

	// Niveau actuel basé sur la difficulté moyenne des questions réussies
	levelQuery := `
		SELECT COALESCE(ROUND(AVG(q.difficulty)), 1)::int
		FROM student_question_history h
		JOIN questions q ON q.id = h.question_id
		WHERE h.student_id = $1 AND h.was_correct = true`

	err = r.pool.QueryRow(context.Background(), levelQuery, studentID).Scan(&stats.CurrentLevel)
	if err != nil {
		stats.CurrentLevel = 1 // Par défaut niveau 1
	}

	return stats, nil
}

func (r *HistoryRepo) Count() (int, error) {
	var count int
	err := r.pool.QueryRow(context.Background(), `SELECT COUNT(*) FROM student_question_history`).Scan(&count)
	return count, err
}

func (r *HistoryRepo) GetSeenQuestionIDs(studentID string) ([]string, error) {
	query := `SELECT question_id FROM student_question_history WHERE student_id = $1`

	rows, err := r.pool.Query(context.Background(), query, studentID)
	if err != nil {
		return nil, fmt.Errorf("get seen questions: %w", err)
	}
	defer rows.Close()

	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("scan question id: %w", err)
		}
		ids = append(ids, id)
	}
	return ids, rows.Err()
}
