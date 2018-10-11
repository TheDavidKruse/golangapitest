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
	err := patient.createNewPatientInJSON()
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

func updatePatient(w http.ResponseWriter, r *http.Request) {

	// The easiest way to make this work was to stream the new json data into the old struct and save it so hopefully that isnt a bad thing

	patientID, _ := strconv.Atoi(mux.Vars(r)["patientID"])

	patient, _ := getSinglePatient(patientID)

	json.NewDecoder(r.Body).Decode(&patient)

	err := updatePatientInJSON(patient, patientID)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("oops I did it again, I played withh you data, I lost in the dataaaaaa"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("{}"))
	}
}
