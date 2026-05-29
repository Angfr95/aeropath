package postgres

import (
    "context"
    "errors"
    "fmt"

    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"

    "aeropath/internal/domain"
)

// StudentRepo implémente domain.StudentRepository avec PostgreSQL.
//
// 📖 DDIA Chapitre 3 : "Storage and Retrieval"
//    PostgreSQL est une base de données relationnelle (lignes/colonnes).
//    Chaque étudiant est une ligne dans la table "students".
//    On utilise des "prepared statements" ($1, $2) pour éviter
//    les injections SQL et optimiser les requêtes répétées.
//
// 📖 DDIA Chapitre 5 : "Replication"
//    En production, on pourrait avoir plusieurs réplicas PostgreSQL :
//    - Un "primary" qui écrit (INSERT, UPDATE, DELETE)
//    - Des "replicas" qui lisent (SELECT)
//    Les requêtes de lecture (FindByID, FindByEmail) iraient
//    sur les replicas pour répartir la charge.
type StudentRepo struct {
    pool *pgxpool.Pool
}

func NewStudentRepo(pool *pgxpool.Pool) *StudentRepo {
    return &StudentRepo{pool: pool}
}

func (r *StudentRepo) Create(s *domain.Student) error {
    query := `
        INSERT INTO students (email, password_hash, lang)
        VALUES ($1, $2, $3)
        RETURNING id, created_at`

    return r.pool.QueryRow(context.Background(), query,
        s.Email, s.PasswordHash, s.Lang,
    ).Scan(&s.ID, &s.CreatedAt)
}

func (r *StudentRepo) FindByEmail(email string) (*domain.Student, error) {
    s := &domain.Student{}
    query := `SELECT id, email, password_hash, lang, created_at FROM students WHERE email = $1`

    err := r.pool.QueryRow(context.Background(), query, email).
        Scan(&s.ID, &s.Email, &s.PasswordHash, &s.Lang, &s.CreatedAt)
    if errors.Is(err, pgx.ErrNoRows) {
        return nil, fmt.Errorf("student introuvable")
    }
    return s, err
}

func (r *StudentRepo) FindByID(id string) (*domain.Student, error) {
    s := &domain.Student{}
    query := `SELECT id, email, password_hash, lang, created_at FROM students WHERE id = $1`

    err := r.pool.QueryRow(context.Background(), query, id).
        Scan(&s.ID, &s.Email, &s.PasswordHash, &s.Lang, &s.CreatedAt)
    if errors.Is(err, pgx.ErrNoRows) {
        return nil, fmt.Errorf("student introuvable")
    }
    return s, err
}

func (r *StudentRepo) Delete(id string) error {
	tag, err := r.pool.Exec(context.Background(), `DELETE FROM students WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete student: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("étudiant introuvable")
	}
	return nil
}

func (r *StudentRepo) Count() (int, error) {
	var count int
	err := r.pool.QueryRow(context.Background(), `SELECT COUNT(*) FROM students`).Scan(&count)
	return count, err
}

func (r *StudentRepo) UpdateLang(id, lang string) error {
    query := `UPDATE students SET lang = $1 WHERE id = $2`
    tag, err := r.pool.Exec(context.Background(), query, lang, id)
    if err != nil {
        return fmt.Errorf("update lang: %w", err)
    }
    if tag.RowsAffected() == 0 {
        return fmt.Errorf("student introuvable")
    }
    return nil
}
