package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

// WriteJSONResponse: writes a JSON response with the given status code and data
func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	// data, _ = json.Marshal(data)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println(err)
	}
}
