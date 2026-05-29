package postgres

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"aeropath/internal/domain"
)

// QuestionRepo implémente domain.QuestionRepository avec PostgreSQL.
//
// 📖 DDIA Chapitre 3 : "Storage and Retrieval"
//    Les questions sont stockées dans une table PostgreSQL avec des index.
//    Les index accélèrent les recherches fréquentes :
//    - Index sur (license, category) pour "questions PPL de météo"
//    - Index sur (theme) pour "toutes les questions de navigation"
//    - Index sur (difficulty) pour "questions niveau 3"
//
//    Sans index, PostgreSQL ferait un "full table scan" (lire TOUTE la table).
//    Avec 10 000 questions, un full scan prend ~50ms.
//    Avec un index, la même requête prend ~0.1ms.
//
// 📖 DDIA Chapitre 4 : "Encoding and Evolution"
//    Les options sont stockées en JSON dans PostgreSQL.
//    Pourquoi pas une table séparée "options" ?
//    - Les options n'ont pas besoin d'être requêtées individuellement
//    - On les charge toujours avec la question
//    - C'est plus simple et plus rapide (pas de JOIN)
//    C'est le principe de "schemaless" : on utilise JSON pour
//    les données qui n'ont pas besoin d'être filtrées.
type QuestionRepo struct {
	pool *pgxpool.Pool
}

func NewQuestionRepo(pool *pgxpool.Pool) *QuestionRepo {
	return &QuestionRepo{pool: pool}
}

const questionColumns = `id, lesson_id, license, category, theme, subtopic, difficulty, question_fr, question_en, options, answer_key, explanation_fr, explanation_en, created_at`

func scanQuestion(scanner interface {
	Scan(dest ...interface{}) error
}) (*domain.Question, error) {
	q := &domain.Question{}
	var optionsJSON []byte
	var lessonID *string
	err := scanner.Scan(&q.ID, &lessonID, &q.License, &q.Category,
		&q.Theme, &q.Subtopic, &q.Difficulty,
		&q.QuestionFr, &q.QuestionEn, &optionsJSON, &q.AnswerKey,
		&q.ExplanationFr, &q.ExplanationEn, &q.CreatedAt)
	if err != nil {
		return nil, err
	}
	if lessonID != nil {
		q.LessonID = *lessonID
	}
	if err := json.Unmarshal(optionsJSON, &q.Options); err != nil {
		return nil, fmt.Errorf("unmarshal options: %w", err)
	}
	return q, nil
}


func (r *QuestionRepo) Create(q *domain.Question) error {
	optionsJSON, err := json.Marshal(q.Options)
	if err != nil {
		return fmt.Errorf("marshal options: %w", err)
	}

	query := `
		INSERT INTO questions (lesson_id, license, category, theme, subtopic, difficulty, question_fr, question_en, options, answer_key, explanation_fr, explanation_en)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id, created_at`

	lessonID := &q.LessonID
	if q.LessonID == "" {
		lessonID = nil
	}

	return r.pool.QueryRow(context.Background(), query,
		lessonID, q.License, q.Category,
		q.Theme, q.Subtopic, q.Difficulty,
		q.QuestionFr, q.QuestionEn,
		optionsJSON, q.AnswerKey,
		q.ExplanationFr, q.ExplanationEn,
	).Scan(&q.ID, &q.CreatedAt)
}


func (r *QuestionRepo) FindByID(id string) (*domain.Question, error) {
	query := fmt.Sprintf(`SELECT %s FROM questions WHERE id = $1`, questionColumns)

	q, err := scanQuestion(r.pool.QueryRow(context.Background(), query, id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("question introuvable")
	}
	return q, err
}

func (r *QuestionRepo) FindAll() ([]*domain.Question, error) {
	query := fmt.Sprintf(`SELECT %s FROM questions ORDER BY created_at DESC`, questionColumns)
	return r.queryQuestions(query)
}

func (r *QuestionRepo) FindAllPaginated(limit, offset int) ([]*domain.Question, error) {
	query := fmt.Sprintf(`SELECT %s FROM questions ORDER BY created_at DESC LIMIT $1 OFFSET $2`, questionColumns)
	return r.queryQuestions(query, limit, offset)
}

func (r *QuestionRepo) FindByTheme(theme string) ([]*domain.Question, error) {
	query := fmt.Sprintf(`SELECT %s FROM questions WHERE theme = $1 ORDER BY difficulty`, questionColumns)
	return r.queryQuestions(query, theme)
}

func (r *QuestionRepo) FindByLicense(license domain.License) ([]*domain.Question, error) {
	query := fmt.Sprintf(`SELECT %s FROM questions WHERE license = $1 ORDER BY category, theme, difficulty`, questionColumns)
	return r.queryQuestions(query, license)
}

func (r *QuestionRepo) FindByCategory(category domain.Category) ([]*domain.Question, error) {
	query := fmt.Sprintf(`SELECT %s FROM questions WHERE category = $1 ORDER BY theme, difficulty`, questionColumns)
	return r.queryQuestions(query, category)
}

func (r *QuestionRepo) FindByLicenseAndCategory(license domain.License, category domain.Category) ([]*domain.Question, error) {
	query := fmt.Sprintf(`SELECT %s FROM questions WHERE license = $1 AND category = $2 ORDER BY difficulty`, questionColumns)
	return r.queryQuestions(query, license, category)
}

func (r *QuestionRepo) FindByDifficulty(difficulty int) ([]*domain.Question, error) {
	query := fmt.Sprintf(`SELECT %s FROM questions WHERE difficulty = $1 ORDER BY category, theme`, questionColumns)
	return r.queryQuestions(query, difficulty)
}

func (r *QuestionRepo) FindBySubtopic(subtopic string) ([]*domain.Question, error) {
	query := fmt.Sprintf(`SELECT %s FROM questions WHERE subtopic = $1 ORDER BY difficulty`, questionColumns)
	return r.queryQuestions(query, subtopic)
}

func (r *QuestionRepo) CountByLicense(license domain.License) (int, error) {
	var count int
	err := r.pool.QueryRow(context.Background(), `SELECT COUNT(*) FROM questions WHERE license = $1`, license).Scan(&count)
	return count, err
}

func (r *QuestionRepo) CountByCategory(category domain.Category) (int, error) {
	var count int
	err := r.pool.QueryRow(context.Background(), `SELECT COUNT(*) FROM questions WHERE category = $1`, category).Scan(&count)
	return count, err
}

func (r *QuestionRepo) CountByTheme(theme string) (int, error) {
	var count int
	err := r.pool.QueryRow(context.Background(), `SELECT COUNT(*) FROM questions WHERE theme = $1`, theme).Scan(&count)
	return count, err
}

func (r *QuestionRepo) Update(q *domain.Question) error {
	optionsJSON, err := json.Marshal(q.Options)
	if err != nil {
		return fmt.Errorf("marshal options: %w", err)
	}

	lessonID := &q.LessonID
	if q.LessonID == "" {
		lessonID = nil
	}

	query := `
		UPDATE questions SET
			lesson_id = $1, license = $2, category = $3,
			theme = $4, subtopic = $5, difficulty = $6,
			question_fr = $7, question_en = $8,
			options = $9, answer_key = $10,
			explanation_fr = $11, explanation_en = $12
		WHERE id = $13`

	tag, err := r.pool.Exec(context.Background(), query,
		lessonID, q.License, q.Category,
		q.Theme, q.Subtopic, q.Difficulty,
		q.QuestionFr, q.QuestionEn,
		optionsJSON, q.AnswerKey,
		q.ExplanationFr, q.ExplanationEn,
		q.ID,
	)
	if err != nil {
		return fmt.Errorf("update question: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("question introuvable")
	}
	return nil
}

func (r *QuestionRepo) Delete(id string) error {
	tag, err := r.pool.Exec(context.Background(), `DELETE FROM questions WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete question: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("question introuvable")
	}
	return nil
}

func (r *QuestionRepo) Count() (int, error) {
	var count int
	err := r.pool.QueryRow(context.Background(), `SELECT COUNT(*) FROM questions`).Scan(&count)
	return count, err
}

// FindRandom retourne des questions aléatoires avec filtres optionnels.
//
// 📖 DDIA Chapitre 1 : "Scalability"
//    ORDER BY RANDOM() est lent sur de grandes tables car PostgreSQL
//    doit charger TOUTES les lignes pour les mélanger.
//    Pour 10 000 questions, ça reste acceptable (< 100ms).
//    Pour 1 million de questions, il faudrait une autre approche
//    (ex: sélectionner un ID aléatoire avec un index).
func (r *QuestionRepo) FindRandom(limit int, license *domain.License, category *domain.Category, theme *string) ([]*domain.Question, error) {
	query := fmt.Sprintf(`SELECT %s FROM questions WHERE 1=1`, questionColumns)
	var args []interface{}
	argIdx := 1

	if license != nil {
		query += fmt.Sprintf(` AND license = $%d`, argIdx)
		args = append(args, *license)
		argIdx++
	}
	if category != nil {
		query += fmt.Sprintf(` AND category = $%d`, argIdx)
		args = append(args, *category)
		argIdx++
	}
	if theme != nil && *theme != "" {
		query += fmt.Sprintf(` AND theme = $%d`, argIdx)
		args = append(args, *theme)
		argIdx++
	}

	query += ` ORDER BY RANDOM()`
	if limit > 0 {
		query += fmt.Sprintf(` LIMIT $%d`, argIdx)
		args = append(args, limit)
	}

	return r.queryQuestions(query, args...)
}

// Search cherche des questions par texte libre.
//
// 📖 DDIA Chapitre 3 : "Storage and Retrieval"
//    On utilise la recherche full-text PostgreSQL (to_tsvector).
//    C'est plus intelligent qu'un simple LIKE :
//    - to_tsvector('french', question_fr) crée un index de mots
//    - plainto_tsquery($1) normalise la recherche
//    - @@ fait la correspondance
//
//    Exemple : chercher "altimètre" trouve aussi "altimètres"
//    (grâce à la normalisation morphologique).
//
//    On combine avec ILIKE pour les cas où la recherche
//    full-text ne suffit pas (mots partiels, codes, etc.).
func (r *QuestionRepo) Search(query string) ([]*domain.Question, error) {
	sqlQuery := fmt.Sprintf(`SELECT %s FROM questions WHERE
		to_tsvector('french', COALESCE(question_fr, '')) || to_tsvector('english', COALESCE(question_en, '')) @@ plainto_tsquery($1)
		OR question_fr ILIKE '%%%%' || $1 || '%%%%'
		OR question_en ILIKE '%%%%' || $1 || '%%%%'
		OR theme ILIKE '%%%%' || $1 || '%%%%'
		OR subtopic ILIKE '%%%%' || $1 || '%%%%'
		ORDER BY difficulty`, questionColumns)
	return r.queryQuestions(sqlQuery, query)
}

func (r *QuestionRepo) queryQuestions(query string, args ...interface{}) ([]*domain.Question, error) {
	rows, err := r.pool.Query(context.Background(), query, args...)
	if err != nil {
		return nil, fmt.Errorf("query questions: %w", err)
	}
	defer rows.Close()

	var questions []*domain.Question
	for rows.Next() {
		q, err := scanQuestion(rows)
		if err != nil {
			return nil, fmt.Errorf("scan question: %w", err)
		}
		questions = append(questions, q)
	}
	return questions, rows.Err()
}
