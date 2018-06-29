package mycontrollers

import (
	"fmt"
	"net/http"
)

var defaultRobotsTxtResponse = `User-agent: *
Disallow: /`

// RobotsHandler returns the robots.txt policy http://www.robotstxt.org/robotstxt.html
func RobotsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, defaultRobotsTxtResponse)
}
