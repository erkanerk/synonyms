package main

import (
	"log"
	"net/http"
	"synonyms-backend/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
