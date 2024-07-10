package routes

import (
	"synonyms-backend/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/word", handlers.AddWord).Methods("POST")
	r.HandleFunc("/synonym/{word}", handlers.AddSynonym).Methods("POST")
	r.HandleFunc("/synonyms/{word}", handlers.GetSynonyms).Methods("GET")
}
