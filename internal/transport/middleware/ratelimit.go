package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// 📖 DDIA Chapitre 11 : "Stream Processing"
//    Le rate limiting est un "backpressure" pattern :
//    on limite le débit pour protéger le système contre
//    les surges de trafic (DoS, bug client, etc.).
//
//    Implémentation : Token Bucket algorithm
//    - Chaque IP a un bucket de N tokens
//    - Chaque requête consomme 1 token
//    - Les tokens se régénèrent au rythme de R par seconde
//    - Si le bucket est vide, la requête est rejetée (429)

type visitor struct {
	tokens    float64
	lastCheck time.Time
}

type RateLimiter struct {
	mu       sync.RWMutex
	visitors map[string]*visitor
	rate     float64   // tokens par seconde
	burst    int       // capacité max du bucket
	cleanup  time.Duration // intervalle de nettoyage
}

// NewRateLimiter crée un nouveau rate limiter.
//
//	rate  : nombre de requêtes autorisées par seconde
//	burst : nombre max de requêtes en rafale
func NewRateLimiter(rate float64, burst int) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*visitor),
		rate:     rate,
		burst:    burst,
		cleanup:  time.Minute,
	}

	// Nettoyage périodique des visiteurs inactifs
	go func() {
		ticker := time.NewTicker(rl.cleanup)
		defer ticker.Stop()
		for range ticker.C {
			rl.mu.Lock()
			for ip, v := range rl.visitors {
				if time.Since(v.lastCheck) > rl.cleanup*2 {
					delete(rl.visitors, ip)
				}
			}
			rl.mu.Unlock()
		}
	}()

	return rl
}

func (rl *RateLimiter) allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	v, exists := rl.visitors[ip]
	now := time.Now()

	if !exists {
		rl.visitors[ip] = &visitor{
			tokens:    float64(rl.burst) - 1,
			lastCheck: now,
		}
		return true
	}

	// Recharger les tokens
	elapsed := now.Sub(v.lastCheck).Seconds()
	v.tokens += elapsed * rl.rate
	if v.tokens > float64(rl.burst) {
		v.tokens = float64(rl.burst)
	}
	v.lastCheck = now

	// Consommer un token
	if v.tokens >= 1 {
		v.tokens--
		return true
	}

	return false
}

// RateLimit retourne un middleware Gin qui limite le débit par IP.
//
//	rate  : requêtes par seconde
//	burst : rafale maximale
//
// Usage:
//
//	r := gin.Default()
//	r.Use(middleware.RateLimit(10, 20)) // 10 req/s, burst 20
func RateLimit(rate float64, burst int) gin.HandlerFunc {
	limiter := NewRateLimiter(rate, burst)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		if !limiter.allow(ip) {
			c.Header("Retry-After", "1")
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":       "trop de requêtes, veuillez réessayer",
				"retry_after": "1 seconde",
			})
			return
		}
		c.Next()
	}
}

// RateLimitStrict est un rate limiter strict pour les routes sensibles
// (auth, création de ressources). Limite à 5 req/s avec burst de 10.
func RateLimitStrict() gin.HandlerFunc {
	return RateLimit(5, 10)
}

// RateLimitDefault est le rate limiter par défaut pour l'API.
// Limite à 30 req/s avec burst de 50.
func RateLimitDefault() gin.HandlerFunc {
	return RateLimit(30, 50)
}
