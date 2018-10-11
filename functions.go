package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strings"
)

type patient struct {
	ID                    int           `json:"id"`
	FirstName             string        `json:"first_name"`
	LastName              string        `json:"last_name"`
	Gender                string        `json:"gender"`
	PhoneNumber           string        `json:"phone_number"`
	Email                 string        `json:"email"`
	Address               string        `json:"address"`
	VisitDate             string        `json:"visit_date"`
	Diagnosis             string        `json:"diagnosis"`
	DrugCode              string        `json:"drug_code"`
	AdditionalInformation []information `json:"additional_information"`
}

type information struct {
	Notes      string `json:"notes"`
	NewPatient bool   `json:"new_patient"`
	Race       string `json:"race"`
	SSN        string `json:"ssn"`
}

type patientIDs struct {
	IDs []int `json:"patient_ids"`
}

func readPatientsFromJSON() ([]patient, error) {
	jsonFile, err := os.Open("MOCK_DATA.json")
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully opened json file!")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var patients []patient

	json.Unmarshal(byteValue, &patients)

	return patients, nil
}

func getSinglePatient(patientID int) (patient, error) {
	patients, err := readPatientsFromJSON()
	if err != nil {
		return patient{}, err
	}

	for i := range patients {
		if patients[i].ID == patientID {
			return patients[i], nil
		}
	}

	return patient{}, errors.New("Patient does not exist")
}

func searchForPatientsMatching(key, value string) (patientIDs, error) {
	patients, err := readPatientsFromJSON()
	if err != nil {
		return patientIDs{}, err
	}

	var filteredPatients patientIDs

	for _, patient := range patients {
		fieldValue := getField(&patient, key)

		if fieldValue == value {
			filteredPatients.IDs = append(filteredPatients.IDs, patient.ID)
			fmt.Printf("found one")
		}
	}

	if filteredPatients.IDs == nil {
		return patientIDs{}, errors.New("hello")
	}
	return filteredPatients, nil
}

func getField(p *patient, fieldName string) string {
	patientTags := patient{}
	patientStructTags := reflect.ValueOf(patientTags)

	var parsedTag string
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	parsedTagName := reg.ReplaceAllString(fieldName, "")

	for i := 0; i < patientStructTags.Type().NumField(); i++ {
		tags := patientStructTags.Type().Field(i).Name
		fmt.Println(tags)
		if strings.ToLower(patientStructTags.Type().Field(i).Name) == parsedTagName {
			parsedTag = patientStructTags.Type().Field(i).Name
		}
	}
	filledPatientStruct := reflect.ValueOf(p)
	valueOfStructTag := strings.ToLower(reflect.Indirect(filledPatientStruct).FieldByName(parsedTag).String())
	return valueOfStructTag
}

func (p *patient) savePatientToJSON() error {
	patients, err := readPatientsFromJSON()
	if err != nil {
		return err
	}
	p.ID = len(patients) + 1
	patients = append(patients, *p)
	patientJSON, _ := json.Marshal(patients)
	ioutil.WriteFile("MOCK_DATA.json", patientJSON, 0644)
	return nil
}
