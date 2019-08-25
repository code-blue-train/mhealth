package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/* Patient struct to support HL7 FHIR Release 4 Patient json representation
https://www.hl7.org/fhir/patient-example.json.html
*/

type Patient struct {
	ResourceType string `json:"resourceType"`
	ID           string `json:"id"`
	Meta         struct {
		VersionID   string `json:"versionId"`
		LastUpdated string `json:"lastUpdated"`
		Profile     string `json:"profile"`
	} `json:"meta"`
	Text struct {
		Status string `json:"status"`
		Div    string `json:"div"`
	} `json:"text"`
	Extension []struct {
		URL       string `json:"url"`
		Extension []struct {
			URL         string `json:"url"`
			ValueCoding struct {
				System  string `json:"system"`
				Code    string `json:"code"`
				Display string `json:"display"`
			} `json:"valueCoding`
			ValueString string `json:"valueString"`
		} `json:"extension"`
		ValueCode string `json:"valueCode"`
	}
	Identifier []struct {
		Use  string `json:"use"`
		Type struct {
			Coding []struct {
				System  string `json:"system"`
				Code    string `json:"code"`
				Display string `json:"display"`
			} `json:"coding"`
			Text string `json:"text"`
		} `json:"type"`
		System string `json:"system"`
		Value  string `json:"value"`
	}
	Active bool `json:"active"`
	Name   []struct {
		Use    string   `json:"use"`
		Family string   `json:"family"`
		Given  []string `json:"given"`
	} `json:"name"`
	Telecom []struct {
		System string `json:"system"`
		Value  string `json:"value"`
		Use    string `json:"use"`
	} `json:"telecom"`
	Gender    string `json:"gender"`
	Birthdate string `json:"birthDate"`
	Address   []struct {
		Line       []string `json:"line"`
		City       string   `json:"city"`
		State      string   `json:"state"`
		PostalCode string   `json:"postalCode"`
		Country    string   `json:"country"`
	}
}

func main() {
	/*
		// Open our jsonFile
		jsonFile, err := os.Open("Patient-example.json")
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully Opened Patient-example.json")

		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()
	*/

	resp, err := http.Get("http://hapi.fhir.org/baseR4/Patient/3624/_history/1?_format=json&_pretty=true")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	// read our opened xmlFile as a byte array.
	//byteValue, _ := ioutil.ReadAll(jsonFile)

	body, err := ioutil.ReadAll(resp.Body)

	// we initialize our Users array
	var patient Patient

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(body, &patient)

	fmt.Println(patient.ResourceType)
	fmt.Println("Patient.Meta.VersionID: " + patient.Meta.VersionID)
	fmt.Println("Patient.Meta.LastUpdated: " + patient.Meta.LastUpdated)
	fmt.Println("Patient.Meta.Profile: " + patient.Meta.Profile)
	fmt.Printf("Patient.Active: %t\n", patient.Active)
	fmt.Println(patient.Text.Status)
	fmt.Println(patient.Text.Div)
	for i := 0; i < len(patient.Extension); i++ {
		fmt.Println(patient.Extension[i].URL)
		fmt.Println(patient.Extension[i].ValueCode)
		for p := 0; p < len(patient.Extension[i].Extension); p++ {
			fmt.Println(patient.Extension[i].Extension[p].URL)
			fmt.Println(patient.Extension[i].Extension[p].ValueCoding.System)
			fmt.Println(patient.Extension[i].Extension[p].ValueCoding.Code)
			fmt.Println(patient.Extension[i].Extension[p].ValueCoding.Display)
			fmt.Println(patient.Extension[i].Extension[p].ValueString)

		}
	}

	for i := 0; i < len(patient.Identifier); i++ {
		fmt.Println(patient.Identifier[i].Use)
		fmt.Println(patient.Identifier[i].Type)
		fmt.Println(patient.Identifier[i].System)
		fmt.Println(patient.Identifier[i].Value)
	}

	for i := 0; i < len(patient.Name); i++ {
		fmt.Println(patient.Name[i].Family)
		fmt.Println(patient.Name[i].Use)
		for g := 0; g < len(patient.Name[i].Given); g++ {
			fmt.Println(patient.Name[i].Given[g])
		}
	}

	for i := 0; i < len(patient.Telecom); i++ {
		fmt.Println(patient.Telecom[i].System)
		fmt.Println(patient.Telecom[i].Value)
		fmt.Println(patient.Telecom[i].Use)
	}

	fmt.Println(patient.Gender)
	fmt.Println(patient.Birthdate)

	for i := 0; i < len(patient.Address); i++ {
		for l := 0; l < len(patient.Address[i].Line); l++ {
			fmt.Println(patient.Address[i].Line[l])
		}
		fmt.Println(patient.Address[i].City)
		fmt.Println(patient.Address[i].State)
		fmt.Println(patient.Address[i].PostalCode)
		fmt.Println(patient.Address[i].Country)
	}
}
