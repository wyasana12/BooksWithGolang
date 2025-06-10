package routes

import (
	"golang_belajar/controllers/bookcontroller"

	"github.com/gorilla/mux"
)

func BookRoutes(r *mux.Router) {

	router := r.PathPrefix("/books").Subrouter()

	router.HandleFunc("", bookcontroller.Index).Methods("GET")
	router.HandleFunc("", bookcontroller.Create).Methods("POST")
	router.HandleFunc("/{id}", bookcontroller.Detail).Methods("GET")
	router.HandleFunc("/{id}", bookcontroller.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", bookcontroller.Destroy).Methods("DELETE")
}
