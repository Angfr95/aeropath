package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

// ============================================================
//  🎯 SERVEUR gRPC — Communication entre services
// ============================================================
// 📖 DDIA Chapitre 4 : "Encoding and Evolution" (RPC)
//
// ❓ C'EST QUOI gRPC ?
//    gRPC permet à deux programmes (ex: API Gateway et Worker)
//    de se parler comme s'ils étaient dans le même programme.
//    Au lieu d'appeler une URL HTTP, tu "appelles une fonction"
//    directement sur l'autre serveur.
//
//    Exemple : tu veux demander au service "Learning" de créer un examen.
//    Sans gRPC : tu fais un POST /api/exam avec du JSON
//    Avec gRPC : tu appelles learning.LearningService.CreateExam()
//    C'est plus rapide, plus structuré, et ça gère les erreurs mieux.
//
// 🧱 COMMENT ÇA MARCHE ?
//    1. On définit les "services" dans des fichiers .proto (voir internal/proto/)
//    2. On génère du code Go automatiquement avec protoc
//    3. On implémente les interfaces générées dans ce fichier
//    4. Le client appelle les fonctions comme si c'était local
//
// 🔗 LIENS UTILES :
//    - https://grpc.io/docs/languages/go/
//    - Les fichiers .proto sont dans internal/proto/
// ============================================================

// Server représente le serveur gRPC.
// C'est le "boîte aux lettres" qui écoute les appels des autres services.
type Server struct {
	srv *grpc.Server
}

// NewServer crée un nouveau serveur gRPC avec intercepteurs.
//
// 🛡️ PRODUCTION :
//    - Keepalive : toutes les 30s, on vérifie que le client est toujours là
//    - Max message size : 4MB max (évite les attaques par gros messages)
//    - Timeout de connexion : 20s max pour établir une connexion
//    - Intercepteur de logging : chaque appel est loggé
//
// 🧠 À SAVOIR :
//    Un "intercepteur" est comme un middleware dans Gin.
//    Il s'exécute AVANT chaque appel gRPC.
//    Utile pour : logger, vérifier l'auth, mesurer le temps, tracer.
func NewServer() *Server {
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			loggingInterceptor,
		),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     5 * time.Minute,  // Fermer les connexions inactives après 5min
			MaxConnectionAge:      30 * time.Minute,  // Forcer le renouvellement toutes les 30min
			MaxConnectionAgeGrace: 5 * time.Second,   // Temps pour terminer les requêtes en cours
			Time:                  30 * time.Second,  // Ping toutes les 30s
			Timeout:               5 * time.Second,   // Timeout du ping
		}),
		grpc.MaxRecvMsgSize(4 * 1024 * 1024), // Max 4MB par message reçu
		grpc.MaxSendMsgSize(4 * 1024 * 1024), // Max 4MB par message envoyé
	)

	// Activer la reflection pour grpcurl (debug)
	reflection.Register(srv)

	return &Server{srv: srv}
}

// Start démarre le serveur gRPC sur le port donné.
func (s *Server) Start(port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("échec écoute port %s: %w", port, err)
	}

	log.Printf("🚀 Serveur gRPC démarré sur :%s", port)
	return s.srv.Serve(lis)
}

// Stop arrête proprement le serveur gRPC.
// GracefulStop() attend que les requêtes en cours finissent
// avant de fermer. C'est plus propre que Kill().
func (s *Server) Stop() {
	if s.srv != nil {
		s.srv.GracefulStop()
	}
}

// ============================================================
//  🔧 INTERCEPTEUR (middleware gRPC)
// ============================================================
// Un intercepteur est une fonction qui s'exécute avant chaque appel.
// Comme les middlewares dans Gin, mais pour gRPC.

func loggingInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	start := time.Now()
	log.Printf("📥 Appel gRPC: %s", info.FullMethod)

	resp, err := handler(ctx, req)

	log.Printf("📤 Réponse: %s en %v", info.FullMethod, time.Since(start))
	return resp, err
}
