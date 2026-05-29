package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"aeropath/internal/domain"
)

type LessonRepo struct {
	pool *pgxpool.Pool
}

func NewLessonRepo(pool *pgxpool.Pool) *LessonRepo {
	return &LessonRepo{pool: pool}
}

const lessonColumns = `id, license, category, theme, title_fr, title_en, content_fr, content_en, difficulty, order_index, created_at`

func scanLesson(scanner interface {
	Scan(dest ...interface{}) error
}) (*domain.Lesson, error) {
	l := &domain.Lesson{}
	err := scanner.Scan(&l.ID, &l.License, &l.Category, &l.Theme,
		&l.TitleFr, &l.TitleEn, &l.ContentFr, &l.ContentEn,
		&l.Difficulty, &l.OrderIndex, &l.CreatedAt)
	return l, err
}

func (r *LessonRepo) Create(l *domain.Lesson) error {
	query := `
		INSERT INTO lessons (license, category, theme, title_fr, title_en, content_fr, content_en, difficulty, order_index)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at`

	return r.pool.QueryRow(context.Background(), query,
		l.License, l.Category, l.Theme,
		l.TitleFr, l.TitleEn, l.ContentFr, l.ContentEn,
		l.Difficulty, l.OrderIndex,
	).Scan(&l.ID, &l.CreatedAt)
}

func (r *LessonRepo) FindByID(id string) (*domain.Lesson, error) {
	query := fmt.Sprintf(`SELECT %s FROM lessons WHERE id = $1`, lessonColumns)

	l, err := scanLesson(r.pool.QueryRow(context.Background(), query, id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("leçon introuvable")
	}
	return l, err
}

func (r *LessonRepo) FindByTheme(theme string) ([]*domain.Lesson, error) {
	query := fmt.Sprintf(`SELECT %s FROM lessons WHERE theme = $1 ORDER BY order_index`, lessonColumns)
	return r.queryLessons(query, theme)
}

func (r *LessonRepo) FindAll() ([]*domain.Lesson, error) {
	query := fmt.Sprintf(`SELECT %s FROM lessons ORDER BY category, theme, order_index`, lessonColumns)
	return r.queryLessons(query)
}

func (r *LessonRepo) FindAllPaginated(limit, offset int) ([]*domain.Lesson, error) {
	query := fmt.Sprintf(`SELECT %s FROM lessons ORDER BY category, theme, order_index LIMIT $1 OFFSET $2`, lessonColumns)
	return r.queryLessons(query, limit, offset)
}

func (r *LessonRepo) FindByLicense(license domain.License) ([]*domain.Lesson, error) {
	query := fmt.Sprintf(`SELECT %s FROM lessons WHERE license = $1 ORDER BY category, theme, order_index`, lessonColumns)
	return r.queryLessons(query, license)
}

func (r *LessonRepo) FindByCategory(category domain.Category) ([]*domain.Lesson, error) {
	query := fmt.Sprintf(`SELECT %s FROM lessons WHERE category = $1 ORDER BY theme, order_index`, lessonColumns)
	return r.queryLessons(query, category)
}

func (r *LessonRepo) FindByLicenseAndCategory(license domain.License, category domain.Category) ([]*domain.Lesson, error) {
	query := fmt.Sprintf(`SELECT %s FROM lessons WHERE license = $1 AND category = $2 ORDER BY order_index`, lessonColumns)
	return r.queryLessons(query, license, category)
}

func (r *LessonRepo) FindByDifficulty(difficulty int) ([]*domain.Lesson, error) {
	query := fmt.Sprintf(`SELECT %s FROM lessons WHERE difficulty = $1 ORDER BY category, theme, order_index`, lessonColumns)
	return r.queryLessons(query, difficulty)
}

func (r *LessonRepo) CountByLicense(license domain.License) (int, error) {
	var count int
	err := r.pool.QueryRow(context.Background(), `SELECT COUNT(*) FROM lessons WHERE license = $1`, license).Scan(&count)
	return count, err
}

func (r *LessonRepo) CountByCategory(category domain.Category) (int, error) {
	var count int
	err := r.pool.QueryRow(context.Background(), `SELECT COUNT(*) FROM lessons WHERE category = $1`, category).Scan(&count)
	return count, err
}

func (r *LessonRepo) Update(l *domain.Lesson) error {
	query := `
		UPDATE lessons SET
			license = $1, category = $2, theme = $3,
			title_fr = $4, title_en = $5,
			content_fr = $6, content_en = $7,
			difficulty = $8, order_index = $9
		WHERE id = $10`

	tag, err := r.pool.Exec(context.Background(), query,
		l.License, l.Category, l.Theme,
		l.TitleFr, l.TitleEn, l.ContentFr, l.ContentEn,
		l.Difficulty, l.OrderIndex,
		l.ID,
	)
	if err != nil {
		return fmt.Errorf("update lesson: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("leçon introuvable")
	}
	return nil
}

func (r *LessonRepo) Delete(id string) error {
	tag, err := r.pool.Exec(context.Background(), `DELETE FROM lessons WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete lesson: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("leçon introuvable")
	}
	return nil
}

func (r *LessonRepo) Count() (int, error) {
	var count int
	err := r.pool.QueryRow(context.Background(), `SELECT COUNT(*) FROM lessons`).Scan(&count)
	return count, err
}

func (r *LessonRepo) queryLessons(query string, args ...interface{}) ([]*domain.Lesson, error) {
	rows, err := r.pool.Query(context.Background(), query, args...)
	if err != nil {
		return nil, fmt.Errorf("query lessons: %w", err)
	}
	defer rows.Close()

	var lessons []*domain.Lesson
	for rows.Next() {
		l, err := scanLesson(rows)
		if err != nil {
			return nil, fmt.Errorf("scan lesson: %w", err)
		}
		lessons = append(lessons, l)
	}
	return lessons, rows.Err()
}
