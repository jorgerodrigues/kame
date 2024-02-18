package utils

import (
	"encoding/json"
	"net/http"

	"github.com/jorgerodrigues/upkame/internal/logger"
)

// ResponsePayload is a generic structure for API responses
type ResponsePayload struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// SendJSONResponse is a helper function to send JSON responses
func SendJSONResponse(w http.ResponseWriter, statusCode int, payload ResponsePayload) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	logger := logger.NewLogger()

	if payload != (ResponsePayload{}) { // Check if the payload is not the zero value before encoding
		err := json.NewEncoder(w).Encode(payload)
		if err != nil {
			// Log the error, and send a 500 internal server error code
			logger.Error("Failed to encode response", err)
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	}
}
