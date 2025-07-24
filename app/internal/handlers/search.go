package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"tracing-app/internal/database"
	"tracing-app/internal/tracing"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracing.Tracer.Start(r.Context(), "handle_request")
	defer span.End()

	query := r.URL.Query().Get("q")
	if query == "" {
		query = "default"
	}

	log.Printf("🔍 Handling search request for query: %s", query)

	// Simulate database query
	func() {
		ctxDatabase, databaseSpan := tracing.Tracer.Start(ctx, "database_query")
		defer databaseSpan.End()

		log.Println("📦 Querying database for current time...")
		now := database.QueryNow(ctxDatabase)
		log.Printf("🕒 Current time from DB: %s", now)

		// Simulate DB processing delay
		time.Sleep(time.Duration(rand.Intn(200)+100) * time.Millisecond)
		log.Println("✅ Database query completed")
	}()

	// Simulate external API call
	func() {
		_, apiSpan := tracing.Tracer.Start(ctx, "external_api_call")
		defer apiSpan.End()

		log.Println("🌐 Calling external API...")
		resp, err := http.Get(fmt.Sprintf("https://httpbin.org/get?query=%s", query))
		if err != nil {
			log.Printf("⚠️ Failed to call external API: %v", err)
			return
		}
		defer resp.Body.Close()
		_, _ = ioutil.ReadAll(resp.Body)
		log.Println("✅ External API call completed")
	}()

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"query":  query,
	})
	log.Println("📤 Response sent to client")
}
