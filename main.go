package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"Metry-API-codetest/counters"

	"Metry-API-codetest/data"
)

var TOKEN string
var BASEURL = "https://app.metry.io/api/v2/"

func GetMeters() {
	response, err := http.Get(BASEURL + "meters/?access_token=" + TOKEN)

	if err != nil {
		fmt.Printf("The http request failed with the error %s\n", err)
	}

	resp, _ := ioutil.ReadAll(response.Body)

	var responseObject data.Response
	json.Unmarshal(resp, &responseObject)

	responseObject.Meters = SetComsumptionToMeter(responseObject.Meters)

	output, err := json.Marshal(responseObject)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}

func SetComsumptionToMeter(meters []data.Meter) []data.Meter {

	for i := 0; i < len(meters); i++ {
		activeOrTemporaryBox := meters[i].Box == "active" || meters[i].Box == "temporary"

		if meters[i].Revoked == false && activeOrTemporaryBox {
			meters[i].Consumption = GetConsumption(meters[i].Root.ID)

		} else {
			meters[i].Consumption = "null"
		}
	}

	return meters
}

func GetConsumption(meterID string) string {
	resp, err := http.Get(BASEURL + "consumptions/" + meterID + "/month/2021?access_token=" + TOKEN)

	if err != nil {
		fmt.Printf("The http request failed with the error %s\n", err)
	}

	consdata, _ := ioutil.ReadAll(resp.Body)

	var response data.ConsumptionResponse
	json.Unmarshal(consdata, &response)

	latestEnergyConsumption := counters.GetLatestEnergyConsumption(response.Data[0].Periods[0].Energy)
	return strconv.Itoa(latestEnergyConsumption)
}

func main() {

	fmt.Println("Starting the application...")
	fmt.Println("Insert access token below: ")
	var token string
	fmt.Scanln(&token)
	TOKEN = token
	fmt.Println("Processing your request...")

	GetMeters()

	fmt.Println("Terminating the application...")
}
