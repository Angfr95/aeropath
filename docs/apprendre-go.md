# 🎯 Plan d'étude Go — De zéro à AeroForge

## Pourquoi Go ?

Go est parfait pour AeroForge car :
- **Simple** : pas de classes, pas d'héritage, pas de génériques tordus
- **Rapide** : aussi rapide que du C, mais aussi facile que du Python
- **Un seul outil** : `go build`, `go test`, `go fmt` — tout est inclus
- **Idéal pour les APIs** : la bibliothèque standard fait tout (HTTP, JSON, bases de données)

---

## 📚 Étape 1 : Les bases (2-3 jours)

### Jour 1 : Le Go Tour (2h)
https://go.dev/tour/

Faire **TOUT** le tour, mais insister sur :
- Packages, variables, fonctions
- `for` (la seule boucle en Go)
- `if`, `switch`
- Structs et méthodes
- Interfaces (concept clé dans AeroForge)
- Erreurs (`error` type)

### Jour 2 : Mini-projets (4h)

**Exercice 1 : Hello World**
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, AeroForge!")
}
```
→ `go run main.go`

**Exercice 2 : Calculatrice**
```go
package main

import "fmt"

func add(a, b int) int {
    return a + b
}

func main() {
    result := add(3, 4)
    fmt.Printf("3 + 4 = %d\n", result)
}
```
→ Ajoute `sub`, `mul`, `div`

**Exercice 3 : Structures**
```go
type Question struct {
    ID       int
    Content  string
    License  string
    Answers  []string
    Correct  int
}

func main() {
    q := Question{
        ID:      1,
        Content: "Quelle est la vitesse du son ?",
        License: "PPL",
        Answers: []string{"340 m/s", "500 m/s", "1000 m/s"},
        Correct: 0,
    }
    fmt.Printf("Question: %s\n", q.Content)
}
```

### Jour 3 : Fichiers et JSON (4h)

**Exercice 4 : Lire/écrire un fichier**
```go
package main

import (
    "encoding/json"
    "os"
)

type Student struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    // Écrire
    s := Student{Name: "Alice", Email: "alice@example.com"}
    data, _ := json.Marshal(s)
    os.WriteFile("student.json", data, 0644)

    // Lire
    data, _ = os.ReadFile("student.json")
    var s2 Student
    json.Unmarshal(data, &s2)
    println(s2.Name)
}
```

---

## 🌐 Étape 2 : Serveur HTTP (2 jours)

### Jour 4 : Premier serveur

```go
package main

import (
    "encoding/json"
    "net/http"
)

func main() {
    http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Hello AeroForge!",
        })
    })
    http.ListenAndServe(":8080", nil)
}
```

**Exercice :** Ajoute un endpoint `/api/questions` qui retourne une liste de questions en JSON.

### Jour 5 : Utiliser Gin (le framework d'AeroForge)

```go
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/api/questions", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "questions": []string{"Q1", "Q2"},
        })
    })
    r.Run(":8080")
}
```

**Exercice :** Ajoute un paramètre `?license=PPL` et filtre les questions.

---

## 🗄️ Étape 3 : Base de données (2 jours)

### Jour 6 : PostgreSQL avec Go

```go
package main

import (
    "database/sql"
    _ "github.com/lib/pq"
)

func main() {
    connStr := "postgres://user:pass@localhost:5432/aeroforge?sslmode=disable"
    db, _ := sql.Open("postgres", connStr)
    
    rows, _ := db.Query("SELECT id, content FROM questions LIMIT 5")
    for rows.Next() {
        var id int
        var content string
        rows.Scan(&id, &content)
        println(id, content)
    }
}
```

### Jour 7 : Repository pattern (comme dans AeroForge)

Regarde `internal/persistence/postgres/question_repo.go` — c'est exactement ce pattern.

---

## 🧵 Étape 4 : Concurrence (2 jours)

### Jour 8 : Goroutines et channels

```go
func main() {
    ch := make(chan string)
    
    go func() {
        ch <- "Question chargée !"
    }()
    
    msg := <-ch
    println(msg)
}
```

### Jour 9 : Worker pool

Regarde comment le worker NATS fonctionne dans `internal/events/consumer.go` — c'est un pattern de worker pool.

---

## 🎯 Étape 5 : Attaquer AeroForge

Une fois les étapes 1-4 maîtrisées, tu peux implémenter dans cet ordre :

1. **Redis** (`internal/persistence/redis/cache.go`) — le plus simple
2. **Prometheus** (`internal/monitoring/prometheus.go`) — juste des compteurs
3. **ClickHouse** (`internal/persistence/clickhouse/analytics.go`) — des requêtes SQL
4. **NATS Producer/Consumer** (`internal/events/`) — connexion + publish/subscribe
5. **OpenTelemetry** (`internal/monitoring/tracing.go`) — le plus conceptuel
6. **gRPC** (`internal/transport/grpc/server.go`) — nécessite de comprendre les .proto

---

## 📖 Ressources

- **Go Tour** (2h) : https://go.dev/tour/
- **Go by Example** : https://gobyexample.com/
- **Effective Go** : https://go.dev/doc/effective_go
- **Documentation Gin** : https://gin-gonic.com/docs/
- **Le code d'AeroForge** : lis les fichiers dans l'ordre :
  1. `internal/domain/` (les modèles)
  2. `internal/persistence/postgres/` (les repositories)
  3. `internal/transport/http/` (les handlers)
  4. `cmd/api-gateway/main.go` (le point d'entrée)

---

## ⏱️ Planning recommandé

| Semaine | Sujet | Objectif |
|---------|-------|----------|
| 1 | Go Tour + mini-projets | Savoir écrire du Go basique |
| 2 | Serveur HTTP + Gin | Savoir créer une API REST |
| 3 | PostgreSQL + Repository | Savoir lire/écrire en base |
| 4 | Goroutines + Channels | Comprendre la concurrence |
| 5 | Redis | Implémenter le cache |
| 6 | Prometheus | Ajouter les métriques |
| 7 | ClickHouse | Analytics |
| 8 | NATS | Event-driven |
| 9 | OpenTelemetry | Tracing |
| 10 | gRPC | Communication inter-services |

**Rappel :** Les squelettes dans `internal/` ont des TODO et des commentaires en français. Tu n'as pas à deviner quoi faire — tout est expliqué ligne par ligne.
