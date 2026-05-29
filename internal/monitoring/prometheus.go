package monitoring

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// ============================================================
//  📈 PROMETHEUS — Métriques et monitoring
// ============================================================
// 📖 DDIA Chapitre 1 : "Reliability" (Monitoring)
//
// ❓ C'EST QUOI PROMETHEUS ?
//    Prometheus est un "scraper" de métriques.
//    Toutes les 15 secondes, il va "gratter" (scrape) les métriques
//    sur chaque service pour voir comment ils se portent.
//
//    Imagine un médecin qui prend tes constantes :
//    - Pouls = nombre de requêtes par seconde
//    - Tension = temps de réponse moyen
//    - Température = taux d'erreur
//
//    Si une constante est anormale, Prometheus alerte.
//
// 🧠 MÉTRIQUES IMPORTANTES (RED Method) :
//    Rate    = nombre de requêtes par seconde
//    Errors  = nombre d'erreurs par seconde
//    Duration = temps de réponse (moyen, médian, 99e percentile)
//
//    Ces 3 métriques suffisent pour savoir si un service va bien.
//
// 🔗 LIENS UTILES :
//    - https://prometheus.io/docs/
//    - https://grafana.com/oss/grafana/
//    - github.com/prometheus/client_golang
// ============================================================

// Metrics regroupe toutes les métriques Prometheus de l'application.
// C'est le "tableau de bord" qui contient tous les indicateurs.
type Metrics struct {
	// Requêtes HTTP
	// Compteur : ne fait qu'augmenter (total de requêtes depuis le démarrage)
	httpRequestsTotal *prometheus.CounterVec

	// Temps de réponse
	// Histogramme : distribue les durées dans des "seaux" (buckets)
	// Exemple : [0.01s, 0.05s, 0.1s, 0.5s, 1s, 5s]
	// Permet de calculer la médiane, le 99e percentile, etc.
	httpRequestDuration *prometheus.HistogramVec

	// Requêtes en cours
	// Jauge (gauge) : peut monter et descendre
	// Utile pour détecter les goulots d'étranglement
	httpRequestsInFlight prometheus.Gauge

	// Réponses aux questions
	questionsAnswered *prometheus.CounterVec

	// Étudiants actifs
	activeStudents prometheus.Gauge
}

// NewMetrics crée et enregistre toutes les métriques Prometheus.
func NewMetrics() *Metrics {
	m := &Metrics{
		httpRequestsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "aeroforge_http_requests_total",
				Help: "Nombre total de requêtes HTTP",
			},
			[]string{"method", "path", "status"}, // Labels : GET /api/questions 200
		),

		httpRequestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "aeroforge_http_request_duration_seconds",
				Help:    "Durée des requêtes HTTP en secondes",
				Buckets: prometheus.DefBuckets, // [0.005, 0.01, 0.025, 0.05, ..., 10]
			},
			[]string{"method", "path"},
		),

		httpRequestsInFlight: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "aeroforge_http_requests_in_flight",
				Help: "Nombre de requêtes HTTP en cours",
			},
		),

		questionsAnswered: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "aeroforge_questions_answered_total",
				Help: "Nombre total de réponses aux questions",
			},
			[]string{"correct"}, // "true" ou "false"
		),

		activeStudents: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "aeroforge_active_students",
				Help: "Nombre d'étudiants actifs (connectés dans les 5 min)",
			},
		),
	}

	// Enregistrer les métriques dans le registre global Prometheus
	prometheus.MustRegister(
		m.httpRequestsTotal,
		m.httpRequestDuration,
		m.httpRequestsInFlight,
		m.questionsAnswered,
		m.activeStudents,
	)

	return m
}

// RecordRequest enregistre une requête HTTP.
// À appeler à la fin de chaque handler HTTP.
//
// 📝 EXEMPLE D'UTILISATION :
//    start := time.Now()
//    // ... traitement de la requête ...
//    metrics.RecordRequest("GET", "/api/questions", "200", time.Since(start))
func (m *Metrics) RecordRequest(method, path, status string, duration time.Duration) {
	m.httpRequestsTotal.WithLabelValues(method, path, status).Inc()
	m.httpRequestDuration.WithLabelValues(method, path).Observe(duration.Seconds())
}

// RecordAnswer enregistre une réponse à une question.
func (m *Metrics) RecordAnswer(correct bool) {
	label := "false"
	if correct {
		label = "true"
	}
	m.questionsAnswered.WithLabelValues(label).Inc()
}

// SetActiveStudents met à jour le nombre d'étudiants actifs.
func (m *Metrics) SetActiveStudents(count int) {
	m.activeStudents.Set(float64(count))
}

// Handler retourne le handler HTTP pour que Prometheus scrape les métriques.
// Prometheus viendra "gratter" les métriques sur cette URL.
//
// 📝 CONFIGURATION PROMETHEUS (prometheus.yml) :
//    scrape_configs:
//      - job_name: 'aeroforge'
//        scrape_interval: 15s
//        static_configs:
//          - targets: ['localhost:8080']
func (m *Metrics) Handler() http.Handler {
	return promhttp.Handler()
}

// StartMetricsServer démarre un serveur HTTP séparé pour les métriques.
// On utilise un port différent (2112) pour ne pas mélanger
// les métriques avec l'API principale.
//
// 📝 POURQUOI UN PORT SÉPARÉ ?
//    - L'API publique est sur le port 8080
//    - Les métriques sont sur le port 2112
//    - Prometheus scrape le port 2112
//    - Les utilisateurs ne voient pas les métriques
func StartMetricsServer(port string) *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	go func() {
		log.Printf("📈 Serveur métriques Prometheus sur :%s/metrics", port)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("❌ Erreur serveur métriques: %v", err)
		}
	}()

	return server
}

// ============================================================
//  🧪 TESTER PROMETHEUS
// ============================================================
// 1. Démarre Prometheus : docker compose up prometheus -d
// 2. Ouvre Prometheus : http://localhost:9090
// 3. Lance l'API : go run cmd/api-gateway/main.go
// 4. Vérifie les métriques : curl http://localhost:2112/metrics
// 5. Dans Prometheus, cherche "aeroforge_http_requests_total"
// 6. Tu devrais voir les métriques apparaître
// ============================================================
