package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Meters []Meter `json:"data"`
}

type Meter struct {
	ID        string `json:"_id"`
	EAN       string `json:"ean"`
	Timestamp string `json:"created"`
	Value     int    `json:"last"`
}

func getMeters() {
	response, err := http.Get("https://app.metry.io/api/2.0/meters/?access_token=a4e9522bb4a335e850967754207859aa972e2e20ca9871d3f276390b757c&id=5825c85d22c8aa00606b9dd5")

	if err != nil {
		fmt.Println("The http request failed with the error %s\n", err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(data))

	var responseObject Response
	json.Unmarshal(data, &responseObject)
	fmt.Println(len(responseObject.Meters))

	for i := 0; i < len(responseObject.Meters); i++ {
		//fmt.Println(responseObject.Meters[i].Value)
		//fmt.Println("The meter's id is " + responseObject.Meters[i].ID + ", the ean is " + responseObject.Meters[i].EAN + ", timestamp is " + responseObject.Meters[i].Timestamp + ", value is " + responseObject.Meters[i].Value)
	}
}

func main() {
	fmt.Println("Starting the application...")

	getMeters()

	fmt.Println("Terminating the application...")
}
