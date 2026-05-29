package monitoring

import (
	"context"
	"fmt"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

// ============================================================
//  🔍 TRACING — Suivre une requête à travers tous les services
// ============================================================
// 📖 DDIA Chapitre 8 : "The Trouble with Distributed Systems"
//
// ❓ C'EST QUOI LE TRACING ?
//    Imagine qu'un étudiant clique "Répondre" sur une question.
//    Cette action traverse plusieurs services :
//    1. API Gateway (HTTP) → 2. Service Learning → 3. PostgreSQL → 4. NATS
//
//    Si la requête prend 5 secondes, où est le problème ?
//    - Sans tracing : tu ne sais pas, tu dois deviner
//    - Avec tracing : tu vois que PostgreSQL a pris 4.5 secondes
//
//    Chaque étape est une "span" (segment) dans le "trace" (parcours).
//    Les spans sont envoyées à Jaeger qui les affiche dans un graphique.
//
// 🧠 COMMENT ÇA MARCHE ?
//    1. Quand une requête arrive, on crée un "trace" (ID unique)
//    2. Chaque service ajoute ses "spans" (ex: "query PostgreSQL")
//    3. Toutes les spans sont envoyées à Jaeger (collecteur)
//    4. Jaeger affiche le parcours complet dans son interface web
//
// 🔗 LIENS UTILES :
//    - https://opentelemetry.io/docs/
//    - https://www.jaegertracing.io/
//    - go.opentelemetry.io/otel
// ============================================================

// TracerProvider gère l'envoi des traces à Jaeger.
// C'est le "responsable" qui collecte toutes les spans
// et les envoie au serveur Jaeger.
type TracerProvider struct {
	provider *sdktrace.TracerProvider
}

// NewTracerProvider crée un nouveau provider de tracing.
//
// 🛡️ PRODUCTION :
//    - Sampler configurable : en prod on trace 10% des requêtes
//      (100% ferait perdre trop de perf)
//    - En dev on trace 100% pour déboguer
//    - Export via OTLP gRPC vers Jaeger
//    - Batching : les spans sont envoyées par lots (optimisé)
//
// 📝 EXEMPLE D'UTILISATION :
//    // En dev : tracer 100% des requêtes
//    tp, err := monitoring.NewTracerProvider("aeroforge-api", "localhost:4317", 1.0)
//
//    // En prod : tracer 10% des requêtes
//    tp, err := monitoring.NewTracerProvider("aeroforge-api", "jaeger:4317", 0.1)
//
//    if err != nil { log.Fatal(err) }
//    defer tp.Shutdown()
//
//    // Créer un tracer pour un package
//    tracer := tp.Tracer("learning")
//
//    // Créer une span (segment de trace)
//    ctx, span := tracer.Start(ctx, "CreateExam")
//    defer span.End()
//
//    // Ajouter des infos à la span
//    span.SetAttributes(attribute.String("student_id", "123"))
func NewTracerProvider(serviceName, jaegerURL string, sampleRate float64) (*TracerProvider, error) {
	// Exporter OTLP = envoie les traces à Jaeger via gRPC
	exporter, err := otlptracegrpc.New(context.Background(),
		otlptracegrpc.WithEndpoint(jaegerURL), // "localhost:4317"
		otlptracegrpc.WithInsecure(),           // Pas de TLS en dev
	)
	if err != nil {
		return nil, fmt.Errorf("création exporter OTLP: %w", err)
	}

	// Resource = infos sur le service qui envoie les traces
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(serviceName), // "aeroforge-api"
	)

	// Sampler = combien de requêtes tracer ?
	// - 1.0 = 100% (dev)
	// - 0.1 = 10% (prod)
	// - 0.01 = 1% (très gros volume)
	sampler := sdktrace.TraceIDRatioBased(sampleRate)

	// Provider = le moteur qui gère les spans
	provider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter), // Envoie les spans par lots (optimisé)
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sampler), // Échantillonnage configurable
	)

	// Définir le provider global pour tout le programme
	otel.SetTracerProvider(provider)

	log.Printf("🔍 Tracing OpenTelemetry initialisé pour %s → %s (sample rate: %.0f%%)",
		serviceName, jaegerURL, sampleRate*100)
	return &TracerProvider{provider: provider}, nil
}

// Tracer retourne un tracer pour un package donné.
// Chaque package peut avoir son propre tracer.
func (tp *TracerProvider) Tracer(name string) trace.Tracer {
	return tp.provider.Tracer(name)
}

// Shutdown arrête proprement l'envoi des traces.
// Important : appelle ça à la fin du programme pour
// être sûr que toutes les spans sont envoyées.
func (tp *TracerProvider) Shutdown() {
	if tp.provider != nil {
		_ = tp.provider.Shutdown(context.Background())
	}
}

// ============================================================
//  🧪 TESTER LE TRACING
// ============================================================
// 1. Démarre Jaeger : docker compose up jaeger -d
// 2. Ouvre Jaeger : http://localhost:16686
// 3. Lance l'API : go run cmd/api-gateway/main.go
// 4. Fais une requête : curl http://localhost:8080/api/questions
// 5. Regarde Jaeger → tu devrais voir la trace apparaître
// ============================================================
