package httpx

// repository of utility functions and things not domain specific. agnostic
import (
	"encoding/json"
	"fmt"
	"net/http"
)

// any - means you can type anything. any == interface{}, an interface with no methods
func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// Go -> JSON (Encode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println("failed to write JSON response", err)
	}
}

func Error(w http.ResponseWriter, message string, status int) {
	http.Error(w, message, status)
}
