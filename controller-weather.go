package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// WeatherResponse wraps the forecast object
type WeatherResponse struct {
	HourlyForecast []HourForecast `json:"hourly_forecast"`
}

// HourForecast wraps the forecast for a given hour
type HourForecast struct {
	ForecastTimestamp   `json:"FCTTIME"`
	ForecastTemperature `json:"temp"`
}

// ForecastTimestamp is the time
type ForecastTimestamp struct {
	Epoch string `json:"epoch"`
}

// ForecastTemperature contains the temperature in both english and metric units
type ForecastTemperature struct {
	Fahrenheit string `json:"english"`
	Celsuis    string `json:"metric"`
}

// WeatherHandler returns the weather for a city
func WeatherHandler(w http.ResponseWriter, r *http.Request, apiKey string) {

	// TODO: totally unsafe for XSS etc.
	// 	vars := mux.Vars(r)
	// fmt.Fprintf(w, "you have requested notes from date: %v\n", vars["date"])

	weatherForecast, err := getWeatherResponse(apiKey, "CA/San_Francisco")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: table of results in a template
	theJSON, err := json.MarshalIndent(weatherForecast, "", "  ")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(theJSON))
}

// getWeatherResponse uses the APIs peculiar city locations https://api.wunderground.com/weather/api/d/docs
func getWeatherResponse(apiKey, location string) (*WeatherResponse, error) {
	c := getClient()
	url := `http://api.wunderground.com/api/` + apiKey + `/hourly/q/` + location + `.json`
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var responseObject WeatherResponse
	err = json.Unmarshal(body, &responseObject)
	if err != nil {
		return nil, err
	}
	return &responseObject, nil
}

func getClient() *http.Client {
	client := &http.Client{
		Timeout: time.Duration(30 * time.Second),
	}
	return client
}
