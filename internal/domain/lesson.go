package domain

import "time"

// License représente le type de licence aéronautique.
//
// 📖 DDIA Chapitre 4 : "Encoding and Evolution"
//    Les licences sont un "enum" en Go, mais stockées comme des strings
//    dans PostgreSQL. Pourquoi pas un ENUM SQL ? Parce que les ENUMs SQL
//    sont rigides : ajouter une licence nécessite ALTER TYPE.
//    Avec des strings, on peut ajouter "BPL" sans migration.
//    C'est le principe "schemaless" du Chapitre 2.
type License string

const (
	LicensePPL  License = "PPL"
	LicenseATPL License = "ATPL"
	LicenseLAPL License = "LAPL"
	LicenseCPL  License = "CPL"
	LicenseIR   License = "IR"
)

// Category représente la matière d'examen.
//
// 📖 DDIA Chapitre 3 : "Storage and Retrieval"
//    Les catégories sont utilisées comme clé de partitionnement.
//    Quand on cherche "toutes les questions de météo",
//    PostgreSQL utilise l'index sur (category) pour éviter
//    de scanner toute la table (full table scan).
type Category string

const (
	CategoryAirLaw           Category = "airlaw"
	CategoryMeteorology      Category = "meteorology"
	CategoryNavigation       Category = "navigation"
	CategoryPerformance      Category = "performance"
	CategoryAircraftGeneral  Category = "aircraft_general"
	CategoryFlightPlanning   Category = "flight_planning"
	CategoryHumanPerformance Category = "human_performance"
	CategoryOperationalProcs Category = "operational_procedures"
	CategoryPrinciplesFlight Category = "principles_of_flight"
	CategoryCommunications   Category = "communications"
	CategoryMassAndBalance   Category = "mass_and_balance"
	CategoryInstrumentation  Category = "instrumentation"
)

// Lesson représente une leçon pédagogique enrichie.
type Lesson struct {
	ID               string            `json:"id"`
	License          License           `json:"license"`
	Category         Category          `json:"category"`
	Theme            string            `json:"theme"`
	TitleFr          string            `json:"title_fr"`
	TitleEn          string            `json:"title_en"`
	ContentFr        string            `json:"content_fr"`
	ContentEn        string            `json:"content_en"`
	Level            int               `json:"level"`                        // 1=Basic, 2=Intermediate, 3=Advanced
	Difficulty       int               `json:"difficulty"`
	OrderIndex       int               `json:"order_index"`
	DurationMinutes  int               `json:"duration_minutes"`             // Durée estimée de la leçon
	Tags             []string          `json:"tags,omitempty"`               // Mots-clés
	LearningObjectives []string        `json:"learning_objectives,omitempty"` // Objectifs pédagogiques
	CreatedAt        time.Time         `json:"created_at"`
}


// LessonRepository définit le contrat pour accéder aux leçons.
type LessonRepository interface {
	Create(l *Lesson) error
	Update(l *Lesson) error
	FindByID(id string) (*Lesson, error)
	FindByTheme(theme string) ([]*Lesson, error)
	FindAll() ([]*Lesson, error)
	FindAllPaginated(limit, offset int) ([]*Lesson, error)
	FindByLicense(license License) ([]*Lesson, error)
	FindByCategory(category Category) ([]*Lesson, error)
	FindByLicenseAndCategory(license License, category Category) ([]*Lesson, error)
	FindByDifficulty(difficulty int) ([]*Lesson, error)
	Delete(id string) error
	Count() (int, error)
	CountByLicense(license License) (int, error)
	CountByCategory(category Category) (int, error)
}
