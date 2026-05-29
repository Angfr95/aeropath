package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// ============================================================
//  🗄️ CACHE REDIS — Accélérer les accès fréquents
// ============================================================
// 📖 DDIA Chapitre 12 : "The Future of Data Systems" (Caching)
//
// ❓ POURQUOI UN CACHE ?
//    Imagine que 1000 étudiants demandent la même leçon en même temps.
//    Sans cache : 1000 requêtes PostgreSQL → la DB ralentit (ou crashe).
//    Avec cache : 1 requête PostgreSQL, 999 réponses depuis Redis → rapide !
//
//    Redis est une base de données "en mémoire" (RAM).
//    C'est 100x plus rapide que PostgreSQL car :
//    - PostgreSQL lit sur le disque dur (SSD) → ~0.1ms
//    - Redis lit dans la RAM → ~0.001ms
//
// 🧠 QUAND UTILISER LE CACHE ?
//    ✅ Données qui changent rarement (leçons, questions)
//    ✅ Données lues souvent (profil étudiant)
//    ✅ Résultats de calculs coûteux (recommandations)
//    ❌ Données qui changent tout le temps (scores en direct)
//    ❌ Données critiques (argent, sécurité)
//
// 🔗 LIENS UTILES :
//    - https://redis.io/docs/
//    - github.com/redis/go-redis
// ============================================================

// Cache gère le cache Redis.
// C'est comme un "post-it" géant : on écrit les infos importantes
// dessus pour les retrouver vite, mais on sait que ça peut
// disparaître (expiration).
type Cache struct {
	client *redis.Client
}

// NewCache crée une nouvelle connexion Redis.
//
// 🛡️ PRODUCTION :
//    - Pool de connexions : jusqu'à 10 connexions simultanées
//    - Timeout de connexion : 5 secondes max
//    - Timeout de lecture/écriture : 3 secondes max
//    - Retry : 3 tentatives si Redis ne répond pas
//    - Ping de vérification au démarrage
//
// 📝 EXEMPLE D'UTILISATION :
//    cache, err := redis.NewCache("localhost:6379", "")
//    if err != nil { log.Fatal(err) }
//    defer cache.Close()
//
//    // Mettre en cache
//    err = cache.Set("lesson:123", maLecon, 1*time.Hour)
//
//    // Lire depuis le cache
//    var lecon domain.Lesson
//    err = cache.Get("lesson:123", &lecon)
func NewCache(addr, password string) (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         addr,            // "localhost:6379"
		Password:     password,        // "" si pas de mot de passe
		DB:           0,               // Base de données Redis #0 (par défaut)
		PoolSize:     10,              // Max 10 connexions simultanées
		MinIdleConns: 3,               // Garder 3 connexions toujours prêtes
		DialTimeout:  5 * time.Second, // Timeout pour établir la connexion
		ReadTimeout:  3 * time.Second, // Timeout pour lire une réponse
		WriteTimeout: 3 * time.Second, // Timeout pour envoyer une commande
		MaxRetries:   3,               // Réessayer 3 fois en cas d'erreur réseau
	})

	// Ping pour vérifier que Redis répond
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("connexion Redis échouée: %w", err)
	}

	log.Println("🗄️ Cache Redis connecté")
	return &Cache{client: client}, nil
}

// Set met une valeur en cache avec une durée d'expiration.
//
// 🧠 POURQUOI UNE EXPIRATION (TTL) ?
//    Sans expiration, le cache grossirait indéfiniment → plus de RAM.
//    Avec expiration, les données "vieilles" sont automatiquement supprimées.
//
//    Durées typiques :
//    - Leçons : 1 heure (changent rarement)
//    - Questions : 30 minutes (peuvent être modifiées)
//    - Profil étudiant : 5 minutes (peut changer)
//    - Recommandations : 10 minutes (recalcul périodique)
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// On convertit la valeur en JSON pour la stocker dans Redis
	// Redis ne stocke que des strings, donc on sérialise
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("sérialisation échouée: %w", err)
	}

	// SET key value EX seconds
	// Exemple : SET lesson:123 "{...}" EX 3600
	return c.client.Set(ctx, key, data, ttl).Err()
}

// Get lit une valeur depuis le cache.
// Si la clé n'existe pas, retourne redis.Nil.
//
// 📝 EXEMPLE :
//    var lecon domain.Lesson
//    err := cache.Get("lesson:123", &lecon)
//    if errors.Is(err, redis.Nil) {
//        // Pas dans le cache → aller chercher dans PostgreSQL
//    }
func (c *Cache) Get(key string, dest interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// On récupère le JSON depuis Redis
	data, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return err // Peut être redis.Nil si la clé n'existe pas
	}

	// On reconvertit le JSON en structure Go
	return json.Unmarshal(data, dest)
}

// Delete supprime une clé du cache.
// Utile quand une donnée est modifiée → on invalide le cache
// pour forcer un rechargement depuis PostgreSQL.
//
// 📝 EXEMPLE :
//    // Quand on modifie une leçon, on supprime son cache
//    cache.Delete("lesson:123")
//    // La prochaine lecture rechargera depuis PostgreSQL
func (c *Cache) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.client.Del(ctx, key).Err()
}

// Exists vérifie si une clé existe dans le cache.
// Plus rapide que Get car on ne transfère pas les données.
func (c *Cache) Exists(key string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	n, err := c.client.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("vérification existence échouée: %w", err)
	}
	return n > 0, nil
}

// IsMiss vérifie si une erreur est un "cache miss" (clé inexistante).
// Utile pour les if/else propres.
//
// 📝 EXEMPLE :
//    var lecon domain.Lesson
//    err := cache.Get("lesson:123", &lecon)
//    if cache.IsMiss(err) {
//        // Pas dans le cache → charger depuis PostgreSQL
//    }
func (c *Cache) IsMiss(err error) bool {
	return errors.Is(err, redis.Nil)
}

// Close ferme la connexion Redis.
func (c *Cache) Close() error {
	return c.client.Close()
}

// ============================================================
//  🧪 TESTER LE CACHE
// ============================================================
// 1. Démarre Redis : docker compose up redis -d
// 2. Teste avec redis-cli :
//    redis-cli ping  →  PONG
//    redis-cli set "hello" "world" EX 60  →  OK
//    redis-cli get "hello"  →  "world"
//    redis-cli ttl "hello"  →  58 (secondes restantes)
// ============================================================
