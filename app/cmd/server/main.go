package main

import (
	"log"
	"net/http"
	"os"
	"tracing-app/internal/database"
	"tracing-app/internal/handlers"
	"tracing-app/internal/tracing"
)

func main() {
	log.Println("🚀 Starting tracing application...")

	shutdown := tracing.InitTracer()
	defer func() {
		log.Println("🛑 Shutting down tracing...")
		shutdown()
	}()

	log.Println("🔗 Connecting to PostgreSQL...")
	database.InitDB()
	log.Println("✅ Successfully connected to PostgreSQL")

	http.HandleFunc("/search", handlers.SearchHandler)
	log.Println("🔍 Endpoint '/search' is ready")

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🌐 Server is running on port %s", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
