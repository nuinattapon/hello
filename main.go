package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"hello/config"
	"hello/fibonacci"
	"hello/handlers"
	"hello/middleware"
)

func main() {
	mux := http.NewServeMux()

	// Register handlers with timing middleware
	mux.Handle("GET /favicon.ico", http.NotFoundHandler())
	mux.Handle("GET /metrics", promhttp.Handler())
	mux.HandleFunc("GET /json", middleware.TimingMiddleware(handlers.JSONHandler))
	mux.HandleFunc("GET /template", middleware.TimingMiddleware(handlers.TemplateHandler))
	mux.HandleFunc("GET /ping", middleware.TimingMiddleware(handlers.PingHandler))
	mux.HandleFunc("GET /version", middleware.TimingMiddleware(handlers.VersionHandler))
	mux.HandleFunc("GET /fibo", middleware.TimingMiddleware(fibonacci.Handler))
	mux.HandleFunc("GET /fibo/", middleware.TimingMiddleware(fibonacci.Handler))
	mux.HandleFunc("GET /fibo/{id}", middleware.TimingMiddleware(fibonacci.Handler))
	mux.HandleFunc("GET /", middleware.TimingMiddleware(handlers.HomeHandler))

	port := config.GetPort()
	log.Printf("Server starting on port %s", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
