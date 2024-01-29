package main

import (
	"log"
	"net/http"

	"github.com/askaroe/go-course24/tsis1-chunin-exam-api/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("starting API server")

	router := mux.NewRouter()
	log.Println("creating routes")

	router.HandleFunc("/health-check", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/genins", handlers.GetGenins).Methods("GET")
	router.HandleFunc("/genins/{name}", handlers.GetGeninByName).Methods("GET")

	http.ListenAndServe(":8080", router)
}
