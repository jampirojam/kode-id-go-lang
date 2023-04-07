package service

import (
	"ojam-test/c3/s1/app/generator"
)

func WeatherCheck() {
	maxWindValue := 30
	maxWaterValue := 16
	windValue := generator.RandomNumber(maxWindValue)
	watervalue := generator.RandomNumber(maxWaterValue)
	generator.WeatherStatusGenerator(windValue, watervalue)
}