package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"

	_ "github.com/lib/pq"
)

// DBRouter gère la connexion aux différentes bases de données
// selon la licence et la langue (architecture multi-BDD).
//
// Exemple d'URLs de connexion (dans .env):
//   DB_PPL_FR=postgres://user:pass@localhost:5432/aero_ppl_fr
//   DB_PPL_EN=postgres://user:pass@localhost:5432/aero_ppl_en
//   DB_PPL_ZH=postgres://user:pass@localhost:5432/aero_ppl_zh
//   DB_PPL_ES=postgres://user:pass@localhost:5432/aero_ppl_es
//   DB_ATPL_FR=postgres://user:pass@localhost:5432/aero_atpl_fr
//   DB_ATPL_EN=postgres://user:pass@localhost:5432/aero_atpl_en
//   DB_GAMIFICATION=postgres://user:pass@localhost:5432/aero_gamification
//   DB_CONFIG=postgres://user:pass@localhost:5432/aero_config
type DBRouter struct {
	mu sync.RWMutex

	// Content databases (licence + langue)
	PPL_FR  *sql.DB
	PPL_EN  *sql.DB
	PPL_ZH  *sql.DB
	PPL_ES  *sql.DB
	PPL_DE  *sql.DB
	ATPL_FR *sql.DB
	ATPL_EN *sql.DB

	// Shared databases
	Gamification *sql.DB
	Config       *sql.DB
}

// NewDBRouter crée un routeur et connecte toutes les BDD configurées.
// Les URLs sont lues depuis un map (peut venir de .env ou autre source).
func NewDBRouter(urls map[string]string) (*DBRouter, error) {
	r := &DBRouter{}

	// Map des connexions à établir
	connections := map[string]**sql.DB{
		"ppl_fr":  &r.PPL_FR,
		"ppl_en":  &r.PPL_EN,
		"ppl_zh":  &r.PPL_ZH,
		"ppl_es":  &r.PPL_ES,
		"ppl_de":  &r.PPL_DE,
		"atpl_fr": &r.ATPL_FR,
		"atpl_en": &r.ATPL_EN,
		"gamification": &r.Gamification,
		"config":       &r.Config,
	}

	for name, dbPtr := range connections {
		url, ok := urls[name]
		if !ok || url == "" {
			log.Printf("ℹ️  DB %s: non configurée, ignorée", name)
			continue
		}

		db, err := sql.Open("postgres", url)
		if err != nil {
			return nil, fmt.Errorf("open %s: %w", name, err)
		}

		// Config pool
		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(5)

		if err := db.Ping(); err != nil {
			return nil, fmt.Errorf("ping %s: %w", name, err)
		}

		*dbPtr = db
		log.Printf("✅ DB %s: connectée", name)
	}

	return r, nil
}

// GetDB retourne la BDD correspondant à la licence et la langue.
// Exemples: ("ppl", "fr") → PPL_FR, ("atpl", "en") → ATPL_EN
func (r *DBRouter) GetDB(license, lang string) (*sql.DB, error) {
	lang = strings.ToLower(lang)
	license = strings.ToLower(license)

	// "gamification" et "config" sont des mots-clés spéciaux
	if license == "gamification" {
		return r.getDB(r.Gamification, "gamification")
	}
	if license == "config" {
		return r.getDB(r.Config, "config")
	}

	key := fmt.Sprintf("%s_%s", license, lang)

	switch key {
	case "ppl_fr":
		return r.getDB(r.PPL_FR, "PPL_FR")
	case "ppl_en":
		return r.getDB(r.PPL_EN, "PPL_EN")
	case "ppl_zh":
		return r.getDB(r.PPL_ZH, "PPL_ZH")
	case "ppl_es":
		return r.getDB(r.PPL_ES, "PPL_ES")
	case "ppl_de":
		return r.getDB(r.PPL_DE, "PPL_DE")
	case "atpl_fr":
		return r.getDB(r.ATPL_FR, "ATPL_FR")
	case "atpl_en":
		return r.getDB(r.ATPL_EN, "ATPL_EN")
	default:
		return nil, fmt.Errorf("DB introuvable pour license=%s lang=%s", license, lang)
	}
}

func (r *DBRouter) getDB(db *sql.DB, name string) (*sql.DB, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if db == nil {
		return nil, fmt.Errorf("DB %s: non configurée", name)
	}
	return db, nil
}

// ListAvailableDBs retourne la liste des BDD connectées.
func (r *DBRouter) ListAvailableDBs() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var dbs []string
	add := func(name string, db *sql.DB) {
		if db != nil {
			dbs = append(dbs, name)
		}
	}

	add("PPL_FR", r.PPL_FR)
	add("PPL_EN", r.PPL_EN)
	add("PPL_ZH", r.PPL_ZH)
	add("PPL_ES", r.PPL_ES)
	add("PPL_DE", r.PPL_DE)
	add("ATPL_FR", r.ATPL_FR)
	add("ATPL_EN", r.ATPL_EN)
	add("gamification", r.Gamification)
	add("config", r.Config)

	return dbs
}

// Close ferme toutes les connexions.
func (r *DBRouter) Close() {
	r.mu.Lock()
	defer r.mu.Unlock()

	closeDB := func(db *sql.DB, name string) {
		if db != nil {
			db.Close()
			log.Printf("🔌 DB %s: déconnectée", name)
		}
	}

	closeDB(r.PPL_FR, "PPL_FR")
	closeDB(r.PPL_EN, "PPL_EN")
	closeDB(r.PPL_ZH, "PPL_ZH")
	closeDB(r.PPL_ES, "PPL_ES")
	closeDB(r.PPL_DE, "PPL_DE")
	closeDB(r.ATPL_FR, "ATPL_FR")
	closeDB(r.ATPL_EN, "ATPL_EN")
	closeDB(r.Gamification, "gamification")
	closeDB(r.Config, "config")
}

// NewDBRouterFromEnv crée un routeur à partir d'un map URL.
// Utilisation typique : charger depuis .env ou viper config.
func NewDBRouterFromEnv(getEnv func(string) string) (*DBRouter, error) {
	urls := map[string]string{
		"ppl_fr":  getEnv("DB_PPL_FR"),
		"ppl_en":  getEnv("DB_PPL_EN"),
		"ppl_zh":  getEnv("DB_PPL_ZH"),
		"ppl_es":  getEnv("DB_PPL_ES"),
		"ppl_de":  getEnv("DB_PPL_DE"),
		"atpl_fr": getEnv("DB_ATPL_FR"),
		"atpl_en": getEnv("DB_ATPL_EN"),
		"gamification": getEnv("DB_GAMIFICATION"),
		"config":       getEnv("DB_CONFIG"),
	}
	return NewDBRouter(urls)
}
