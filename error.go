package main

import "net/http"

// renderError return a generic HTTP error with StatusCode BadRequest.
func renderError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Something bad happens"))
}
