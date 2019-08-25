package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Diagnostic Report
type DiagnosticReport struct {
	FullURL  string `json:"fullUrl"`
	Resource struct {
		ResourceType string `json:"resourceType"`
		ID           string `json:"id"`
		Text         struct {
			Status string `json:"status"`
			Div    string `json:"div"`
		} `json:"text"`
		Status string `json:"status"`
		Code   struct {
			Coding []struct {
				System  string `json:"system"`
				Code    string `json:"code"`
				Display string `json:"display"`
			} `json:"coding"`
			Text string `json:"text"`
		} `json:"code"`
		Subject struct {
			Reference string `json:"reference"`
		} `json:"subject"`
		Performer []struct {
			Reference string `json:"reference"`
			Display   string `json:"display"`
		} `json:"performer"`
		ValueQuantity struct {
			Value  uint   `json:"value"`
			Unit   string `json:"unit"`
			System string `json:"system"`
			Code   string `json:"code"`
		} `json:"valueQuantity`
		ReferenceRange []struct {
			Low struct {
				Value  uint   `json:"value"`
				Unit   string `json:"unit"`
				System string `json:"system"`
				Code   string `json:"code"`
			} `json:"low"`
			High struct {
				Value  uint   `json:"value"`
				Unit   string `json:"unit"`
				System string `json:"system"`
				Code   string `json:"code"`
			} `json:"high"`
		} `json:"referenceRange"`
	} `json:"resource"`
}

func main() {

	// Open our jsonFile
	jsonFile, err := os.Open("sample/one-test.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened one-test.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var dr DiagnosticReport

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &dr)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	fmt.Println(dr.FullURL)
	fmt.Println(dr.Resource.Text)
	fmt.Println(dr.Resource.Status)

	for i := 0; i < len(dr.Resource.Code.Coding); i++ {
		fmt.Println(dr.Resource.Code.Coding[i].System)
		fmt.Println(dr.Resource.Code.Coding[i].Code)
		fmt.Println(dr.Resource.Code.Coding[i].Display)
	}

	fmt.Println(dr.Resource.Subject)

	for i := 0; i < len(dr.Resource.Performer); i++ {
		fmt.Println(dr.Resource.Performer[i].Reference)
		fmt.Println(dr.Resource.Performer[i].Display)
	}

	fmt.Println(dr.Resource.ValueQuantity.Value)
	fmt.Println(dr.Resource.ValueQuantity.Unit)
	fmt.Println(dr.Resource.ValueQuantity.System)
	fmt.Println(dr.Resource.ValueQuantity.Code)

	for i := 0; i < len(dr.Resource.ReferenceRange); i++ {

		fmt.Println(dr.Resource.ReferenceRange[i].High.Value)
		fmt.Println(dr.Resource.ReferenceRange[i].High.Unit)
		fmt.Println(dr.Resource.ReferenceRange[i].High.System)
		fmt.Println(dr.Resource.ReferenceRange[i].High.Code)

		fmt.Println(dr.Resource.ReferenceRange[i].Low.Value)
		fmt.Println(dr.Resource.ReferenceRange[i].Low.Unit)
		fmt.Println(dr.Resource.ReferenceRange[i].Low.System)
		fmt.Println(dr.Resource.ReferenceRange[i].Low.Code)

	}

}
