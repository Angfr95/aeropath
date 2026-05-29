package monitoring

import (
	"fmt"
	"log"
	"time"
)

// ============================================================
//  🚨 ALERTING — Être prévenu quand quelque chose ne va pas
// ============================================================
// 📖 DDIA Chapitre 12 : "The Future of Data Systems" (SLOs)
//
// ❓ C'EST QUOI UNE ALERTE ?
//    Une alerte, c'est quand Prometheus te dit :
//    "Hé, il y a un problème, regarde ça !"
//
//    Exemples d'alertes :
//    - "Le taux d'erreur dépasse 5% depuis 5 minutes"
//    - "Le temps de réponse moyen est > 2 secondes"
//    - "Le service est down (pas de réponse depuis 1 minute)"
//
// 🧠 CONCEPTS IMPORTANTS :
//    SLO (Service Level Objective) = objectif de qualité
//    Exemple : "99.9% des requêtes doivent répondre en moins de 1 seconde"
//
//    SLI (Service Level Indicator) = mesure réelle
//    Exemple : "99.5% des requêtes répondent en moins de 1 seconde"
//
//    Error Budget = temps d'erreur autorisé
//    Si SLO = 99.9% sur 30 jours, l'error budget = 43 minutes
//    Tant qu'on est dans l'error budget, pas d'inquiétude
//
// 🔗 LIENS UTILES :
//    - https://prometheus.io/docs/alerting/latest/alertmanager/
//    - https://grafana.com/oss/alertmanager/
// ============================================================

// AlertRule définit une règle d'alerte Prometheus.
// C'est la "recette" qui dit : "si X arrive, préviens-moi".
type AlertRule struct {
	Name        string        // Nom de l'alerte (ex: "HighErrorRate")
	Description string        // Description lisible (ex: "Le taux d'erreur est élevé")
	Expr        string        // Expression PromQL (ex: "rate(errors_total[5m]) > 0.05")
	For         time.Duration // Durée avant de déclencher (ex: 5 minutes)
	Severity    string        // "critical", "warning", "info"
}

// AlertingConfig regroupe toutes les règles d'alerte.
// C'est le "cahier" qui contient toutes les recettes d'alertes.
type AlertingConfig struct {
	Rules []AlertRule
}

// DefaultAlertingConfig retourne les alertes par défaut pour AeroForge.
//
// 🧠 POURQUOI CES ALERTES ?
//    - HighErrorRate : si >5% des requêtes échouent, il y a un bug
//    - HighLatency : si les réponses sont lentes, la DB est peut-être saturée
//    - ServiceDown : si le service ne répond plus, il faut redémarrer
//    - LowSuccessRate : si les étudiants ratent plus que d'habitude,
//      peut-être que les questions sont trop difficiles ou mal formulées
func DefaultAlertingConfig() *AlertingConfig {
	return &AlertingConfig{
		Rules: []AlertRule{
			{
				Name:        "HighErrorRate",
				Description: "Le taux d'erreur HTTP dépasse 5%",
				Expr:        "rate(aeroforge_http_requests_total{status=~\"5..\"}[5m]) / rate(aeroforge_http_requests_total[5m]) > 0.05",
				For:         5 * time.Minute,
				Severity:    "critical",
			},
			{
				Name:        "HighLatency",
				Description: "Le temps de réponse moyen dépasse 2 secondes",
				Expr:        "rate(aeroforge_http_request_duration_seconds_sum[5m]) / rate(aeroforge_http_request_duration_seconds_count[5m]) > 2",
				For:         5 * time.Minute,
				Severity:    "warning",
			},
			{
				Name:        "ServiceDown",
				Description: "Le service API ne répond plus",
				Expr:        "up{job=\"aeroforge\"} == 0",
				For:         1 * time.Minute,
				Severity:    "critical",
			},
			{
				Name:        "LowSuccessRate",
				Description: "Le taux de réussite des étudiants est < 40%",
				Expr:        "rate(aeroforge_questions_answered_total{correct=\"true\"}[1h]) / rate(aeroforge_questions_answered_total[1h]) < 0.4",
				For:         30 * time.Minute,
				Severity:    "warning",
			},
		},
	}
}

// GeneratePrometheusRules génère le fichier YAML des règles Prometheus.
// Ce fichier est chargé par Prometheus au démarrage.
//
// 📝 EXEMPLE DE SORTIE :
//    groups:
//      - name: aeroforge
//        rules:
//          - alert: HighErrorRate
//            expr: rate(aeroforge_http_requests_total{status=~"5.."}[5m]) / rate(aeroforge_http_requests_total[5m]) > 0.05
//            for: 5m
//            labels:
//              severity: critical
//            annotations:
//              description: "Le taux d'erreur HTTP dépasse 5%"
func (c *AlertingConfig) GeneratePrometheusRules() string {
	yaml := "groups:\n  - name: aeroforge\n    rules:\n"

	for _, rule := range c.Rules {
		yaml += fmt.Sprintf(`      - alert: %s
        expr: %s
        for: %s
        labels:
          severity: %s
        annotations:
          description: "%s"
`, rule.Name, rule.Expr, rule.For.String(), rule.Severity, rule.Description)
	}

	return yaml
}

// PrintRules affiche les règles d'alerte dans la console.
// Utile pour le débogage.
func (c *AlertingConfig) PrintRules() {
	log.Println("🚨 Règles d'alerte configurées:")
	for _, rule := range c.Rules {
		log.Printf("  - %s [%s]: %s", rule.Name, rule.Severity, rule.Description)
	}
}

// ============================================================
//  🧪 TESTER LES ALERTES
// ============================================================
// 1. Démarre Prometheus + AlertManager : docker compose up -d
// 2. Ouvre Prometheus : http://localhost:9090/alerts
// 3. Tu devrais voir les alertes "pending" (en attente)
// 4. Si une condition est remplie pendant assez longtemps,
//    l'alerte passe en "firing" (déclenchée)
// 5. AlertManager envoie une notification (email, Slack, etc.)
// ============================================================
