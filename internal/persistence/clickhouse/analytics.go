package clickhouse

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

// ============================================================
//  📊 ANALYTICS CLICKHOUSE — Statistiques et rapports
// ============================================================
// 📖 DDIA Chapitre 3 : "Storage and Retrieval" (Column-Oriented Storage)
//
// ❓ POURQUOI CLICKHOUSE ET PAS POSTGRESQL ?
//    PostgreSQL = base "lignes" (row-oriented)
//    ClickHouse = base "colonnes" (column-oriented)
//
//    Imagine la table "réponses des étudiants" :
//    | étudiant | question | correct | date       |
//    |----------|----------|---------|------------|
//    | Alice    | Q1       | oui     | 2024-01-01 |
//    | Bob      | Q2       | non     | 2024-01-02 |
//    | Alice    | Q3       | oui     | 2024-01-03 |
//
//    Pour calculer "combien de réponses correctes par jour ?"
//    PostgreSQL lit TOUTES les colonnes de TOUTES les lignes
//    ClickHouse lit seulement la colonne "correct" et "date"
//    → 3x moins de données à lire → 3x plus rapide
//
//    Pour des millions de réponses, ClickHouse est 10-100x plus rapide
//    pour les requêtes d'analyse (agrégations, GROUP BY, etc.)
//
// 🧠 QUAND UTILISER CLICKHOUSE ?
//    ✅ Rapports hebdomadaires : "taux de réussite par thème"
//    ✅ Tableaux de bord : "étudiants actifs aujourd'hui"
//    ✅ Analytics : "questions les plus ratées"
//    ❌ Transactions : "créer un étudiant" → PostgreSQL
//    ❌ Recherche en temps réel : "chercher une question" → PostgreSQL
//
// 🔗 LIENS UTILES :
//    - https://clickhouse.com/docs
//    - github.com/ClickHouse/clickhouse-go
// ============================================================

// Analytics gère les connexions à ClickHouse pour les stats.
type Analytics struct {
	conn driver.Conn
}

// NewAnalytics crée une connexion à ClickHouse.
//
// 🛡️ PRODUCTION :
//    - Pool de connexions : jusqu'à 10 connexions simultanées
//    - Timeout dial : 5 secondes max pour établir la connexion
//    - Timeout exécution : 30 secondes max pour les requêtes longues
//    - Ping de vérification au démarrage
//
// 📝 EXEMPLE D'UTILISATION :
//    analytics, err := clickhouse.NewAnalytics("localhost:9000")
//    if err != nil { log.Fatal(err) }
//    defer analytics.Close()
//
//    // Enregistrer une réponse
//    analytics.RecordAnswer("student123", "question456", true)
//
//    // Lire les stats
//    stats, err := analytics.GetStudentStats("student123")
func NewAnalytics(host string) (*Analytics, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{host}, // "localhost:9000"
		Auth: clickhouse.Auth{
			Database: "aeroforge", // Base de données ClickHouse
			Username: "default",
			Password: "",
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 30, // Timeout max des requêtes (secondes)
		},
		DialTimeout:  5 * time.Second, // Timeout connexion
		MaxOpenConns: 10,              // Max connexions simultanées
		MaxIdleConns: 5,               // Connexions inactives gardées en vie
		ConnMaxLifetime: time.Hour,    // Durée de vie max d'une connexion
	})
	if err != nil {
		return nil, fmt.Errorf("connexion ClickHouse échouée: %w", err)
	}

	// Vérifier que ClickHouse répond
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := conn.Ping(ctx); err != nil {
		conn.Close()
		return nil, fmt.Errorf("ping ClickHouse échoué: %w", err)
	}

	log.Println("📊 Analytics ClickHouse connecté")
	return &Analytics{conn: conn}, nil
}

// EnsureTable crée la table answer_events si elle n'existe pas.
// À appeler au démarrage de l'application.
//
// 🛡️ PRODUCTION : on vérifie que la table existe avant d'insérer.
// Si la table n'existe pas, on la crée automatiquement.
func (a *Analytics) EnsureTable() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := a.conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS answer_events (
			student_id  String,
			question_id String,
			was_correct Bool,
			timestamp   DateTime
		) ENGINE = MergeTree()
		ORDER BY (timestamp, student_id)
	`)
	if err != nil {
		return fmt.Errorf("création table answer_events échouée: %w", err)
	}

	log.Println("📊 Table answer_events prête")
	return nil
}

// RecordAnswer enregistre une réponse dans ClickHouse pour les stats.
//
// 🧠 STRUCTURE DE LA TABLE :
//    CREATE TABLE answer_events (
//        student_id  String,
//        question_id String,
//        was_correct Bool,
//        timestamp   DateTime
//    ) ENGINE = MergeTree()
//    ORDER BY (timestamp, student_id)
//
//    MergeTree est le moteur de stockage par défaut de ClickHouse.
//    Il trie les données par (timestamp, student_id) pour
//    accélérer les requêtes qui filtrent par date ou étudiant.
func (a *Analytics) RecordAnswer(studentID, questionID string, wasCorrect bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// INSERT INTO answer_events VALUES (...)
	// On insère une ligne à la fois (pour les événements en temps réel)
	return a.conn.AsyncInsert(ctx,
		`INSERT INTO answer_events (student_id, question_id, was_correct, timestamp) VALUES (?, ?, ?, now())`,
		false, // wait_async_insert = false → ne pas attendre la confirmation
		studentID, questionID, wasCorrect,
	)
}

// GetStudentStats retourne les stats d'un étudiant.
//
// 📝 EXEMPLE DE REQUÊTE SQL GÉNÉRÉE :
//    SELECT
//        count() AS total_answers,
//        countIf(was_correct = 1) AS correct_answers,
//        round(countIf(was_correct = 1) / count() * 100, 1) AS success_rate
//    FROM answer_events
//    WHERE student_id = 'student123'
//
//    countIf() est une fonction ClickHouse qui compte
//    seulement les lignes qui satisfont la condition.
//    C'est l'équivalent de SUM(CASE WHEN ...) en SQL standard.
type StudentStats struct {
	TotalAnswers   int     `json:"total_answers"`
	CorrectAnswers int     `json:"correct_answers"`
	SuccessRate    float64 `json:"success_rate"`
}

func (a *Analytics) GetStudentStats(studentID string) (*StudentStats, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var stats StudentStats
	err := a.conn.QueryRow(ctx, `
		SELECT
			count() AS total_answers,
			countIf(was_correct = 1) AS correct_answers,
			round(countIf(was_correct = 1) / count() * 100, 1) AS success_rate
		FROM answer_events
		WHERE student_id = $1
	`, studentID).Scan(&stats.TotalAnswers, &stats.CorrectAnswers, &stats.SuccessRate)

	if err != nil {
		return nil, fmt.Errorf("requête stats échouée: %w", err)
	}

	return &stats, nil
}

// GetDailyStats retourne les stats quotidiennes pour un étudiant.
// Utile pour le graphique "progression jour par jour" sur le dashboard.
//
// 📝 EXEMPLE DE REQUÊTE :
//    SELECT
//        toDate(timestamp) AS day,
//        countIf(was_correct = 1) AS correct,
//        count() AS total
//    FROM answer_events
//    WHERE student_id = 'student123'
//      AND timestamp >= now() - INTERVAL 30 DAY
//    GROUP BY day
//    ORDER BY day
type DailyStat struct {
	Day     string `json:"day"`     // "2024-01-15"
	Correct int    `json:"correct"` // Réponses correctes ce jour-là
	Total   int    `json:"total"`   // Total réponses ce jour-là
}

func (a *Analytics) GetDailyStats(studentID string, days int) ([]DailyStat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := a.conn.Query(ctx, `
		SELECT
			toDate(timestamp) AS day,
			countIf(was_correct = 1) AS correct,
			count() AS total
		FROM answer_events
		WHERE student_id = $1
		  AND timestamp >= now() - INTERVAL $2 DAY
		GROUP BY day
		ORDER BY day
	`, studentID, days)
	if err != nil {
		return nil, fmt.Errorf("requête daily stats échouée: %w", err)
	}
	defer rows.Close()

	var stats []DailyStat
	for rows.Next() {
		var s DailyStat
		var t time.Time
		if err := rows.Scan(&t, &s.Correct, &s.Total); err != nil {
			return nil, fmt.Errorf("scan daily stat: %w", err)
		}
		s.Day = t.Format("2006-01-02")
		stats = append(stats, s)
	}

	return stats, nil
}

// Close ferme la connexion ClickHouse.
func (a *Analytics) Close() error {
	return a.conn.Close()
}

// ============================================================
//  🧪 TESTER CLICKHOUSE
// ============================================================
// 1. Démarre ClickHouse : docker compose up clickhouse -d
// 2. Connecte-toi : docker compose exec clickhouse clickhouse-client
// 3. Crée la table :
//    CREATE TABLE answer_events (
//        student_id  String,
//        question_id String,
//        was_correct Bool,
//        timestamp   DateTime
//    ) ENGINE = MergeTree()
//    ORDER BY (timestamp, student_id)
//
// 4. Insère un test :
//    INSERT INTO answer_events VALUES ('student123', 'q1', true, now())
//
// 5. Vérifie :
//    SELECT * FROM answer_events
// ============================================================
