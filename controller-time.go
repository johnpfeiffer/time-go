package main

import (
	"fmt"
	"net/http"
	"time"
)

// EpochHandler returns epoch seconds https://en.wikipedia.org/wiki/Unix_time
func EpochHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%d", t.Unix())
}

// TimeHandler returns various time formats
func TimeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().UTC()
	timeString := t.Format(time.RFC3339)
	output := "UTC , " + timeString

	// https://en.wikipedia.org/wiki/Time_in_China
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	china := time.FixedZone("China Standard Time", secondsEastOfUTC)
	regionalTime := t.In(china)
	output = output + "\n" + "China Standard Time , " + regionalTime.Format(time.RFC3339)

	// TODO: // https://en.wikipedia.org/wiki/Daylight_saving_time_by_country

	// TODO: https://golang.org/pkg/time/#LoadLocationFromTZData,
	// https://golang.org/doc/go1.10 ,  https://www.iana.org/time-zones

	// for _, v := range timezones {
	// 	location, err := time.LoadLocation(v)
	// 	if err != nil {
	// 		log.Fatalf("bad time", err)
	// 	}
	// 	regionalTime := t.In(location)
	// 	timeString := regionalTime.Format(time.RFC3339)
	// 	output = output + v + "," + timeString
	// }

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", output)
}
