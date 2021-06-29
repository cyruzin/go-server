package rest

import (
	"encoding/json"
	"net/http"
)

// APIMessage is a struct for generic JSON response.
type APIMessage struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}

// ToJSON returns a JSON response.
func ToJSON(
	w http.ResponseWriter,
	httpCode int,
	dest interface{},
) {
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(dest)
}
