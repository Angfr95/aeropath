package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	log.Printf("worker démarré")
	_ = os.Getenv("DATABASE_URL")

	// TODO: implémenter le worker
	select {}
}
