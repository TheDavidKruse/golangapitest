package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

func amIRunning(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("I am up and running!"))
}

func getAllPatients(w http.ResponseWriter, r *http.Request) {
	allPatients, err := readPatientsFromJSON()

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("server could not read json file"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(allPatients)
	}
}

func getPatientByID(w http.ResponseWriter, r *http.Request) {
	patientID, err := strconv.Atoi(mux.Vars(r)["patientID"])

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Id parameter is unacceptable or patient does not exist"))
	}

	patient, getPatientErr := getSinglePatient(patientID)

	if getPatientErr != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("There was an error on our end"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(patient)
	}

}

func createPatient(w http.ResponseWriter, r *http.Request) {
	patient := patient{}
	json.NewDecoder(r.Body).Decode(&patient)
	err := patient.savePatientToJSON()
	if err != nil {
		fmt.Printf("Error writing patient to json file: %+v", err)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("patient could not be created or already exists"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("{}"))
	}
}

func searchPatients(w http.ResponseWriter, r *http.Request) {
	urlQuery := r.URL.Query()
	mapKeys := reflect.ValueOf(urlQuery).MapKeys()
	key := mapKeys[0].String()
	value := urlQuery[key][0]

	patientIDs, err := searchForPatientsMatching(key, value)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("patient could not be created or already exists"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(patientIDs)
	}

}