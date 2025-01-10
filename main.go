package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/Klimentin0/sheptalka/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db             *database.Queries
}

func main() {
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	dbQueries := database.New(db)

	const port = "8080"
	const root = "."

	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
		db:             dbQueries,
	}

	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(root)))))
	mux.HandleFunc("GET /api/healthz", healthHandler)
	mux.HandleFunc("POST /api/users", apiCfg.createUserHandler)
	//api шепотов
	mux.HandleFunc("POST /api/shepots", apiCfg.shepotHandler)
	mux.HandleFunc("GET /api/shepots", apiCfg.getShepotsHandler)
	mux.HandleFunc("GET /api/shepots/{shepotID}", apiCfg.getShepotHandler)
	//админка
	mux.HandleFunc("GET /admin/count", apiCfg.reqCountHandler)
	mux.HandleFunc("POST /admin/reset", apiCfg.resetDb)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("serving data from %s on port: %s\n", root, port)
	log.Fatal(srv.ListenAndServe())
}
