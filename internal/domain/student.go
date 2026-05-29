package domain

import "time"

// Student représente un étudiant inscrit sur AeroForge.
//
// 📖 DDIA Chapitre 5 : "Replication"
//    Les sessions utilisateur sont stockées dans Redis (pas ici).
//    Pourquoi ? Parce que les sessions changent à chaque requête.
//    Si on écrivait dans PostgreSQL à chaque requête, ce serait trop lent.
//    Redis est plus adapté car il est en RAM et supporte le TTL.
//
// 📖 DDIA Chapitre 11 : "Stream Processing"
//    Quand un étudiant s'inscrit, on publie un événement "student.registered"
//    sur NATS. Le service Analytics reçoit l'événement et crée un profil
//    dans ClickHouse. Le service Notification envoie un email de bienvenue.
type Student struct {
    ID           string    `json:"id"`
    Email        string    `json:"email"`
    PasswordHash string    `json:"-"`
    Lang         string    `json:"lang"`
    CreatedAt    time.Time `json:"created_at"`
}

// StudentRepository définit le contrat pour accéder aux étudiants.
//
// 📖 DDIA Chapitre 2 : "Data Models and Query Languages"
//    Même pattern que QuestionRepository : on cache le stockage.
//    Ici, on utilise PostgreSQL car on a besoin de ACID :
//    - Atomicité : la création d'un étudiant est tout ou rien
//    - Cohérence : l'email doit être unique
//    - Isolation : deux inscriptions simultanées ne s'écrasent pas
//    - Durabilité : une fois créé, l'étudiant existe même après un crash
type StudentRepository interface {
    Create(s *Student) error
    FindByEmail(email string) (*Student, error)
    FindByID(id string) (*Student, error)
    UpdateLang(id, lang string) error
    Delete(id string) error
    Count() (int, error)
}
