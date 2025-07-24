package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	}

	var err error
	log.Println("🔌 Connecting to PostgreSQL database...")
	pool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("❌ Failed to connect to PostgreSQL: %v", err)
	}
	log.Println("✅ Connected to PostgreSQL successfully!")
}

func QueryNow(ctx context.Context) string {
	var now string
	err := pool.QueryRow(ctx, "SELECT NOW()").Scan(&now)
	if err != nil {
		log.Printf("⚠️ Failed to fetch current time from database: %v", err)
		return "error"
	}
	log.Printf("⏰ Current time from DB: %s", now)
	return now
}
