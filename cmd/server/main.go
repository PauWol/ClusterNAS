package main

import (
	"log"
	"net/http"

	"github.com/PauWol/ClusterNAS/internal/api"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, 
		AllowedMethods: []string{"GET","POST","PUT","DELETE"},
	}))
	
	// Routes
	r.Get("/api/drives", api.GetDrives)
	r.Post("/api/drives", api.AddDrive)

	log.Println("ClusterNAS API running on :8080")
	http.ListenAndServe(":8080", r)
}
