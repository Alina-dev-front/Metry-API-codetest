package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Response struct {
	Meters []Meter `json:"data"`
}

type Meter struct {
	Root struct {
		ID string `json:"_id"`
	} `json:"root"`
	EAN         string `json:"ean"`
	Consumption int
	Box         string
	Revoked     bool
}

type ConsumptionResponse struct {
	Data []Consumption `json:"data"`
}
type Consumption struct {
	MeterID string `json:"meter_id"`
	Periods []struct {
		Energy    []int  `json:"energy"`
		StartDate string `json:"start_date"`
	} `json:"periods"`
}

func GetMeters() {
	response, err := http.Get("https://app.metry.io/api/v2/meters/?access_token=a4e9522bb4a335e850967754207859aa972e2e20ca9871d3f276390b757c&id=5825c85d22c8aa00606b9dd5")

	if err != nil {
		fmt.Println("The http request failed with the error %s\n", err)
	}

	data, _ := ioutil.ReadAll(response.Body)

	var responseObject Response
	json.Unmarshal(data, &responseObject)

	for i := 0; i < len(responseObject.Meters); i++ {
		meter := responseObject.Meters[i]

		if meter.Revoked == false && (meter.Box == "active" || meter.Box == "temporary") {
			meter.Consumption = GetConsumption(meter.Root.ID)
		} else {
			meter.Consumption = 0
		}

		fmt.Println("The meter's id is " + meter.Root.ID + ", the ean is " + meter.EAN + ", consumption is " + strconv.Itoa(meter.Consumption))

	}

	output, err := json.Marshal(responseObject)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}

func GetConsumption(meterID string) int {
	resp, err := http.Get("https://app.metry.io/api/v2/consumptions/" + meterID + "/month/2021?access_token=a4e9522bb4a335e850967754207859aa972e2e20ca9871d3f276390b757c")

	if err != nil {
		fmt.Println("The http request failed with the error %s\n", err)
	}

	consdata, _ := ioutil.ReadAll(resp.Body)

	var response ConsumptionResponse
	json.Unmarshal(consdata, &response)

	var latestEnergyConsumption = 0
	for _, val := range response.Data[0].Periods[0].Energy {
		if val > 0 {
			latestEnergyConsumption = val
		}
	}
	return latestEnergyConsumption
}

func main() {
	fmt.Println("Starting the application...")

	GetMeters()

	fmt.Println("Terminating the application...")
}
