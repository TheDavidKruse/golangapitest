package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", amIRunning).Methods("GET")
	r.HandleFunc("/patient", getAllPatients).Methods("GET")
	r.HandleFunc("/search", searchPatients).Methods("GET")
	r.HandleFunc("/patient/{patientID}", getPatientByID).Methods("GET")
	r.HandleFunc("/patient/{patientID}", updatePatient).Methods("PATCH")
	r.HandleFunc("/patient", createPatient).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
