package utils

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, responseBody []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
	}
	w.Write(responseBody)
}

func ErrorJsonResponse(w http.ResponseWriter, errorMessage string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
	}
	jsonResponse, _ := json.Marshal(ErrorResponse{Error: errorMessage})
	w.Write(jsonResponse)
}
