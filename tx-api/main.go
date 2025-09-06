package main

import (
	"fmt"
	"log"
	"net/http"
	"tx-api/config"
	logger "tx-api/core/logging"
	"tx-api/core/postgres"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Docker! 🚀")
}

func main() {

	if err := config.Load(); err != nil {
		log.Fatalf("config load failed: %v", err)
	}

	if err := logger.Init(); err != nil {
		log.Fatalf("log init failed: %v", err)
	}
	defer logger.Sync()

	pgService, err := postgres.NewPostgresService()
	if err != nil {
		log.Fatalf("log init pg service: %v", err)
	}

	fmt.Println("pgService", pgService)

	http.HandleFunc("/", handler)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
