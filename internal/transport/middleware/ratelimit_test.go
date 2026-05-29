package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRateLimiterAllow(t *testing.T) {
	rl := NewRateLimiter(10, 5)

	// Les premières requêtes doivent passer (burst)
	for i := 0; i < 5; i++ {
		if !rl.allow("test-ip") {
			t.Fatalf("requête %d devrait être autorisée (burst)", i+1)
		}
	}

	// La 6ème devrait être refusée (burst épuisé)
	if rl.allow("test-ip") {
		t.Fatal("la 6ème requête devrait être refusée")
	}
}

func TestRateLimiterDifferentIPs(t *testing.T) {
	rl := NewRateLimiter(10, 3)

	// Chaque IP a son propre bucket
	for i := 0; i < 3; i++ {
		if !rl.allow("ip-a") {
			t.Fatalf("ip-a requête %d devrait passer", i+1)
		}
		if !rl.allow("ip-b") {
			t.Fatalf("ip-b requête %d devrait passer", i+1)
		}
	}

	// Les deux IPs devraient être bloquées
	if rl.allow("ip-a") {
		t.Fatal("ip-a devrait être bloquée")
	}
	if rl.allow("ip-b") {
		t.Fatal("ip-b devrait être bloquée")
	}
}

func TestRateLimiterTokenRefill(t *testing.T) {
	rl := NewRateLimiter(100, 1) // 100 tokens/s, burst 1

	// Consommer l'unique token
	if !rl.allow("test-ip") {
		t.Fatal("première requête devrait passer")
	}

	// La seconde devrait être refusée immédiatement
	if rl.allow("test-ip") {
		t.Fatal("seconde requête devrait être refusée")
	}
}

func TestRateLimitMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Créer un routeur avec rate limiting strict
	r := gin.New()
	r.Use(RateLimit(100, 5)) // 100 req/s, burst 5

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Les 5 premières requêtes doivent passer
	for i := 0; i < 5; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.RemoteAddr = "192.168.1.1:12345"
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("requête %d devrait être 200, got %d", i+1, w.Code)
		}
	}

	// La 6ème devrait être 429
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.RemoteAddr = "192.168.1.1:12345"
	r.ServeHTTP(w, req)

	if w.Code != http.StatusTooManyRequests {
		t.Fatalf("devrait être 429, got %d", w.Code)
	}
}

func TestRateLimitMiddlewareDifferentIPs(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.Use(RateLimit(100, 2)) // 100 req/s, burst 2

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// IP A : 2 requêtes OK
	for i := 0; i < 2; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.RemoteAddr = "10.0.0.1:12345"
		r.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Fatalf("IP A requête %d devrait être 200", i+1)
		}
	}

	// IP B : 2 requêtes OK (bucket indépendant)
	for i := 0; i < 2; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.RemoteAddr = "10.0.0.2:12345"
		r.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Fatalf("IP B requête %d devrait être 200", i+1)
		}
	}

	// IP A : 3ème requête bloquée
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.RemoteAddr = "10.0.0.1:12345"
	r.ServeHTTP(w, req)
	if w.Code != http.StatusTooManyRequests {
		t.Fatalf("IP A 3ème requête devrait être 429, got %d", w.Code)
	}
}

func TestRateLimitStrict(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.Use(RateLimitStrict()) // 5 req/s, burst 10

	r.GET("/auth", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Les 10 premières requêtes doivent passer (burst)
	for i := 0; i < 10; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/auth", nil)
		req.RemoteAddr = "10.0.0.1:12345"
		r.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Fatalf("requête %d devrait être 200, got %d", i+1, w.Code)
		}
	}

	// La 11ème devrait être 429
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/auth", nil)
	req.RemoteAddr = "10.0.0.1:12345"
	r.ServeHTTP(w, req)
	if w.Code != http.StatusTooManyRequests {
		t.Fatalf("devrait être 429, got %d", w.Code)
	}
}

func TestRateLimitDefault(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.Use(RateLimitDefault()) // 30 req/s, burst 50

	r.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 50 requêtes doivent passer (burst)
	for i := 0; i < 50; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api", nil)
		req.RemoteAddr = "10.0.0.1:12345"
		r.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Fatalf("requête %d devrait être 200, got %d", i+1, w.Code)
		}
	}

	// La 51ème devrait être 429
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api", nil)
	req.RemoteAddr = "10.0.0.1:12345"
	r.ServeHTTP(w, req)
	if w.Code != http.StatusTooManyRequests {
		t.Fatalf("devrait être 429, got %d", w.Code)
	}
}
