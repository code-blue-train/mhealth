package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

/* DiagnosticReport struct to support HL7 FHIR Release 4 Diagnostic Report json representation
https://www.hl7.org/fhir/diagnosticreport-example.json.html
*/
type DiagnosticReport struct {
	ResourceType string `json:"resourceType"`
	ID           string `json:"id"`
	Type         string `json:"type"`
	Entry        []struct {
		FullURL  string `json:"fullUrl"`
		Resource struct {
			ResourceType string `json:"resourceType"`
			ID           string `json:"id"`
			Meta         struct {
				LastUpdated string `json:"lastUpdated"`
				Tag         []struct {
					System  string `json:"system"`
					Code    string `json:"code"`
					Display string `json:"display"`
				} `json:"tag"`
			} `json:"meta"`
			Text struct {
				Status string `json:"status"`
				Div    string `json:"div"`
			} `json:"text"`
			AccessionIdentifier struct {
				System string `json:"system"`
				Value  string `json:"value"`
			} `json:"accessionIdentifier"`
			Identifier []struct {
				System string `json:"system"`
				Value  string `json:"value"`
			} `json:"identifier"`
			Status   string `json:"status"`
			Category []struct {
				Coding []struct {
					System string `json:"system"`
					Code   string `json:"code"`
				} `json:"coding"`
			} `json:"category"`
			Code struct {
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
			Collection struct {
				Collector struct {
					Reference string `json:"reference"`
				} `json:"collector"`
				CollectedDateTime string `json:"collectedDateTime"`
			} `json:"collection"`
			Container []struct {
				Type struct {
					Coding []struct {
						System  string `json:"system"`
						Code    string `json:"code"`
						Display string `json:"display"`
					} `json:"coding"`
				} `json:"type"`
			} `json:"container"`
			EffectiveDateTime string `json:"effectiveDateTime"`
			Issued            string `json:"issued"`
			Performer         []struct {
				Reference string `json:"reference"`
				Display   string `json:"display"`
			} `json:"performer"`
			ValueCodeableConcept struct {
				Coding []struct {
					System  string `json:"system"`
					Code    string `json:"code"`
					Display string `json:"display"`
				} `json:"coding"`
			} `json:"valueCodeableConcept"`
			Specimen []struct {
				Reference string `json:"reference"`
				Display   string `json:"display"`
			} `json:"specimen"`
			HasMember []struct {
				Reference string `json:"reference"`
				Display   string `json:"display"`
			} `json:"hasMember"`
			Result []struct {
				Reference string `json:"reference"`
				Display   string `json:"display"`
			} `json:"result"`
			ValueQuantity struct {
				Value  uint64 `json:"value"`
				Unit   string `json:"unit"`
				System string `json:"system"`
				Code   string `json:"code"`
			} `json:"valueQuantity"`
			Interpretation []struct {
				Coding []struct {
					System string `json:"system"`
					Code   string `json:"code"`
				} `json:"coding"`
			} `json:"interpretation"`
			ReferenceRange []struct {
				Low struct {
					Value  uint64 `json:"value"`
					Unit   string `json:"unit"`
					System string `json:"system"`
					Code   string `json:"code"`
				} `json:"low"`
				High struct {
					Value  uint64 `json:"value"`
					Unit   string `json:"unit"`
					System string `json:"system"`
					Code   string `json:"code"`
				} `json:"high"`
			} `json:"referenceRange"`
		} `json:"resource"`
	} `json:"entry"`
}

func main() {
	/*
		// Open our jsonFile
		//jsonFile, err := os.Open("diagnosticreport-example.json")
		jsonFile, err := os.Open("diagnostic-example-ghp.json")
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Successfully Opened diagnosticreport-example.json")
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()
	*/

	resp, err := http.Get("http://hapi.fhir.org/baseR4/Bundle?_format=json&_pretty=true")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	// read our opened xmlFile as a byte array.
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}

	// we initialize our Users array
	var dr DiagnosticReport

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(body, &dr)

	fmt.Println(dr.ResourceType)

	for i := 0; i < len(dr.Entry); i++ {
		fmt.Println("---------------------------------------------------------")
		fmt.Println("Entry Number: " + strconv.Itoa(i) + "\n")
		fmt.Println("Entry.Resource.EffectiveDateTime: " + dr.Entry[i].Resource.EffectiveDateTime + "\n")
		fmt.Println("Entry.Resource.Issued: " + dr.Entry[i].Resource.Issued + "\n")
		fmt.Println("Entry.Resource.Subject.Reference: " + dr.Entry[i].Resource.Subject.Reference + "\n")
		fmt.Println("Entry.Resource.Collection.Collector.Reference: " + dr.Entry[i].Resource.Collection.Collector.Reference + "\n")
		fmt.Println("Entry.Resource.Collection.Collector.CollectedDateTime: " + dr.Entry[i].Resource.Collection.CollectedDateTime + "\n")
		fmt.Println("Entry.FullURL: " + dr.Entry[i].FullURL + "\n")
		fmt.Println("Entry.ResourceType:" + dr.Entry[i].Resource.ResourceType + "\n")
		fmt.Println("Entry.Resource.Id:" + dr.Entry[i].Resource.ID + "\n")
		fmt.Println("Entry.Resource.Text.Status: " + dr.Entry[i].Resource.Text.Status + "\n")
		fmt.Println("Entry.Resource.Text.Div: " + dr.Entry[i].Resource.Text.Div + "\n")

		fmt.Println("Entry.Resource.Code.Text: " + dr.Entry[i].Resource.Code.Text + "\n")

		for a := 0; a < len(dr.Entry[i].Resource.Code.Coding); a++ {

			fmt.Println("Entry.Resource.Code.Coding.System: " + dr.Entry[i].Resource.Code.Coding[a].System + "\n")
			fmt.Println("Entry.Resource.Code.Coding.Code: " + dr.Entry[i].Resource.Code.Coding[a].Code + "\n")
			fmt.Println("Entry.Resource.Code.Coding.Display: " + dr.Entry[i].Resource.Code.Coding[a].Display + "\n")
		}

		for b := 0; b < len(dr.Entry[i].Resource.Performer); b++ {
			fmt.Println("Entry.Performer.Resource.Reference: " + dr.Entry[i].Resource.Performer[b].Reference + "\n")
			fmt.Println("Entry.Performer.Resource.Display: " + dr.Entry[i].Resource.Performer[b].Display + "\n")
		}
		//this isn't working
		for t := 0; t < len(dr.Entry[i].Resource.ValueCodeableConcept.Coding); t++ {
			fmt.Println("Entry.Resource.ValueCodeableConcept")
			fmt.Println("System:" + dr.Entry[i].Resource.ValueCodeableConcept.Coding[t].System)
			fmt.Println("Code:" + dr.Entry[i].Resource.ValueCodeableConcept.Coding[t].Code)
			fmt.Println("Display: " + dr.Entry[i].Resource.ValueCodeableConcept.Coding[t].Display)
		}

		for d := 0; d < len(dr.Entry[i].Resource.Specimen); d++ {
			fmt.Println("Entry.Resource.Specimen.Reference: " + dr.Entry[i].Resource.Specimen[d].Reference + "\n")
			fmt.Println("Entry.Resource.Specimen.Display: " + dr.Entry[i].Resource.Specimen[d].Display + "\n")
		}

		for m := 0; m < len(dr.Entry[i].Resource.HasMember); m++ {
			fmt.Println("Entry.Resource.HasMember.Reference: " + dr.Entry[i].Resource.HasMember[m].Reference + "\n")
			fmt.Println("Entry.Resource.HasMember.Display: " + dr.Entry[i].Resource.HasMember[m].Display + "\n")
		}

		for r := 0; r < len(dr.Entry[i].Resource.Result); r++ {
			fmt.Println("Entry.Resource.Result.Reference: " + dr.Entry[i].Resource.Result[r].Reference + "\n")
			fmt.Println("Entry.Resource.Result.Display: " + dr.Entry[i].Resource.Result[r].Display + "\n")
		}

		fmt.Println("Entry.Resource.ValueQuantity.Value: " + strconv.FormatUint(dr.Entry[i].Resource.ValueQuantity.Value, 10) + "\n")
		fmt.Println("Entry.Resource.ValueQuantity.Unit: " + dr.Entry[i].Resource.ValueQuantity.Unit + "\n")
		fmt.Println("Entry.Resource.ValueQuantity.System: " + dr.Entry[i].Resource.ValueQuantity.System + "\n")
		fmt.Println("Entry.Resource.ValueQuantity.Code: " + dr.Entry[i].Resource.ValueQuantity.Code + "\n")

		for z := 0; z < len(dr.Entry[i].Resource.Interpretation); z++ {
			for v := 0; v < len(dr.Entry[i].Resource.Interpretation[z].Coding); v++ {
				fmt.Println("Entry.Resource.Interpretation.Coding.System: " + dr.Entry[i].Resource.Interpretation[z].Coding[v].System + "\n")
				fmt.Println("Entry.Resource.Interpretation.Coding.Code: " + dr.Entry[i].Resource.Interpretation[z].Coding[v].Code + "\n")
			}
		}

		for c := 0; c < len(dr.Entry[i].Resource.ReferenceRange); c++ {
			fmt.Println("Entry.Resource.ReferenceRange.Low.Value: " + strconv.FormatUint(dr.Entry[i].Resource.ReferenceRange[c].Low.Value, 10) + "\n")
			fmt.Println("Entry.Resource.ReferenceRange.Low.Unit: " + dr.Entry[i].Resource.ReferenceRange[c].Low.Unit + "\n")
			fmt.Println("Entry.Resource.ReferenceRange.Low.System: " + dr.Entry[i].Resource.ReferenceRange[c].Low.System + "\n")
			fmt.Println("Entry.Resource.ReferenceRange.Low.Code: " + dr.Entry[i].Resource.ReferenceRange[c].Low.Code + "\n")

			fmt.Println("Entry.Resource.ReferenceRange.High.Value: " + strconv.FormatUint(dr.Entry[i].Resource.ReferenceRange[c].High.Value, 10) + "\n")
			fmt.Println("Entry.Resource.ReferenceRange.High.Unit: " + dr.Entry[i].Resource.ReferenceRange[c].High.Unit + "\n")
			fmt.Println("Entry.Resource.ReferenceRange.High.System: " + dr.Entry[i].Resource.ReferenceRange[c].High.System + "\n")
			fmt.Println("Entry.Resource.ReferenceRange.High.Code: " + dr.Entry[i].Resource.ReferenceRange[c].High.Code + "\n")
		}

		for x := 0; x < len(dr.Entry[i].Resource.Container); x++ {
			for y := 0; y < len(dr.Entry[i].Resource.Container[x].Type.Coding); y++ {
				fmt.Println("Entry.Resource.Container.Type.Coding.System: " + dr.Entry[i].Resource.Container[x].Type.Coding[y].System + "\n")
				fmt.Println("Entry.Resource.Container.Type.Coding.Code: " + dr.Entry[i].Resource.Container[x].Type.Coding[y].Code + "\n")
				fmt.Println("Entry.Resource.Container.Type.Coding.Display: " + dr.Entry[i].Resource.Container[x].Type.Coding[y].Display + "\n")
			}
		}
	}
}
