package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func router(weatherKey string) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/weather",
		func(w http.ResponseWriter, r *http.Request) {
			WeatherHandler(w, r, weatherKey)
		}).Methods("GET")

	r.HandleFunc("/time/epoch", EpochHandler).Methods("GET")
	r.HandleFunc("/time", TimeHandler).Methods("GET")
	r.HandleFunc("/robots.txt", RobotsHandler).Methods("GET")
	r.HandleFunc("/", IndexHandler).Methods("GET")
	return r
}

func main() {
	port := getEnvOrDefault("PORT", "8080")
	weatherKey := getEnvOrDefault("WEATHERKEY", "")
	r := router(weatherKey)
	http.Handle("/", r)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, loggedRouter))
}

func getEnvOrDefault(key, defaultValue string) string {
	result := defaultValue
	val, ok := os.LookupEnv(key)
	if ok {
		result = val
	}
	return result
}

func exitIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
