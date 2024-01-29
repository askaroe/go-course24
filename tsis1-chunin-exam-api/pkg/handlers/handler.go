package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/askaroe/go-course24/tsis1-chunin-exam-api/api"
	"github.com/gorilla/mux"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API server is running")
}

func GetGenins(w http.ResponseWriter, r *http.Request) {
	genins := api.Genins

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(genins)
	if err != nil {
		return
	}

	w.Write(jsonResponse)

}

func GetGeninByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	genin := api.GetGeninByName(name)
	w.Header().Set("Content-Type", "application/json")

	if genin == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(genin)
	if err != nil {
		return
	}

	w.Write(jsonResponse)

}
