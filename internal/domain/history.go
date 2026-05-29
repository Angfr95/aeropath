package domain

import "time"

// QuestionHistory représente l'historique des réponses d'un étudiant.
//
// 📖 DDIA Chapitre 11 : "Stream Processing"
//    Chaque réponse est un "événement" dans le système.
//    On pourrait stocker ces événements dans un "event log" (Kafka/NATS)
//    et les rejouer pour recalculer les stats en cas de bug.
//    C'est le principe de l'Event Sourcing.
//
// 📖 DDIA Chapitre 12 : "The Future of Data Systems"
//    L'historique est utilisé par le moteur de recommandation pour :
//    - Calculer le score de maîtrise (combien de bonnes réponses ?)
//    - Déterminer la répétition espacée (quand revoir cette question ?)
//    - Identifier les sujets faibles (où l'étudiant fait le plus d'erreurs ?)
//    C'est un exemple de "derived data" : on part de données brutes
//    (les réponses) pour créer des données dérivées (les recommandations).
type QuestionHistory struct {
	StudentID  string    `json:"student_id"`
	QuestionID string    `json:"question_id"`
	SeenAt     time.Time `json:"seen_at"`
	WasCorrect bool      `json:"was_correct"`
}

// StudentStats contient les statistiques agrégées d'un étudiant.
//
// 📖 DDIA Chapitre 3 : "Storage and Retrieval"
//    Ces stats sont calculées à partir de l'historique.
//    C'est un exemple de "materialized view" : on pré-calcule
//    les résultats pour éviter de les recalculer à chaque requête.
//    Sans ça, chaque page de dashboard ferait un COUNT(*) sur
//    des milliers de lignes.
type StudentStats struct {
	TotalQuestions int     `json:"total_questions"`
	CorrectAnswers int     `json:"correct_answers"`
	WrongAnswers   int     `json:"wrong_answers"`
	SuccessRate    float64 `json:"success_rate"`
	CurrentLevel   int     `json:"current_level"` // 1-5, basé sur la moyenne des difficultés réussies
}

// QuestionHistoryRepository définit le contrat pour l'historique.
//
// 📖 DDIA Chapitre 5 : "Replication"
//    L'historique est écrit en PostgreSQL (garantie ACID).
//    Mais on pourrait aussi l'écrire dans ClickHouse pour les analytics,
//    et garder PostgreSQL seulement pour les données transactionnelles.
//    C'est le pattern CQRS (Command Query Responsibility Segregation).
type QuestionHistoryRepository interface {
	RecordAnswer(studentID, questionID string, wasCorrect bool) error
	GetHistory(studentID string) ([]*QuestionHistory, error)
	GetStats(studentID string) (*StudentStats, error)
	GetSeenQuestionIDs(studentID string) ([]string, error)
	Count() (int, error)
}
