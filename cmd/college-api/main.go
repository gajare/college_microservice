package main

import (
	"log"
	"net/http"

	handlers "github.com/gajare/college_api/handler"
	"github.com/gajare/college_api/middleware"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize router
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)
	// Define endpoints
	router.HandleFunc("/colleges", handlers.GetColleges).Methods("GET")
	router.HandleFunc("/colleges/{id}", handlers.GetCollege).Methods("GET")
	router.HandleFunc("/colleges", handlers.CreateCollege).Methods("POST")
	router.HandleFunc("/colleges/{id}", handlers.UpdateCollege).Methods("PUT")
	router.HandleFunc("/colleges/{id}", handlers.DeleteCollege).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", router))
}
