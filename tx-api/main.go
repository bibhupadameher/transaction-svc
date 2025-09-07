package main

import (
	"context"
	"log"
	"net/http"
	"tx-api/config"
	logger "tx-api/core/logging"
	"tx-api/endpoint"
	httptransport "tx-api/http"
	"tx-api/service"
)

func main() {

	if err := config.Load(); err != nil {
		log.Fatalf("config load failed: %v", err)
	}

	if err := logger.Init(); err != nil {
		log.Fatalf("log init failed: %v", err)
	}
	defer logger.Sync()

	ctx := context.Background()
	logger := logger.GetLogger()
	svc, err := service.New(ctx, logger)
	if err != nil {
		log.Fatalf("service init failed: %v", err)
	}
	endpoints := endpoint.NewEndpointSet(svc)

	handler := httptransport.NewHTTPHandler(endpoints)

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", handler)

}
