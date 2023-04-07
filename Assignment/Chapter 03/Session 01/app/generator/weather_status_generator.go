package generator

import (
	"encoding/json"
	"fmt"
	"ojam-test/c3/s1/app/model"
)

func WeatherStatusGenerator(windValue, waterValue int) {
	var windStatus, waterStatus string
	if windValue > 15 {
		windStatus = "Danger"
	} else if windValue < 7 {
		windStatus = "Safe"
	} else {
		windStatus = "Alert"
	}

	if waterValue > 8 {
		waterStatus = "Danger"
	} else if waterValue < 6 {
		waterStatus = "Safe"
	} else {
		waterStatus = "Alert"
	}

	status := model.ElementValue{
		Wind: windValue,
		Water: waterValue,
	}

	body, err := json.MarshalIndent(status, "", "   ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
	fmt.Printf("wind status: %-5s\n", windStatus)
	fmt.Printf("water status: %-5s\n", waterStatus)
}