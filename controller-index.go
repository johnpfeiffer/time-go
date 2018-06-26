package main

import (
	"fmt"
	"net/http"
)

var defaultIndexResponse = `<html><body>index web page</body></html>`

// IndexHandler sends the Index page HTTP response
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, defaultIndexResponse)
}
