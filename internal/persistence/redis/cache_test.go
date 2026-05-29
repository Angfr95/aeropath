package redis

import (
	"os"
	"testing"
	"time"
)

// Test d'intégration pour le cache Redis.
//
// Prérequis : Redis doit tourner sur localhost:6379
// Lancement : docker compose up redis -d
//
// Ces tests sont désactivés par défaut car ils nécessitent Redis.
// Pour les lancer : go test -tags=integration ./internal/persistence/redis/
//
// Si Redis n'est pas disponible, les tests sont skip automatiquement.

func getRedisAddr() string {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	return addr
}

func skipIfRedisUnavailable(t *testing.T) *Cache {
	t.Helper()

	cache, err := NewCache(getRedisAddr(), "")
	if err != nil {
		t.Skipf("Redis non disponible (%v), skip test d'intégration", err)
	}
	return cache
}

func TestIntegrationCacheSetGet(t *testing.T) {
	cache := skipIfRedisUnavailable(t)
	defer cache.Close()

	key := "test:int:setget:" + time.Now().Format("150405.000")
	value := map[string]interface{}{
		"name":  "test-student",
		"score": 42,
	}

	err := cache.Set(key, value, 30*time.Second)
	if err != nil {
		t.Fatalf("Set a échoué: %v", err)
	}

	var result map[string]interface{}
	err = cache.Get(key, &result)
	if err != nil {
		t.Fatalf("Get a échoué: %v", err)
	}

	if result["name"] != "test-student" {
		t.Fatalf("name attendu 'test-student', got '%v'", result["name"])
	}
	if result["score"] != float64(42) { // JSON unmarshal convertit les nombres en float64
		t.Fatalf("score attendu 42, got '%v'", result["score"])
	}
}

func TestIntegrationCacheGetMiss(t *testing.T) {
	cache := skipIfRedisUnavailable(t)
	defer cache.Close()

	key := "test:int:miss:" + time.Now().Format("150405.000")

	var result map[string]interface{}
	err := cache.Get(key, &result)

	if !cache.IsMiss(err) {
		t.Fatalf("devrait retourner un cache miss pour une clé inexistante, got: %v", err)
	}
}

func TestIntegrationCacheDelete(t *testing.T) {
	cache := skipIfRedisUnavailable(t)
	defer cache.Close()

	key := "test:int:delete:" + time.Now().Format("150405.000")

	err := cache.Set(key, "test-value", 30*time.Second)
	if err != nil {
		t.Fatalf("Set a échoué: %v", err)
	}

	exists, err := cache.Exists(key)
	if err != nil {
		t.Fatalf("Exists a échoué: %v", err)
	}
	if !exists {
		t.Fatal("la clé devrait exister après Set")
	}

	err = cache.Delete(key)
	if err != nil {
		t.Fatalf("Delete a échoué: %v", err)
	}

	exists, err = cache.Exists(key)
	if err != nil {
		t.Fatalf("Exists a échoué: %v", err)
	}
	if exists {
		t.Fatal("la clé ne devrait plus exister après Delete")
	}
}

func TestIntegrationCacheTTL(t *testing.T) {
	cache := skipIfRedisUnavailable(t)
	defer cache.Close()

	key := "test:int:ttl:" + time.Now().Format("150405.000")

	err := cache.Set(key, "expires-soon", 2*time.Second)
	if err != nil {
		t.Fatalf("Set a échoué: %v", err)
	}

	// Vérifier que la clé existe
	exists, err := cache.Exists(key)
	if err != nil {
		t.Fatalf("Exists a échoué: %v", err)
	}
	if !exists {
		t.Fatal("la clé devrait exister")
	}

	// Attendre l'expiration
	time.Sleep(3 * time.Second)

	exists, err = cache.Exists(key)
	if err != nil {
		t.Fatalf("Exists a échoué: %v", err)
	}
	if exists {
		t.Fatal("la clé devrait avoir expiré")
	}
}

func TestIntegrationCacheExists(t *testing.T) {
	cache := skipIfRedisUnavailable(t)
	defer cache.Close()

	key := "test:int:exists:" + time.Now().Format("150405.000")

	// Vérifier qu'une clé inexistante retourne false
	exists, err := cache.Exists(key)
	if err != nil {
		t.Fatalf("Exists a échoué: %v", err)
	}
	if exists {
		t.Fatal("la clé ne devrait pas exister")
	}

	// Créer la clé
	err = cache.Set(key, "exists-test", 30*time.Second)
	if err != nil {
		t.Fatalf("Set a échoué: %v", err)
	}

	// Vérifier qu'elle existe maintenant
	exists, err = cache.Exists(key)
	if err != nil {
		t.Fatalf("Exists a échoué: %v", err)
	}
	if !exists {
		t.Fatal("la clé devrait exister après Set")
	}
}

func TestIntegrationCacheConcurrentAccess(t *testing.T) {
	cache := skipIfRedisUnavailable(t)
	defer cache.Close()

	key := "test:int:concurrent:" + time.Now().Format("150405.000")

	// Test de lecture/écriture concurrente
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func(n int) {
			k := key + ":" + string(rune('0'+n))
			err := cache.Set(k, n, 10*time.Second)
			if err != nil {
				t.Errorf("Set concurrent %d a échoué: %v", n, err)
			}

			var val int
			err = cache.Get(k, &val)
			if err != nil && !cache.IsMiss(err) {
				t.Errorf("Get concurrent %d a échoué: %v", n, err)
			}

			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}

// Test d'intégration avec redis.Nil
func TestIntegrationCacheIsMiss(t *testing.T) {
	cache := skipIfRedisUnavailable(t)
	defer cache.Close()

	// Vérifier que IsMiss détecte correctement redis.Nil
	err := cache.Get("nonexistent:key:"+time.Now().Format("150405.000"), nil)
	if !cache.IsMiss(err) {
		t.Fatalf("IsMiss devrait retourner true pour une clé inexistante, got: %v", err)
	}

	// Vérifier que IsMiss retourne false pour d'autres erreurs
	if cache.IsMiss(nil) {
		t.Fatal("IsMiss devrait retourner false pour nil")
	}
}
