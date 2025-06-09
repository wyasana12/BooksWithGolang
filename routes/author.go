package routes

import (
	"golang_belajar/controllers/authorcontroller"

	"github.com/gorilla/mux"
)

func AuthorRoutes(r *mux.Router) {
	router := r.PathPrefix("/authors").Subrouter()

	router.HandleFunc("", authorcontroller.Index).Methods("GET")
	router.HandleFunc("", authorcontroller.Create).Methods("POST")
	router.HandleFunc("/{id}", authorcontroller.Detail).Methods("GET")
	router.HandleFunc("/{id}", authorcontroller.Update).Methods("PUT")
	router.HandleFunc("/{id}/Delete", authorcontroller.Destroy).Methods("DELETE")
}
