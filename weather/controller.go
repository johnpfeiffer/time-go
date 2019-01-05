package weather

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// weatherTemplate is the handle for the singleton
var weatherTemplate *template.Template

func init() {
	weatherTemplate = template.Must(template.ParseFiles("base.tmpl", "weather/weather.html"))
}

// Controller returns the weather for a city in the requested format
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
		weatherTemplate.Execute(w, weatherForecast)
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
