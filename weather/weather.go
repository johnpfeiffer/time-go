package weather

import (
	"encoding/json"
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
	ForecastCondition   string `json:"condition"`
	Windspeed           `json:"wspd"`
	Humidity            string `json:"humidity"`
}

// ForecastTimestamp is the time
type ForecastTimestamp struct {
	Epoch string `json:"epoch"`
	Hour  string `json:"hour"`
}

// ForecastTemperature contains the temperature in both english and metric units
type ForecastTemperature struct {
	Fahrenheit string `json:"english"`
	Celsuis    string `json:"metric"`
}

// Windspeed contains the speed in both english and metric units
type Windspeed struct {
	MilesPerHour      string `json:"english"`
	KilometersPerHour string `json:"metric"`
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
