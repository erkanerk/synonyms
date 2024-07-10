package main

import (
	"log"
	"net/http"
	"synonyms-backend/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)


	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
    origins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	
	log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", handlers.CORS(origins, headers, methods)(r)))
}
