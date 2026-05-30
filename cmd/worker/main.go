package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	log.Printf("worker démarré")
	_ = os.Getenv("DATABASE_URL")

	// TODO: implémenter le worker
	for {
		log.Printf("worker: en attente...")
		time.Sleep(30 * time.Second)
	}
}
