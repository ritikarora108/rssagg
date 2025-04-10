package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
    "github.com/go-chi/cors"
)

func main() {
	
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Rss Aggregator")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not set")
	}

    
	router := chi.NewRouter()

    router.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"https://*", "http://*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: false,
        MaxAge:           300,
    }))


    v1Router := chi.NewRouter()

    v1Router.Get("/healthz", readinessHandler)
    v1Router.Get("/err", errorHandler)

    router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
    }
    
    log.Printf("Server starting on port %v", portString)
    
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
