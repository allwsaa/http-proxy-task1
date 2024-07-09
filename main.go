package main

import (
	"log"
	"net/http"

	"http-proxy-task1/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//get/post
	r.Get("/health", handlers.HealthCheckHandler)
	r.Post("/proxy", handlers.PrHandler)

	log.Println("Starting server: 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

}
