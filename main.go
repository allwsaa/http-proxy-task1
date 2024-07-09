package main

import (
	"log"
	"net/http"

	"http-proxy-task1/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title HTTP Proxy Service
// @version 1.0
// @description This is a simple server for proxying HTTP requests
// @host localhost:8080
// @BasePath/
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	//swagger endpoint
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	//healthcheck endpoint
	r.Get("/health", handlers.HealthCheckHandler)
	//proxy endpoint
	r.Post("/proxy", handlers.PrHandler)

	log.Println("Starting server: 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

}
