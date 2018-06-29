package weather

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

var WeatherTemplate = GetWeatherTemplate()

// GetWeatherTemplate returns the parsed file as a template object
func GetWeatherTemplate() *template.Template {
	return template.Must(template.ParseFiles("base.tmpl", "weather/weather.html"))
}

// Controller is returns the weather for a city in the requested format
func Controller(w http.ResponseWriter, r *http.Request, apiKey string) {
	// ?city=CA/San_Francisco
	city := r.FormValue("city")
	if city == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: a city parameter is required\n"))
		return
	}

	weatherForecast, err := getWeatherResponse(apiKey, city)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	format := r.FormValue("format")
	if format == "html" {
		WeatherTemplate.Execute(w, weatherForecast)
		return
	}

	theJSON, err := json.MarshalIndent(weatherForecast, "", "  ")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(theJSON))
}
