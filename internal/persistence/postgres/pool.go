package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPool crée un pool de connexions PostgreSQL.
//
// 📖 DDIA Chapitre 5 : "Replication"
//    Un pool de connexions est un groupe de connexions réutilisables.
//    Au lieu d'ouvrir/fermer une connexion à chaque requête (lent),
//    on garde un pool de connexions ouvertes en permanence.
//
//    C'est comme un parking de taxis :
//    - Sans pool : chaque client appelle un taxi → 5 min d'attente
//    - Avec pool : les taxis attendent déjà → 0 sec d'attente
//
// 🛡️ PRODUCTION :
//    - MaxConns : 25 connexions max (évite de saturer PostgreSQL)
//    - MinConns : 5 connexions toujours prêtes
//    - MaxConnLifetime : 30 min (renouvellement périodique)
//    - MaxConnIdleTime : 5 min (fermer les inactives)
//    - HealthCheck : ping toutes les 30s
//    - Timeout de connexion : 5s
func NewPool(databaseURL string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	config.MaxConns = 25
	config.MinConns = 5
	config.MaxConnLifetime = 30 * time.Minute
	config.MaxConnIdleTime = 5 * time.Minute
	config.HealthCheckPeriod = 30 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("create pool: %w", err)
	}

	// Vérifier que PostgreSQL répond
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping PostgreSQL: %w", err)
	}

	log.Println("🗄️ Pool PostgreSQL connecté (25 max, 5 min)")
	return pool, nil
}
