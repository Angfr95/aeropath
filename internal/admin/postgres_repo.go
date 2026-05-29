package admin

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresRepository implémente Repository avec PostgreSQL
type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

func (r *PostgresRepository) CountStudents(ctx context.Context) (int, error) {
	var count int
	err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM students`).Scan(&count)
	return count, err
}

func (r *PostgresRepository) CountQuestions(ctx context.Context) (int, error) {
	var count int
	err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM questions`).Scan(&count)
	return count, err
}

func (r *PostgresRepository) CountLessons(ctx context.Context) (int, error) {
	var count int
	err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM lessons`).Scan(&count)
	return count, err
}

func (r *PostgresRepository) CountAnswers(ctx context.Context) (int, error) {
	var count int
	err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM answer_history`).Scan(&count)
	return count, err
}

func (r *PostgresRepository) CountCorrectAnswers(ctx context.Context) (int, error) {
	var count int
	err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM answer_history WHERE was_correct = true`).Scan(&count)
	return count, err
}

func (r *PostgresRepository) GetQuestionsByLicense(ctx context.Context) (map[string]int, error) {
	rows, err := r.pool.Query(ctx, `SELECT license, COUNT(*) FROM questions GROUP BY license ORDER BY license`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]int)
	for rows.Next() {
		var license string
		var count int
		if err := rows.Scan(&license, &count); err != nil {
			return nil, err
		}
		result[license] = count
	}
	return result, nil
}

func (r *PostgresRepository) GetQuestionsByCategory(ctx context.Context) (map[string]int, error) {
	rows, err := r.pool.Query(ctx, `SELECT category, COUNT(*) FROM questions GROUP BY category ORDER BY category`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]int)
	for rows.Next() {
		var category string
		var count int
		if err := rows.Scan(&category, &count); err != nil {
			return nil, err
		}
		result[category] = count
	}
	return result, nil
}

func (r *PostgresRepository) GetStudentsByLang(ctx context.Context) (map[string]int, error) {
	rows, err := r.pool.Query(ctx, `SELECT lang, COUNT(*) FROM students GROUP BY lang ORDER BY lang`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]int)
	for rows.Next() {
		var lang string
		var count int
		if err := rows.Scan(&lang, &count); err != nil {
			return nil, err
		}
		result[lang] = count
	}
	return result, nil
}

func (r *PostgresRepository) GetAnswersLast7Days(ctx context.Context) ([]DailyCount, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT DATE(answered_at) as date, COUNT(*) as count
		FROM answer_history
		WHERE answered_at >= NOW() - INTERVAL '7 days'
		GROUP BY DATE(answered_at)
		ORDER BY date
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []DailyCount
	for rows.Next() {
		var dc DailyCount
		var t time.Time
		if err := rows.Scan(&t, &dc.Count); err != nil {
			return nil, err
		}
		dc.Date = t.Format("2006-01-02")
		result = append(result, dc)
	}
	return result, nil
}

func (r *PostgresRepository) GetRegistrationsByDay(ctx context.Context, days int) ([]DailyCount, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT DATE(created_at) as date, COUNT(*) as count
		FROM students
		WHERE created_at >= NOW() - $1::INTERVAL
		GROUP BY DATE(created_at)
		ORDER BY date
	`, fmt.Sprintf("%d days", days))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []DailyCount
	for rows.Next() {
		var dc DailyCount
		var t time.Time
		if err := rows.Scan(&t, &dc.Count); err != nil {
			return nil, err
		}
		dc.Date = t.Format("2006-01-02")
		result = append(result, dc)
	}
	return result, nil
}

func (r *PostgresRepository) GetTopWeakTopics(ctx context.Context, limit int) ([]TopicStat, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT q.theme,
			   AVG(CASE WHEN ah.was_correct THEN 100.0 ELSE 0.0 END) as score,
			   COUNT(*) as count
		FROM answer_history ah
		JOIN questions q ON q.id = ah.question_id
		GROUP BY q.theme
		ORDER BY score ASC
		LIMIT $1
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []TopicStat
	for rows.Next() {
		var ts TopicStat
		if err := rows.Scan(&ts.Theme, &ts.Score, &ts.Count); err != nil {
			return nil, err
		}
		result = append(result, ts)
	}
	return result, nil
}

func (r *PostgresRepository) ListStudents(ctx context.Context, offset, limit int) ([]Student, int, error) {
	total, err := r.CountStudents(ctx)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.pool.Query(ctx, `
		SELECT id, email, lang, created_at, updated_at
		FROM students
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var s Student
		if err := rows.Scan(&s.ID, &s.Email, &s.Lang, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, 0, err
		}
		students = append(students, s)
	}
	return students, total, nil
}

func (r *PostgresRepository) GetStudentByID(ctx context.Context, id string) (*Student, error) {
	var s Student
	err := r.pool.QueryRow(ctx, `
		SELECT id, email, lang, created_at, updated_at
		FROM students WHERE id = $1
	`, id).Scan(&s.ID, &s.Email, &s.Lang, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("student not found")
		}
		return nil, err
	}
	return &s, nil
}

func (r *PostgresRepository) DeleteStudent(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM students WHERE id = $1`, id)
	return err
}

func (r *PostgresRepository) CreateQuestion(ctx context.Context, q *Question) error {
	_, err := r.pool.Exec(ctx, `
		INSERT INTO questions (id, question_fr, question_en, options, answer_key,
			explanation_fr, explanation_en, license, category, theme, subtopic, difficulty, reference)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
	`, q.ID, q.QuestionFr, q.QuestionEn, q.Options, q.AnswerKey,
		q.ExplanationFr, q.ExplanationEn, q.License, q.Category, q.Theme,
		q.Subtopic, q.Difficulty, q.Reference)
	return err
}

func (r *PostgresRepository) UpdateQuestion(ctx context.Context, q *Question) error {
	_, err := r.pool.Exec(ctx, `
		UPDATE questions SET question_fr=$1, question_en=$2, options=$3, answer_key=$4,
			explanation_fr=$5, explanation_en=$6, license=$7, category=$8, theme=$9,
			subtopic=$10, difficulty=$11, reference=$12
		WHERE id=$13
	`, q.QuestionFr, q.QuestionEn, q.Options, q.AnswerKey,
		q.ExplanationFr, q.ExplanationEn, q.License, q.Category, q.Theme,
		q.Subtopic, q.Difficulty, q.Reference, q.ID)
	return err
}

func (r *PostgresRepository) DeleteQuestion(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM questions WHERE id = $1`, id)
	return err
}

func (r *PostgresRepository) CreateLesson(ctx context.Context, l *Lesson) error {
	_, err := r.pool.Exec(ctx, `
		INSERT INTO lessons (id, title_fr, title_en, content_fr, content_en,
			license, category, theme, difficulty, order_index)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
	`, l.ID, l.TitleFr, l.TitleEn, l.ContentFr, l.ContentEn,
		l.License, l.Category, l.Theme, l.Difficulty, l.OrderIndex)
	return err
}

func (r *PostgresRepository) UpdateLesson(ctx context.Context, l *Lesson) error {
	_, err := r.pool.Exec(ctx, `
		UPDATE lessons SET title_fr=$1, title_en=$2, content_fr=$3, content_en=$4,
			license=$5, category=$6, theme=$7, difficulty=$8, order_index=$9
		WHERE id=$10
	`, l.TitleFr, l.TitleEn, l.ContentFr, l.ContentEn,
		l.License, l.Category, l.Theme, l.Difficulty, l.OrderIndex, l.ID)
	return err
}

func (r *PostgresRepository) DeleteLesson(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM lessons WHERE id = $1`, id)
	return err
}
