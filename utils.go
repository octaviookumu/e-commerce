package main

import (
	"encoding/json"
	"net/http"
)

// any - means you can type anything. any == interface{}, an interface with no methods
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data) // Go -> JSON (Encode)
}

func handleError(w http.ResponseWriter, message string, status int) {
	http.Error(w, message, status)
}
