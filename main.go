package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/deanhattenhauer/reservation-api/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// apiConfig holds shared server state accessible across all request handlers.
// Using a struct allows state to be injected into handlers without global variables.
type apiConfig struct {
	// dbQueries provides type-safe access to all database operations via SQLC.
	dbQueries *database.Queries
	// jwt secret for token creation
	jwtSecret string
}

func main() {
	// Load environment variables from .env file before reading any config.
	// Must be called before os.Getenv to ensure variables are available.
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	// SQLC-generated query wrapper provides type-safe database access.
	dbQueries := database.New(db)

	jwtSecret := os.Getenv("JWT_SECRET")

	port := os.Getenv("PORT")

	// apiCfg is the single source of truth for shared server state.
	// Passed to handlers as a pointer receiver so all handlers share the same instance.
	apiCfg := apiConfig{
		dbQueries: dbQueries,
		jwtSecret: jwtSecret,
	}

	// ServeMux routes incoming requests to the appropriate handler.
	// Without registered routes, all requests return 404 by default.
	mux := http.NewServeMux()

	// User endpoints
	mux.HandleFunc("POST /api/v1/users", apiCfg.handlerCreateUser)

	// Login endpoints
	mux.HandleFunc("POST /api/v1/login" ,apiCfg.handlerUserLogin)

	// Token endpoints
	mux.HandleFunc("POST /api/v1/refresh", apiCfg.handlerRefreshToken)
	mux.HandleFunc("POST /api/v1/revoke", apiCfg.handlerRevokeToken)

	// Category endpoints
	mux.HandleFunc("POST /api/v1/categories", apiCfg.handlerCreateCategory)
	mux.HandleFunc("GET /api/v1/categories", apiCfg.handlerGetActiveCategories)

	// Reservation endpoints
	mux.HandleFunc("POST /api/v1/reservations", apiCfg.handlerCreateReservation)
	mux.HandleFunc("GET /api/v1/reservations", apiCfg.handlerGetReservationsByUser)
	mux.HandleFunc("PATCH /api/v1/reservations/{reservationID}", apiCfg.handlerCancelReservation)

	// Admin endpoints
	mux.HandleFunc("GET /api/v1/admin/reservations", apiCfg.handlerGetAllReservations)
	mux.HandleFunc("PATCH /api/v1/admin/reservations/{reservationID}", apiCfg.handlerCancelReservationByAdmin)

	// The mux is injected as the handler so all routing decisions
	// flow through a single, centrally managed router.
	s := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// Logged before blocking so the operator knows the server is ready.
	// Code after ListenAndServe only executes on shutdown or error.
	log.Printf("Server running on port: %s\n", port)

	// ErrServerClosed is not a real error — it signals a clean shutdown.
	// Any other error indicates an unexpected failure and is fatal.
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}